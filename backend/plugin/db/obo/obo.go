// Package obo is for OceanBase Oracle mode
package obo

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/common/log"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	v1pb "github.com/bytebase/bytebase/backend/generated-go/v1"
	"github.com/bytebase/bytebase/backend/plugin/db"
	"github.com/bytebase/bytebase/backend/plugin/db/util"
	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	plsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/plsql"

	// Register OceanBase Oracle mode driver.
	_ "github.com/mattn/go-oci8"
)

func init() {
	db.Register(storepb.Engine_OCEANBASE_ORACLE, newDriver)
}

type Driver struct {
	db           *sql.DB
	databaseName string
}

func newDriver() db.Driver {
	return &Driver{}
}

func (d *Driver) Open(_ context.Context, _ storepb.Engine, config db.ConnectionConfig) (db.Driver, error) {
	databaseName := func() string {
		if config.ConnectionContext.DatabaseName != "" {
			return config.ConnectionContext.DatabaseName
		}
		i := strings.Index(config.DataSource.Username, "@")
		if i == -1 {
			return config.DataSource.Username
		}
		return config.DataSource.Username[:i]
	}()

	// Usename format: {user}@{tenant}#{cluster}
	// User is required, others are optional.
	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", url.PathEscape(config.DataSource.Username), url.PathEscape(config.Password), config.DataSource.Host, config.DataSource.Port, url.PathEscape(databaseName))

	db, err := sql.Open("oci8", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open connection")
	}

	d.db = db
	d.databaseName = databaseName
	return d, nil
}

func (d *Driver) Close(context.Context) error {
	return d.db.Close()
}

func (d *Driver) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

func (d *Driver) GetDB() *sql.DB {
	return d.db
}

func (d *Driver) Execute(ctx context.Context, statement string, opts db.ExecuteOptions) (int64, error) {
	if opts.CreateDatabase {
		return 0, errors.New("create database is not supported for OceanBase Oracle mode")
	}

	// Parse transaction mode from the script
	transactionMode, cleanedStatement := base.ParseTransactionMode(statement)
	statement = cleanedStatement

	// Apply default when transaction mode is not specified
	if transactionMode == common.TransactionModeUnspecified {
		transactionMode = common.GetDefaultTransactionMode()
	}

	// Use Oracle sql parser.
	singleSQLs, err := plsqlparser.SplitSQL(statement)
	if err != nil {
		return 0, err
	}
	singleSQLs = base.FilterEmptySQL(singleSQLs)
	if len(singleSQLs) == 0 {
		return 0, nil
	}

	// Execute based on transaction mode
	if transactionMode == common.TransactionModeOff {
		return d.executeInAutoCommitMode(ctx, singleSQLs, opts)
	}
	return d.executeInTransactionMode(ctx, singleSQLs, opts)
}

func (d *Driver) executeInTransactionMode(ctx context.Context, singleSQLs []base.SingleSQL, opts db.ExecuteOptions) (int64, error) {
	conn, err := d.db.Conn(ctx)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get connection")
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		opts.LogTransactionControl(storepb.TaskRunLog_TransactionControl_BEGIN, err.Error())
		return 0, errors.Wrapf(err, "failed to begin transaction")
	}
	opts.LogTransactionControl(storepb.TaskRunLog_TransactionControl_BEGIN, "")

	committed := false
	defer func() {
		if !committed {
			err := tx.Rollback()
			var rerr string
			if err != nil && !errors.Is(err, sql.ErrTxDone) {
				rerr = err.Error()
				slog.Debug("failed to rollback transaction", log.BBError(err))
			}
			opts.LogTransactionControl(storepb.TaskRunLog_TransactionControl_ROLLBACK, rerr)
		}
	}()

	totalRowsAffected := int64(0)
	for i, singleSQL := range singleSQLs {
		opts.LogCommandExecute([]int32{int32(i)})
		sqlResult, err := tx.ExecContext(ctx, singleSQL.Text)
		if err != nil {
			opts.LogCommandResponse([]int32{int32(i)}, 0, nil, err.Error())
			return 0, &db.ErrorWithPosition{
				Err:   errors.Wrapf(err, "failed to execute context in a transaction"),
				Start: singleSQL.Start,
				End:   singleSQL.End,
			}
		}
		rowsAffected, err := sqlResult.RowsAffected()
		if err != nil {
			// Since we cannot differentiate DDL and DML yet, we have to ignore the error.
			slog.Debug("rowsAffected returns error", log.BBError(err))
		}
		opts.LogCommandResponse([]int32{int32(i)}, int32(rowsAffected), nil, "")
		totalRowsAffected += rowsAffected
	}

	if err := tx.Commit(); err != nil {
		opts.LogTransactionControl(storepb.TaskRunLog_TransactionControl_COMMIT, err.Error())
		return 0, errors.Wrapf(err, "failed to commit transaction")
	}
	opts.LogTransactionControl(storepb.TaskRunLog_TransactionControl_COMMIT, "")
	committed = true
	return totalRowsAffected, nil
}

