// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: v1/anomaly_service.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "bytebase.v1";

export interface SearchAnomaliesRequest {
  /**
   * The parent resource whose anomalies are to be listed.
   * Format: projects/{project}
   */
  parent: string;
  /**
   * filter is the filter to apply on the search anomaly request,
   * follow the [ebnf](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) syntax.
   * Only support filter by resource and type for now.
   * For example:
   * Search the anomalies of a specific resource: 'resource="instances/{instance}/databases/{database}".'
   * Search the specified types of anomalies: 'type="MIGRATION_SCHEMA".'
   */
  filter: string;
  /**
   * Not used.
   * The maximum number of anomalies to return. The service may return fewer than
   * this value.
   * If unspecified, at most 10 anomalies will be returned.
   * The maximum value is 1000; values above 1000 will be coerced to 1000.
   */
  pageSize: number;
  /**
   * Not used.
   * A page token, received from a previous `SearchAnomalies` call.
   * Provide this to retrieve the subsequent page.
   *
   * When paginating, all other parameters provided to `SearchAnomalies` must match
   * the call that provided the page token.
   */
  pageToken: string;
}

export interface SearchAnomaliesResponse {
  /** anomalies is the list of anomalies. */
  anomalies: Anomaly[];
  /**
   * Not used. A token, which can be sent as `page_token` to retrieve the next page.
   * If this field is omitted, there are no subsequent pages.
   */
  nextPageToken: string;
}

export interface Anomaly {
  /**
   * The resource that is the target of the operation.
   * Format:
   * - Database: instances/{instance}/databases/{database}
   */
  resource: string;
  /** type is the type of the anomaly. */
  type: Anomaly_AnomalyType;
  /** severity is the severity of the anomaly. */
  severity: Anomaly_AnomalySeverity;
  createTime: Timestamp | undefined;
}

