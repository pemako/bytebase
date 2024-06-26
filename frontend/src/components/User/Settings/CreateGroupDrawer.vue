<template>
  <Drawer @close="$emit('close')">
    <DrawerContent
      class="w-[50rem] max-w-[90vw] relative"
      :title="
        isCreating
          ? $t('settings.members.groups.add-group')
          : $t('settings.members.groups.update-group')
      "
    >
      <template #default>
        <div class="flex flex-col gap-y-4">
          <div class="flex flex-col gap-y-2">
            <div>
              <label class="textlabel block">
                {{ $t("settings.members.groups.form.email") }}
                <span class="text-red-600">*</span>
              </label>
              <span class="textinfolabel">
                {{ $t("settings.members.groups.form.email-tips") }}
              </span>
            </div>
            <NInput
              v-model:value="state.group.email"
              :disabled="!isCreating || !allowEdit"
              placeholder="The unique email"
            />
          </div>
          <div class="flex flex-col gap-y-2">
            <label class="textlabel block">
              {{ $t("settings.members.groups.form.title") }}
              <span class="text-red-600">*</span>
            </label>
            <NInput v-model:value="state.group.title" :disabled="!allowEdit" />
          </div>
          <div class="flex flex-col gap-y-2">
            <label class="textlabel block">
              {{ $t("settings.members.groups.form.description") }}
            </label>
            <NInput
              v-model:value="state.group.description"
              :disabled="!allowEdit"
            />
          </div>
          <div class="flex flex-col gap-y-2">
            <label class="textlabel block">
              {{ $t("settings.members.groups.form.members") }}
              <span class="text-red-600">*</span>
            </label>
            <div class="flex flex-col space-y-2">
              <div
                v-for="(member, i) in state.group.members"
                :key="member.member"
                class="w-full flex items-center space-x-3"
              >
                <UserSelect
                  :user="getUserUidForMember(member)"
                  :multiple="false"
                  :size="'medium'"
                  :include-all="false"
                  @update:user="(uid: string) => updateMemberEmail(i, uid)"
                />
                <GroupMemberRoleSelect
                  :value="member.role"
                  :size="'medium'"
                  @update:value="
                    (role: UserGroupMember_Role) => updateMemberRole(i, role)
                  "
                />
                <div class="pl-5 flex justify-end">
                  <NButton
                    quaternary
                    circle
                    size="tiny"
                    @click="deleteMember(i)"
                  >
                    <template #icon>
                      <Trash2Icon class="w-4 h-auto" />
                    </template>
                  </NButton>
                </div>
              </div>
            </div>
            <div>
              <NButton :disabled="!allowEdit" @click="addMember">
                {{ $t("settings.members.groups.form.add-member") }}
              </NButton>
            </div>
          </div>
        </div>
      </template>
      <template #footer>
        <div class="w-full flex justify-between items-center">
          <div>
            <NPopconfirm
              v-if="!isCreating"
              :disabled="!allowEdit"
              @positive-click="handleDeleteGroup"
            >
              <template #trigger>
                <NButton
                  quaternary
                  size="small"
                  :disabled="!allowEdit"
                  @click.stop
                >
                  <template #icon>
                    <Trash2Icon class="w-4 h-auto" />
                  </template>
                  <template #default>
                    {{ $t("common.delete") }}
                  </template>
                </NButton>
              </template>

              <template #default>
                <div>
                  {{ $t("settings.members.action.delete-confirm-title") }}
                </div>
              </template>
            </NPopconfirm>
          </div>

          <div class="flex flex-row items-center justify-end gap-x-3">
            <NButton @click="$emit('close')">
              {{ $t("common.cancel") }}
            </NButton>
            <NTooltip :disabled="!errorMessage">
              <template #trigger>
                <NButton
                  type="primary"
                  :disabled="!allowEdit || !allowConfirm"
                  :loading="state.isRequesting"
                  @click="tryCreateOrUpdateGroup"
                >
                  {{ $t("common.confirm") }}
                </NButton>
              </template>
              <span class="w-56 text-sm">
                {{ errorMessage }}
              </span>
            </NTooltip>
          </div>
        </div>
      </template>
    </DrawerContent>
  </Drawer>
</template>

