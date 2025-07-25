<template>
  <NPopconfirm
    v-bind="popconfirmProps"
    :disabled="!tooltip"
    @positive-click="handleConfirm"
  >
    <template #trigger>
      <NButton
        :loading="loading"
        style="--n-icon-margin: 4px; --n-icon-size: 14px"
        v-bind="buttonAttrs"
      >
        <slot />
      </NButton>
    </template>

    <div v-if="tooltip" :class="tooltipClass">{{ tooltip }}</div>
  </NPopconfirm>
</template>

<script lang="ts">
import { defineComponent } from "vue";

defineComponent({
  inheritAttrs: false,
});
</script>

<script lang="ts" setup>
import { omit } from "lodash-es";
import {
  type ButtonProps,
  type PopconfirmProps,
  NButton,
  NPopconfirm,
} from "naive-ui";
import { computed, ref, useAttrs } from "vue";
import type { VueClass } from "@/utils";

export interface SpinnerButtonProps extends /* @vue-ignore */ ButtonProps {
  onConfirm: () => Promise<any>;
  tooltip?: string;
  tooltipClass?: VueClass;
  popconfirmProps?: PopconfirmProps;
}
const props = defineProps<SpinnerButtonProps>();

const attrs = useAttrs();
const buttonAttrs = computed(() =>
  omit(attrs, "tooltip", "tooltipClass", "popconfirmProps")
);

const loading = ref(false);

const handleConfirm = async () => {
  if (loading.value) return;

  loading.value = true;
  try {
    await props.onConfirm();
  } finally {
    loading.value = false;
  }
};
</script>
