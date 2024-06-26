<template>
  <Suspense>
    <template #default>
      <ProvideInstanceContext :instance-id="instanceId">
        <template v-if="hasPermission">
          <div
            v-if="instance.name === UNKNOWN_INSTANCE_NAME"
            class="flex items-center gap-x-2 m-4"
          >
            <BBSpin :size="20" />
            Loading instance...
          </div>
          <router-view v-else :instance-id="instanceId" />
        </template>
        <NoPermissionPlaceholder v-else />
      </ProvideInstanceContext>
    </template>
    <template #fallback>
      <span>Loading instance...</span>
    </template>
  </Suspense>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useRouter } from "vue-router";
import ProvideInstanceContext from "@/components/ProvideInstanceContext.vue";
import { useCurrentUserV1, useInstanceV1Store } from "@/store";
import { instanceNamePrefix } from "@/store/modules/v1/common";
import { UNKNOWN_INSTANCE_NAME } from "@/types";
import { hasWorkspacePermissionV2 } from "@/utils";

const props = defineProps<{
  instanceId: string;
}>();

const router = useRouter();
const currentUser = useCurrentUserV1();
const instanceStore = useInstanceV1Store();

const requiredPermissions = computed(() => {
  const getPermissionListFunc =
    router.currentRoute.value.meta.requiredWorkspacePermissionList;
  return getPermissionListFunc ? getPermissionListFunc() : [];
});

const hasPermission = computed(() => {
  return requiredPermissions.value.every((permission) =>
    hasWorkspacePermissionV2(currentUser.value, permission)
  );
});

const instance = computed(() =>
  instanceStore.getInstanceByName(`${instanceNamePrefix}${props.instanceId}`)
);
</script>
