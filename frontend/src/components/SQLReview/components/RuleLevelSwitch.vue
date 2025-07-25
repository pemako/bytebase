<template>
  <div class="flex items-center" :class="[!editable && 'pointer-events-none']">
    <button
      v-for="item in availableLevel"
      :key="item.level"
      :disabled="disabled || !editable"
      :class="['button', item.class, level === item.level && 'active']"
      @click="$emit('level-change', item.level)"
    >
      {{ item.title }}
    </button>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { SQLReviewRuleLevel } from "@/types/proto-es/v1/org_policy_service_pb";

const props = withDefaults(
  defineProps<{
    level: SQLReviewRuleLevel;
    disabled?: boolean;
    editable?: boolean;
  }>(),
  {
    disabled: false,
    editable: true,
  }
);

defineEmits<{
  (event: "level-change", level: SQLReviewRuleLevel): void;
}>();

const { t } = useI18n();

const availableLevel = computed(() => {
  return [
    {
      level: SQLReviewRuleLevel.ERROR,
      title: t("sql-review.level.error"),
      class: "error",
    },
    {
      level: SQLReviewRuleLevel.WARNING,
      title: t("sql-review.level.warning"),
      class: "warning",
    },
    {
      level: SQLReviewRuleLevel.DISABLED,
      title: t("sql-review.level.disabled"),
      class: "",
    },
  ].filter((item) => {
    if (props.editable) {
      return item.level !== SQLReviewRuleLevel.DISABLED;
    }
    return props.level === item.level;
  });
});
</script>

<style lang="postcss" scoped>
.button {
  @apply relative py-1 w-[4.5rem] whitespace-nowrap border border-control-border text-control font-medium hover:z-[2];
  @apply disabled:cursor-not-allowed disabled:bg-control-bg disabled:opacity-50;
}
.button:not(:first-child) {
  @apply -ml-px;
}
.button:first-child {
  @apply rounded-l;
}
.button:last-child {
  @apply rounded-r;
}
.button.active {
  @apply z-[1];
}
.button.error.active {
  @apply bg-red-100 text-red-800 border-red-800;
}
.button.warning.active {
  @apply bg-yellow-100 text-yellow-800 border-yellow-800;
}
.button:not(:disabled).error:hover {
  @apply bg-red-100 text-red-800 border-red-800;
}
.button:not(:disabled).warning:hover {
  @apply bg-yellow-100 text-yellow-800 border-yellow-800;
}
</style>