func (d *Driver) executeInAutoCommitMode(ctx context.Context, singleSQLs []base.SingleSQL, opts db.ExecuteOptions) (int64, error) {
	conn, err := d.db.Conn(ctx)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get connection")
	}
	defer conn.Close()

	totalRowsAffected := int64(0)
	for i, singleSQL := range singleSQLs {
		opts.LogCommandExecute([]int32{int32(i)})
		sqlResult, err := conn.ExecContext(ctx, singleSQL.Text)
		if err != nil {
			opts.LogCommandResponse([]int32{int32(i)}, 0, nil, err.Error())
			return totalRowsAffected, &db.ErrorWithPosition{
				Err:   errors.Wrapf(err, "failed to execute context in auto-commit mode"),
				Start: singleSQL.Start,
				End:   singleSQL.End,
			}
		}
		rowsAffected, err := sqlResult.RowsAffected()
		if err != nil {
			// Since we cannot differentiate DDL and DML yet, we have to ignore the error.
			slog.Debug("rowsAffected returns error", log.BBError(err))
		}
		opts.LogCommandResponse([]int32{int32(i)}, int32(rowsAffected), nil, "")
		totalRowsAffected += rowsAffected
	}
	return totalRowsAffected, nil
}

func (*Driver) QueryConn(ctx context.Context, conn *sql.Conn, statement string, queryContext db.QueryContext) ([]*v1pb.QueryResult, error) {
	singleSQLs, err := plsqlparser.SplitSQL(statement)
	if err != nil {
		return nil, err
	}
	singleSQLs = base.FilterEmptySQL(singleSQLs)
	if len(singleSQLs) == 0 {
		return nil, nil
	}

	var results []*v1pb.QueryResult
	for _, singleSQL := range singleSQLs {
		statement := singleSQL.Text
		if queryContext.Explain {
			startTime := time.Now()
			randNum, err := rand.Int(rand.Reader, big.NewInt(999))
			if err != nil {
				return nil, errors.Wrapf(err, "failed to generate random statement ID")
			}
			randomID := fmt.Sprintf("%d%d", startTime.UnixMilli(), randNum.Int64())

			if _, err := conn.ExecContext(ctx, fmt.Sprintf("EXPLAIN PLAN SET STATEMENT_ID = '%s' FOR %s", randomID, statement)); err != nil {
				return nil, err
			}
			statement = fmt.Sprintf(`SELECT LPAD(' ', LEVEL-1) || OPERATION || ' (' || OPTIONS || ')' "Operation", OBJECT_NAME "Object", OPTIMIZER "Optimizer", COST "Cost", CARDINALITY "Cardinality", BYTES "Bytes", PARTITION_START "Partition Start", PARTITION_ID "Partition ID", ACCESS_PREDICATES "Access Predicates" FROM PLAN_TABLE START WITH ID = 0 AND statement_id = '%s' CONNECT BY PRIOR ID=PARENT_ID AND statement_id = '%s' ORDER BY id`, randomID, randomID)
		}

		if !queryContext.Explain && queryContext.Limit > 0 {
			statement = getStatementWithResultLimit(statement, queryContext.Limit)
		}

		_, allQuery, err := base.ValidateSQLForEditor(storepb.Engine_ORACLE, statement)
		if err != nil {
			return nil, err
		}
		startTime := time.Now()
		queryResult, err := func() (*v1pb.QueryResult, error) {
			if allQuery {
				rows, err := conn.QueryContext(ctx, statement)
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				r, err := util.RowsToQueryResult(rows, makeValueByTypeName, convertValue, queryContext.MaximumSQLResultSize)
				if err != nil {
					return nil, err
				}
				if err := rows.Err(); err != nil {
					return nil, err
				}
				return r, nil
			}

			sqlResult, err := conn.ExecContext(ctx, statement)
			if err != nil {
				return nil, err
			}
			affectedRows, err := sqlResult.RowsAffected()
			if err != nil {
				slog.Info("rowsAffected returns error", log.BBError(err))
			}
			return util.BuildAffectedRowsResult(affectedRows, nil), nil
		}()
		stop := false
		if err != nil {
			queryResult = &v1pb.QueryResult{
				Error: err.Error(),
			}
			stop = true
		}
		queryResult.Statement = statement
		queryResult.Latency = durationpb.New(time.Since(startTime))
		queryResult.RowsCount = int64(len(queryResult.Rows))
		results = append(results, queryResult)
		if stop {
			break
		}
	}

	return results, nil
}

func getStatementWithResultLimit(statement string, limit int) string {
	return fmt.Sprintf("SELECT * FROM (%s) WHERE ROWNUM <= %d", util.TrimStatement(statement), limit)
}
