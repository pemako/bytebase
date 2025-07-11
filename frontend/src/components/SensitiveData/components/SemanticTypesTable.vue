<template>
  <NDataTable
    :size="size"
    :data="semanticItemList"
    :columns="columnList"
    :striped="true"
    :bordered="true"
    :row-props="rowProps"
    :row-key="(item: SemanticItem) => item.item.id"
  />

  <MaskingAlgorithmsCreateDrawer
    :show="state.pendingEditSemanticIndex !== undefined"
    :algorithm="
      state.pendingEditSemanticIndex !== undefined
        ? getMaskingType(
            semanticItemList[state.pendingEditSemanticIndex].item.algorithm
          )
          ? semanticItemList[state.pendingEditSemanticIndex].item.algorithm
          : undefined
        : undefined
    "
    :readonly="readonly"
    @apply="onAlgorithmUpsert"
    @dismiss="onDrawerDismiss"
  />
</template>

<script lang="tsx" setup>
import {
  CheckIcon,
  PencilIcon,
  TrashIcon,
  Undo2Icon,
  InfoIcon,
} from "lucide-vue-next";
import { NPopconfirm, NInput, NDataTable, NTooltip } from "naive-ui";
import type { DataTableColumn } from "naive-ui";
import { computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { MiniActionButton } from "@/components/v2";
import {
  type Algorithm,
  type SemanticTypeSetting_SemanticType,
} from "@/types/proto-es/v1/setting_service_pb";
import MaskingAlgorithmsCreateDrawer from "./MaskingAlgorithmsCreateDrawer.vue";
import { getMaskingType } from "./utils";

type SemanticItemMode = "NORMAL" | "CREATE" | "EDIT";

export interface SemanticItem {
  mode: SemanticItemMode;
  dirty: boolean;
  item: SemanticTypeSetting_SemanticType;
}

interface LocalState {
  pendingEditSemanticIndex?: number;
  processing: boolean;
}

const props = withDefaults(
  defineProps<{
    size?: "small" | "medium" | "large";
    readonly: boolean;
    semanticItemList: SemanticItem[];
    rowClickable: boolean;
  }>(),
  { size: "medium" }
);

const emit = defineEmits<{
  (event: "select", id: string): void;
  (event: "remove", index: number): void;
  (event: "cancel", index: number): void;
  (event: "confirm", index: number): void;
}>();

const state = reactive<LocalState>({
  processing: false,
});

const { t } = useI18n();

const onDrawerDismiss = () => {
  state.pendingEditSemanticIndex = undefined;
};

const isBuiltinSemanticType = (item: SemanticTypeSetting_SemanticType) => {
  return item.id.startsWith("bb.");
};

const isReadonly = (item: SemanticTypeSetting_SemanticType) => {
  return props.readonly || isBuiltinSemanticType(item);
};

const onAlgorithmUpsert = async (maskingAlgorithm: Algorithm) => {
  if (state.pendingEditSemanticIndex === undefined) {
    return;
  }
  onInput(
    state.pendingEditSemanticIndex,
    (data) => (data.item.algorithm = maskingAlgorithm)
  );
  if (
    !isConfirmDisabled(props.semanticItemList[state.pendingEditSemanticIndex])
  ) {
    emit("confirm", state.pendingEditSemanticIndex);
  }
  onDrawerDismiss();
};

const columnList = computed(() => {
  const columns: DataTableColumn<SemanticItem>[] = [
    {
      key: "id",
      title: "ID",
      width: 350,
      resizable: true,
      render: (item) => {
        return <h3 class="break-normal">{item.item.id}</h3>;
      },
    },
    {
      key: "title",
      title: t("settings.sensitive-data.semantic-types.table.semantic-type"),
      width: "minmax(min-content, auto)",
      resizable: true,
      render: (item, row) => {
        if (item.mode === "NORMAL") {
          return <h3 class="break-normal">{item.item.title}</h3>;
        }
        return (
          <NInput
            value={item.item.title}
            size="small"
            type="text"
            placeholder={t(
              "settings.sensitive-data.semantic-types.table.semantic-type"
            )}
            onInput={(val: string) =>
              onInput(row, (data) => (data.item.title = val))
            }
          />
        );
      },
    },
    {
      key: "description",
      title: t("settings.sensitive-data.semantic-types.table.description"),
      resizable: true,
      width: "minmax(min-content, auto)",
      render: (item, row) => {
        if (item.mode === "NORMAL") {
          return <h3>{item.item.description}</h3>;
        }
        return (
          <NInput
            value={item.item.description}
            size="small"
            type="text"
            placeholder={t(
              "settings.sensitive-data.semantic-types.table.description"
            )}
            onInput={(val: string) =>
              onInput(row, (data) => (data.item.description = val))
            }
          />
        );
      },
    },
  ];

  columns.push({
    key: "algorithm",
    title: t("settings.sensitive-data.semantic-types.table.masking-algorithm"),
    resizable: true,
    width: "minmax(min-content, auto)",
    render: (item, row) => {
      return (
        <div class="flex items-center space-x-1">
          {isBuiltinSemanticType(item.item) ? (
            <h3>
              {t(
                `dynamic.settings.sensitive-data.semantic-types.template.${item.item.id.split(".").join("-")}.title`
              )}
            </h3>
          ) : (
            <h3>
              {getMaskingType(item.item.algorithm)
                ? t(
                    `settings.sensitive-data.algorithms.${getMaskingType(item.item.algorithm)?.toLowerCase()}.self`
                  )
                : "N/A"}
            </h3>
          )}
          {isBuiltinSemanticType(item.item) && (
            <NTooltip>
              {{
                trigger: () => <InfoIcon class="w-4" />,
                default: () => (
                  <div class="whitespace-pre-line">
                    {t(
                      `dynamic.settings.sensitive-data.semantic-types.template.${item.item.id.split(".").join("-")}.algorithm.description`
                    )}
                  </div>
                ),
              }}
            </NTooltip>
          )}
          {!isReadonly(item.item) && (
            <MiniActionButton
              onClick={() => {
                state.pendingEditSemanticIndex = row;
              }}
            >
              <PencilIcon class="w-4 h-4" />
            </MiniActionButton>
          )}
        </div>
      );
    },
  });

  if (!props.readonly) {
    // operation.
    columns.push({
      key: "operation",
      title: t("common.edit"),
      width: "minmax(min-content, auto)",
      render: (item, row) => {
        if (isBuiltinSemanticType(item.item)) {
          return (
            <NPopconfirm onPositiveClick={() => emit("remove", row)}>
              {{
                trigger: () => {
                  return (
                    <MiniActionButton>
                      <TrashIcon class="w-4 h-4" />
                    </MiniActionButton>
                  );
                },
                default: () => (
                  <div class="whitespace-nowrap">
                    {t("settings.sensitive-data.semantic-types.table.delete")}
                  </div>
                ),
              }}
            </NPopconfirm>
          );
        }

        return (
          <div>
            {item.mode === "EDIT" && (
              <NPopconfirm onPositiveClick={() => emit("remove", row)}>
                {{
                  trigger: () => {
                    return (
                      <MiniActionButton>
                        <TrashIcon class="w-4 h-4" />
                      </MiniActionButton>
                    );
                  },
                  default: () => (
                    <div class="whitespace-nowrap">
                      {t("settings.sensitive-data.semantic-types.table.delete")}
                    </div>
                  ),
                }}
              </NPopconfirm>
            )}
            {item.mode !== "NORMAL" && (
              <MiniActionButton onClick={() => emit("cancel", row)}>
                <Undo2Icon class="w-4 h-4" />
              </MiniActionButton>
            )}
            {item.mode !== "NORMAL" && (
              <MiniActionButton
                type={"primary"}
                disabled={isConfirmDisabled(item)}
                onClick={() => emit("confirm", row)}
              >
                <CheckIcon class="w-4 h-4" />
              </MiniActionButton>
            )}
            {item.mode === "NORMAL" && (
              <MiniActionButton onClick={() => (item.mode = "EDIT")}>
                <PencilIcon class="w-4 h-4" />
              </MiniActionButton>
            )}
          </div>
        );
      },
    });
  }
  return columns;
});

const rowProps = (item: SemanticItem) => {
  return {
    style: props.rowClickable ? "cursor: pointer;" : "",
    onClick: () => {
      if (!props.rowClickable) {
        return;
      }
      emit("select", item.item.id);
    },
  };
};

const onInput = (index: number, callback: (item: SemanticItem) => void) => {
  const item = props.semanticItemList[index];
  if (!item) {
    return;
  }
  callback(item);
  item.dirty = true;
};

const isConfirmDisabled = (data: SemanticItem): boolean => {
  if (!data.item.title) {
    return true;
  }
  if (data.mode === "EDIT" && !data.dirty) {
    return true;
  }
  return false;
};
</script>
