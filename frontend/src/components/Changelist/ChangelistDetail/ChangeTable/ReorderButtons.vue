<template>
  <NButton
    size="tiny"
    style="--n-padding: 0 6px"
    :disabled="!allowMoveUp"
    @click.stop="$emit('move', -1)"
  >
    <template #icon>
      <heroicons:arrow-up />
    </template>
  </NButton>
  <NButton
    size="tiny"
    style="--n-padding: 0 6px"
    :disabled="!allowMoveDown"
    @click.stop="$emit('move', 1)"
  >
    <template #icon>
      <heroicons:arrow-down />
    </template>
  </NButton>
</template>

<script lang="ts" setup>
import { NButton } from "naive-ui";
import { computed } from "vue";
import type { Changelist_Change as Change } from "@/types/proto-es/v1/changelist_service_pb";

const props = defineProps<{
  changes: Change[];
  row: number;
}>();
defineEmits<{
  (event: "move", delta: -1 | 1): void;
}>();

const allowMoveUp = computed(() => {
  return props.row > 0;
});
const allowMoveDown = computed(() => {
  return props.row < props.changes.length - 1;
});
</script>
