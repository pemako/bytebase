<template>
  <div class="space-y-1 w-full">
    <div>
      <span>{{ $t("common.credentials") }}</span>
      <span class="text-red-600">*</span>
    </div>
    <p class="textinfolabel">
      <span>{{ $t("instance.create-gcp-credentials") }}</span>
      <a
        href="https://docs.bytebase.com/get-started/instance/#create-a-google-cloud-service-account-as-the-credential?source=console"
        target="_blank"
        class="normal-link inline-flex items-center ml-1"
      >
        <span>{{ $t("common.detailed-guide") }}</span>
        <heroicons-outline:external-link class="w-4 h-4 ml-1" />
      </a>
    </p>
    <DroppableTextarea
      :value="value"
      :resizable="false"
      :placeholder="placeholder"
      class="w-full h-24 mt-2 whitespace-pre-wrap"
      @update:value="$emit('update:value', $event ?? '')"
    />
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import DroppableTextarea from "@/components/misc/DroppableTextarea.vue";

const props = withDefaults(
  defineProps<{
    value: string;
    writeOnly?: boolean;
  }>(),
  {
    writeOnly: false,
  }
);

defineEmits<{
  (name: "update:value", value: string): void;
}>();

const { t } = useI18n();

const placeholder = computed(() => {
  return props.writeOnly
    ? t("instance.type-or-paste-credentials-write-only")
    : t("instance.type-or-paste-credentials");
});
</script>
