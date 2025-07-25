<template>
  <div class="flex flex-col space-y-4">
    <div class="flex justify-between items-end">
      <TabFilter
        :value="state.selectedTab"
        :items="tabItemList"
        @update:value="(val) => (state.selectedTab = val as LocalTabType)"
      />

      <SearchBox
        v-model:value="state.searchText"
        :placeholder="$t('common.filter-by-name')"
      />
    </div>
    <div class="">
      <PagedProjectTable
        v-if="state.selectedTab == 'PROJECT'"
        session-key="bb.project-table.archived"
        :filter="{
          query: state.searchText,
          state: State.DELETED,
          excludeDefault: true,
        }"
        :bordered="true"
        :show-selection="false"
      />
      <PagedInstanceTable
        v-else-if="state.selectedTab == 'INSTANCE'"
        session-key="bb.instance-table.archived"
        :bordered="true"
        :show-selection="false"
        :filter="{
          query: state.searchText,
          state: State.DELETED,
        }"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import {
  SearchBox,
  TabFilter,
  PagedProjectTable,
  PagedInstanceTable,
} from "@/components/v2";
import { State } from "@/types/proto-es/v1/common_pb";
import { hasWorkspacePermissionV2 } from "@/utils";

type LocalTabType = "PROJECT" | "INSTANCE";

interface LocalState {
  selectedTab: LocalTabType;
  searchText: string;
}

const { t } = useI18n();
const state = reactive<LocalState>({
  selectedTab: "PROJECT",
  searchText: "",
});

const tabItemList = computed(() => {
  const list: { value: LocalTabType; label: string }[] = [
    { value: "PROJECT", label: t("common.project") },
  ];

  if (hasWorkspacePermissionV2("bb.instances.undelete")) {
    list.push({ value: "INSTANCE", label: t("common.instance") });
  }

  return list;
});
</script>