<script lang="ts" setup>
import { cloneDeep, isEqual } from "lodash-es";
import { Trash2Icon } from "lucide-vue-next";
import { NPopconfirm, NButton, NInput, NTooltip } from "naive-ui";
import { computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import {
  extractGroupEmail,
  useUserGroupStore,
  useCurrentUserV1,
  pushNotification,
  useUserStore,
} from "@/store";
import { userNamePrefix, userGroupNamePrefix } from "@/store/modules/v1/common";
import { getUserEmailFromIdentifier } from "@/store/modules/v1/common";
import {
  UserGroup,
  UserGroupMember,
  UserGroupMember_Role,
} from "@/types/proto/v1/user_group";
import {
  isValidEmail,
  extractUserUID,
  hasWorkspacePermissionV2,
} from "@/utils";

interface LocalState {
  isRequesting: boolean;
  group: {
    email: string;
    title: string;
    description: string;
    members: UserGroupMember[];
  };
}

const props = defineProps<{
  group?: UserGroup;
}>();

const emit = defineEmits<{
  (event: "close"): void;
}>();

const { t } = useI18n();
const userStore = useUserStore();
const groupStore = useUserGroupStore();
const currentUserV1 = useCurrentUserV1();

const state = reactive<LocalState>({
  isRequesting: false,
  group: {
    email: extractGroupEmail(props.group?.name ?? ""),
    title: props.group?.title ?? "",
    description: props.group?.description ?? "",
    members: cloneDeep(
      props.group?.members ?? [
        UserGroupMember.fromPartial({
          role: UserGroupMember_Role.OWNER,
          member: `${userNamePrefix}${currentUserV1.value.email}`,
        }),
      ]
    ),
  },
});

const isCreating = computed(() => !props.group);

const selfMemberInGroup = computed(() => {
  return props.group?.members.find(
    (member) =>
      getUserEmailFromIdentifier(member.member) === currentUserV1.value.email
  );
});

const allowEdit = computed(() => {
  if (selfMemberInGroup.value?.role === UserGroupMember_Role.OWNER) {
    return true;
  }
  return hasWorkspacePermissionV2(
    currentUserV1.value,
    isCreating.value ? "bb.userGroups.create" : "bb.userGroups.update"
  );
});

const validGroup = computed(() => {
  const memberMap = new Map<string, UserGroupMember>();
  for (const member of state.group.members) {
    if (!member.member) {
      continue;
    }
    if (
      !memberMap.has(member.member) ||
      member.role === UserGroupMember_Role.OWNER
    ) {
      memberMap.set(member.member, member);
    }
  }
  return UserGroup.fromPartial({
    name: `${userGroupNamePrefix}${state.group.email}`,
    title: state.group.title,
    creator: props.group?.creator ?? "",
    description: state.group.description,
    members: [...memberMap.values()],
  });
});

const errorMessage = computed(() => {
  if (!isValidEmail(state.group.email)) {
    return "Invalid group email";
  }
  if (!state.group.title) {
    return "Title is required";
  }
  if (!validGroup.value.members.length) {
    return "At least select 1 member for the group";
  }
  if (
    !validGroup.value.members.some(
      (member) => member.role === UserGroupMember_Role.OWNER
    )
  ) {
    return "At least has 1 owner in the group";
  }

  return "";
});

const allowConfirm = computed(() => {
  if (errorMessage.value) {
    return false;
  }

  return !isEqual(props.group, validGroup.value);
});

const addMember = () => {
  const member = UserGroupMember.fromPartial({
    role:
      state.group.members.length === 0
        ? UserGroupMember_Role.OWNER
        : UserGroupMember_Role.MEMBER,
  });
  state.group.members.push(member);
};

const getUserUidForMember = (member: UserGroupMember) => {
  if (!member.member) {
    return;
  }
  const email = getUserEmailFromIdentifier(member.member);
  const user = userStore.getUserByEmail(email);
  if (!user) {
    return;
  }
  return extractUserUID(user.name);
};

const updateMemberEmail = (index: number, uid: string) => {
  const user = userStore.getUserById(uid);
  if (!user) {
    return;
  }
  state.group.members[index] = {
    ...state.group.members[index],
    member: `${userNamePrefix}${user.email}`,
  };
};

const updateMemberRole = (index: number, role: UserGroupMember_Role) => {
  state.group.members[index] = {
    ...state.group.members[index],
    role,
  };
};

const deleteMember = (index: number) => {
  state.group.members.splice(index, 1);
};

const handleDeleteGroup = async () => {
  if (!props.group) {
    return;
  }
  await groupStore.deleteGroup(props.group.name);
  pushNotification({
    module: "bytebase",
    style: "INFO",
    title: t("common.deleted"),
  });
  emit("close");
};

const tryCreateOrUpdateGroup = async () => {
  state.isRequesting = true;

  try {
    if (isCreating.value) {
      await groupStore.createGroup(validGroup.value);
      pushNotification({
        module: "bytebase",
        style: "SUCCESS",
        title: t("common.created"),
      });
    } else {
      await groupStore.updateGroup(validGroup.value, [
        "title",
        "description",
        "members",
      ]);
      pushNotification({
        module: "bytebase",
        style: "SUCCESS",
        title: t("common.updated"),
      });
    }
    emit("close");
  } finally {
    state.isRequesting = false;
  }
};
</script>