/** AnomalyType is the type of the anomaly. */
export enum Anomaly_AnomalyType {
  /** ANOMALY_TYPE_UNSPECIFIED - Unspecified anomaly type. */
  ANOMALY_TYPE_UNSPECIFIED = "ANOMALY_TYPE_UNSPECIFIED",
  /**
   * DATABASE_CONNECTION - Database level anomaly.
   *
   * DATABASE_CONNECTION is the anomaly type for database connection, e.g. the database had been deleted.
   */
  DATABASE_CONNECTION = "DATABASE_CONNECTION",
  /**
   * DATABASE_SCHEMA_DRIFT - DATABASE_SCHEMA_DRIFT is the anomaly type for database schema drift,
   * e.g. the database schema had been changed without bytebase migration.
   */
  DATABASE_SCHEMA_DRIFT = "DATABASE_SCHEMA_DRIFT",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function anomaly_AnomalyTypeFromJSON(object: any): Anomaly_AnomalyType {
  switch (object) {
    case 0:
    case "ANOMALY_TYPE_UNSPECIFIED":
      return Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED;
    case 5:
    case "DATABASE_CONNECTION":
      return Anomaly_AnomalyType.DATABASE_CONNECTION;
    case 6:
    case "DATABASE_SCHEMA_DRIFT":
      return Anomaly_AnomalyType.DATABASE_SCHEMA_DRIFT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Anomaly_AnomalyType.UNRECOGNIZED;
  }
}

export function anomaly_AnomalyTypeToJSON(object: Anomaly_AnomalyType): string {
  switch (object) {
    case Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED:
      return "ANOMALY_TYPE_UNSPECIFIED";
    case Anomaly_AnomalyType.DATABASE_CONNECTION:
      return "DATABASE_CONNECTION";
    case Anomaly_AnomalyType.DATABASE_SCHEMA_DRIFT:
      return "DATABASE_SCHEMA_DRIFT";
    case Anomaly_AnomalyType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function anomaly_AnomalyTypeToNumber(object: Anomaly_AnomalyType): number {
  switch (object) {
    case Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED:
      return 0;
    case Anomaly_AnomalyType.DATABASE_CONNECTION:
      return 5;
    case Anomaly_AnomalyType.DATABASE_SCHEMA_DRIFT:
      return 6;
    case Anomaly_AnomalyType.UNRECOGNIZED:
    default:
      return -1;
  }
}

/** AnomalySeverity is the severity of the anomaly. */
export enum Anomaly_AnomalySeverity {
  /** ANOMALY_SEVERITY_UNSPECIFIED - Unspecified anomaly severity. */
  ANOMALY_SEVERITY_UNSPECIFIED = "ANOMALY_SEVERITY_UNSPECIFIED",
  /** MEDIUM - MEDIUM is the info level anomaly severity. */
  MEDIUM = "MEDIUM",
  /** HIGH - HIGH is the warning level anomaly severity. */
  HIGH = "HIGH",
  /** CRITICAL - CRITICAL is the critical level anomaly severity. */
  CRITICAL = "CRITICAL",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function anomaly_AnomalySeverityFromJSON(object: any): Anomaly_AnomalySeverity {
  switch (object) {
    case 0:
    case "ANOMALY_SEVERITY_UNSPECIFIED":
      return Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED;
    case 1:
    case "MEDIUM":
      return Anomaly_AnomalySeverity.MEDIUM;
    case 2:
    case "HIGH":
      return Anomaly_AnomalySeverity.HIGH;
    case 3:
    case "CRITICAL":
      return Anomaly_AnomalySeverity.CRITICAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Anomaly_AnomalySeverity.UNRECOGNIZED;
  }
}

export function anomaly_AnomalySeverityToJSON(object: Anomaly_AnomalySeverity): string {
  switch (object) {
    case Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED:
      return "ANOMALY_SEVERITY_UNSPECIFIED";
    case Anomaly_AnomalySeverity.MEDIUM:
      return "MEDIUM";
    case Anomaly_AnomalySeverity.HIGH:
      return "HIGH";
    case Anomaly_AnomalySeverity.CRITICAL:
      return "CRITICAL";
    case Anomaly_AnomalySeverity.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function anomaly_AnomalySeverityToNumber(object: Anomaly_AnomalySeverity): number {
  switch (object) {
    case Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED:
      return 0;
    case Anomaly_AnomalySeverity.MEDIUM:
      return 1;
    case Anomaly_AnomalySeverity.HIGH:
      return 2;
    case Anomaly_AnomalySeverity.CRITICAL:
      return 3;
    case Anomaly_AnomalySeverity.UNRECOGNIZED:
    default:
      return -1;
  }
}

function createBaseSearchAnomaliesRequest(): SearchAnomaliesRequest {
  return { parent: "", filter: "", pageSize: 0, pageToken: "" };
}

export const SearchAnomaliesRequest: MessageFns<SearchAnomaliesRequest> = {
  encode(message: SearchAnomaliesRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.parent !== "") {
      writer.uint32(10).string(message.parent);
    }
    if (message.filter !== "") {
      writer.uint32(18).string(message.filter);
    }
    if (message.pageSize !== 0) {
      writer.uint32(24).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(34).string(message.pageToken);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SearchAnomaliesRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchAnomaliesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.parent = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.filter = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.pageToken = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SearchAnomaliesRequest {
    return {
      parent: isSet(object.parent) ? globalThis.String(object.parent) : "",
      filter: isSet(object.filter) ? globalThis.String(object.filter) : "",
      pageSize: isSet(object.pageSize) ? globalThis.Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? globalThis.String(object.pageToken) : "",
    };
  },

  toJSON(message: SearchAnomaliesRequest): unknown {
    const obj: any = {};
    if (message.parent !== "") {
      obj.parent = message.parent;
    }
    if (message.filter !== "") {
      obj.filter = message.filter;
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    if (message.pageToken !== "") {
      obj.pageToken = message.pageToken;
    }
    return obj;
  },

  create(base?: DeepPartial<SearchAnomaliesRequest>): SearchAnomaliesRequest {
    return SearchAnomaliesRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SearchAnomaliesRequest>): SearchAnomaliesRequest {
    const message = createBaseSearchAnomaliesRequest();
    message.parent = object.parent ?? "";
    message.filter = object.filter ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseSearchAnomaliesResponse(): SearchAnomaliesResponse {
  return { anomalies: [], nextPageToken: "" };
}

export const SearchAnomaliesResponse: MessageFns<SearchAnomaliesResponse> = {
  encode(message: SearchAnomaliesResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.anomalies) {
      Anomaly.encode(v!, writer.uint32(10).fork()).join();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SearchAnomaliesResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchAnomaliesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.anomalies.push(Anomaly.decode(reader, reader.uint32()));
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SearchAnomaliesResponse {
    return {
      anomalies: globalThis.Array.isArray(object?.anomalies)
        ? object.anomalies.map((e: any) => Anomaly.fromJSON(e))
        : [],
      nextPageToken: isSet(object.nextPageToken) ? globalThis.String(object.nextPageToken) : "",
    };
  },

  toJSON(message: SearchAnomaliesResponse): unknown {
    const obj: any = {};
    if (message.anomalies?.length) {
      obj.anomalies = message.anomalies.map((e) => Anomaly.toJSON(e));
    }
    if (message.nextPageToken !== "") {
      obj.nextPageToken = message.nextPageToken;
    }
    return obj;
  },

  create(base?: DeepPartial<SearchAnomaliesResponse>): SearchAnomaliesResponse {
    return SearchAnomaliesResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SearchAnomaliesResponse>): SearchAnomaliesResponse {
    const message = createBaseSearchAnomaliesResponse();
    message.anomalies = object.anomalies?.map((e) => Anomaly.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseAnomaly(): Anomaly {
  return {
    resource: "",
    type: Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED,
    severity: Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED,
    createTime: undefined,
  };
}

export const Anomaly: MessageFns<Anomaly> = {
  encode(message: Anomaly, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.resource !== "") {
      writer.uint32(10).string(message.resource);
    }
    if (message.type !== Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED) {
      writer.uint32(16).int32(anomaly_AnomalyTypeToNumber(message.type));
    }
    if (message.severity !== Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED) {
      writer.uint32(24).int32(anomaly_AnomalySeverityToNumber(message.severity));
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(message.createTime, writer.uint32(74).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Anomaly {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnomaly();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.resource = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.type = anomaly_AnomalyTypeFromJSON(reader.int32());
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.severity = anomaly_AnomalySeverityFromJSON(reader.int32());
          continue;
        }
        case 9: {
          if (tag !== 74) {
            break;
          }

          message.createTime = Timestamp.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Anomaly {
    return {
      resource: isSet(object.resource) ? globalThis.String(object.resource) : "",
      type: isSet(object.type)
        ? anomaly_AnomalyTypeFromJSON(object.type)
        : Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED,
      severity: isSet(object.severity)
        ? anomaly_AnomalySeverityFromJSON(object.severity)
        : Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED,
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
    };
  },

  toJSON(message: Anomaly): unknown {
    const obj: any = {};
    if (message.resource !== "") {
      obj.resource = message.resource;
    }
    if (message.type !== Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED) {
      obj.type = anomaly_AnomalyTypeToJSON(message.type);
    }
    if (message.severity !== Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED) {
      obj.severity = anomaly_AnomalySeverityToJSON(message.severity);
    }
    if (message.createTime !== undefined) {
      obj.createTime = fromTimestamp(message.createTime).toISOString();
    }
    return obj;
  },

  create(base?: DeepPartial<Anomaly>): Anomaly {
    return Anomaly.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Anomaly>): Anomaly {
    const message = createBaseAnomaly();
    message.resource = object.resource ?? "";
    message.type = object.type ?? Anomaly_AnomalyType.ANOMALY_TYPE_UNSPECIFIED;
    message.severity = object.severity ?? Anomaly_AnomalySeverity.ANOMALY_SEVERITY_UNSPECIFIED;
    message.createTime = (object.createTime !== undefined && object.createTime !== null)
      ? Timestamp.fromPartial(object.createTime)
      : undefined;
    return message;
  },
};

export type AnomalyServiceDefinition = typeof AnomalyServiceDefinition;
export const AnomalyServiceDefinition = {
  name: "AnomalyService",
  fullName: "bytebase.v1.AnomalyService",
  methods: {
    searchAnomalies: {
      name: "SearchAnomalies",
      requestType: SearchAnomaliesRequest,
      requestStream: false,
      responseType: SearchAnomaliesResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          800016: [new Uint8Array([2])],
          578365826: [
            new Uint8Array([
              45,
              58,
              1,
              42,
              34,
              40,
              47,
              118,
              49,
              47,
              123,
              112,
              97,
              114,
              101,
              110,
              116,
              61,
              112,
              114,
              111,
              106,
              101,
              99,
              116,
              115,
              47,
              42,
              125,
              47,
              97,
              110,
              111,
              109,
              97,
              108,
              105,
              101,
              115,
              58,
              115,
              101,
              97,
              114,
              99,
              104,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(Math.trunc(date.getTime() / 1_000));
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
