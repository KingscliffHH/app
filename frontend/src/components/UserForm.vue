<script setup lang="ts">
import { computed } from "vue";

import { ascending } from "@/utils/sort";

import { type User } from "@/types/user";

import { useQuery } from "@/hooks/fetch";
import services from "@/services";

import BaseImagePreviewInput from "./base/BaseImagePreviewInput.vue";
import BaseInput from "./base/BaseInput.vue";
import BaseSelect from "./base/BaseSelect.vue";

const props = defineProps<{
  modelValue: User;
  image?: File | null;
  error?: Record<string, string>;
  readonly: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
  (e: "update:image", value: any): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const image = computed({
  get: () => props.image,
  set: (value) => emits("update:image", value)
});

const { data: availableOrgs } = useQuery({
  queryFn: () => services.preferences.getOrganisations()
});

const organisations = computed<string[]>(() => {
  const orgs = availableOrgs.value || [];
  return orgs.sort(ascending);
});
</script>

<template>
  <form
    action="javascript:void(0);"
    autocomplete="off"
  >
    <div class="bg-white shadow-md rounded-lg p-6 mb-4">
      <div class="grid grid-cols-1 w-1/2 gap-4">
        <!-- User Type -->
        <BaseSelect
          v-model="data.type"
          label="User Type"
          name="type"
          :options="[
            { label: 'C&I', value: 'member' },
            { label: 'Client', value: 'client' }
          ]"
          :errors="props.error"
          :disabled="data.id !== ''"
        />
        <BaseInput
          v-model="data.fullName"
          label="Full Name"
          name="fullName"
          type="text"
          :errors="props.error"
          :readonly="props.readonly"
        />
        <BaseInput
          v-model="data.email"
          label="Email"
          name="email"
          type="email"
          :errors="props.error"
          :readonly="props.readonly"
        />

        <!-- Avatar -->
        <BaseImagePreviewInput
          :url="data.avatar"
          :readonly="props.readonly"
          @update:image="(value) => (image = value)"
        />

        <!-- Bio -->
        <BaseInput
          v-if="data.type === 'member'"
          v-model="data.bio"
          label="Bio"
          name="bio"
          type="text"
          :errors="props.error"
          :readonly="props.readonly"
        />
        <!-- Organisation -->
        <v-combobox
          v-if="data.type === 'client'"
          v-model="data.organisation"
          label="Organisation"
          name="organisation"
          variant="outlined"
          clearable
          :items="organisations"
          :errors="props.error"
          :error-messages="
            (!!props.error && props.error!['organisation']) || ''
          "
          :error="!!props.error && 'organisation' in props.error"
          :readonly="props.readonly"
        ></v-combobox>
        <!-- Role -->
        <BaseInput
          v-if="data.type === 'client'"
          v-model="data.clientRole"
          label="Role"
          name="role"
          type="text"
          :errors="props.error"
          :readonly="props.readonly"
        />
      </div>
    </div>
  </form>
</template>
