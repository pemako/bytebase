<template>
  <NInput
    :value="value"
    :style="style"
    v-bind="$attrs"
    @focus="focused = true"
    @blur="focused = false"
    @update:value="$emit('update:value', $event)"
  />
</template>

<script setup lang="ts">
import { NInput } from "naive-ui";
import type { CSSProperties } from "vue";
import { computed, ref } from "vue";

defineOptions({
  inheritAttrs: false,
});

defineProps<{
  value: string | undefined | null;
}>();
defineEmits<{
  (event: "update:value", value: string): void;
}>();

const focused = ref(false);

const style = computed(() => {
  const style: CSSProperties = {
    cursor: "default",
    "--n-color": "transparent",
    "--n-color-disabled": "transparent",
  };
  const border = focused.value
    ? "1px solid rgb(var(--color-control-border))"
    : "none";
  style["--n-border"] = border;
  style["--n-border-disabled"] = border;

  return style;
});
</script>
