<template>
  <div class="flex justify-start items-center">
    <NPopconfirm @positive-click="$emit('drop')">
      <template #trigger>
        <MiniActionButton tag="div">
          <TrashIcon class="w-4 h-4" />
        </MiniActionButton>
      </template>
      <template #default>
        {{ $t("branch.index.drop-index-confirm") }}
      </template>
    </NPopconfirm>
  </div>
</template>
<script lang="ts" setup>
import { TrashIcon } from "lucide-vue-next";
import { NPopconfirm } from "naive-ui";
import { MiniActionButton } from "@/components/v2";
import type { IndexMetadata } from "@/types/proto-es/v1/database_service_pb";

defineProps<{
  index: IndexMetadata;
}>();
defineEmits<{
  (event: "drop"): void;
}>();
</script>
