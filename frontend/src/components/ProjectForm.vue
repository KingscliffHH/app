<script setup lang="ts">
import { computed } from "vue";

import { ascending } from "@/utils/sort";

import type { Project } from "@/types/project";

import BaseDateInput from "@/components/base/BaseDateInput.vue";
import BaseInput from "@/components/base/BaseInput.vue";
import BaseSelect from "@/components/base/BaseSelect.vue";
import BaseToggle from "@/components/base/BaseToggle.vue";

import { useQuery } from "@/hooks/fetch";
import services from "@/services";

const props = defineProps<{
  modelValue: Project;
  error?: Record<string, string>;
  readonly: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const { data: availableUsers } = useQuery({
  queryFn: () => services.users.getAll()
});

const members = computed(() => {
  return (availableUsers.value || [])
    .filter((u) => u.type === "member")
    .map((u) => ({
      label: u.fullName,
      value: u.id
    }))
    .sort((a, b) => ascending(a.label, b.label));
});

const clients = computed(() => {
  return (availableUsers.value || [])
    .filter((u) => u.type === "client" && u.organisation === data.value.client)
    .map((u) => ({
      label: u.fullName,
      value: u.id
    }))
    .sort((a, b) => ascending(a.label, b.label));
});

const limit = 3;
const canAddTeamMember = computed(() => {
  return data.value.team.teamMembers.length < limit;
});
const addTeamMember = () => {
  data.value.team.teamMembers.push({ id: "" });
};

const { data: availableOrgs } = useQuery({
  queryFn: () => services.preferences.getOrganisations()
});

const organisations = computed<string[]>(() => {
  const orgs = availableOrgs.value || [];
  return orgs.sort(ascending);
});

const resetClientRepresentative = () => {
  data.value.clientRepresentative.id = "";
};
</script>

<template>
  <div class="bg-white shadow-md rounded-lg p-6 mb-4">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">Project</span>
      <div class="flex-grow h-0.5 bg-gray-400"></div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <BaseInput
        v-model="data.name"
        label="Name"
        name="name"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />

      <BaseInput
        v-model="data.region"
        label="Region"
        name="region"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />

      <BaseInput
        v-model="data.ciProjectNumber"
        label="C&I Project Number"
        name="ciProjectNumber"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />

      <BaseInput
        v-model="data.clientProjectNumber"
        label="Client Project Number"
        name="clientProjectNumber"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />

      <BaseSelect
        v-model="data.client"
        label="Client (Organization)"
        name="client"
        :errors="props.error"
        :readonly="props.readonly"
        :options="organisations"
        @update:model-value="resetClientRepresentative"
      />

      <BaseSelect
        v-model="data.clientRepresentative.id"
        label="Client Representative"
        name="clientRepresentative"
        :errors="props.error"
        :readonly="props.readonly"
        :options="clients"
      />

      <BaseDateInput
        v-model="data.startDate"
        label="Start Date"
        name="startDate"
        :errors="props.error"
        :readonly="props.readonly"
      />
    </div>
  </div>
  <div class="bg-white shadow-md rounded-lg p-6 mb-4">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">Project Team</span>
      <div class="flex-grow h-0.5 bg-gray-400"></div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <BaseSelect
        v-model="data.team.projectLead.id"
        label="Project Lead"
        name="projectLead"
        :errors="props.error"
        :readonly="props.readonly"
        :options="members"
      />
      <template
        v-for="(_, index) in data.team.teamMembers"
        :key="index"
      >
        <div class="flex">
          <BaseSelect
            v-model="data.team.teamMembers[index].id"
            :label="`Team Member ${index + 1}`"
            :name="`teamMembers${index}`"
            :readonly="props.readonly"
            :options="members"
          />

          <v-btn
            icon
            flat
            @click="data.team.teamMembers.splice(index, 1)"
          >
            <i class="material-icons-round">delete</i>
          </v-btn>
        </div>
      </template>
    </div>
    <button
      class="mt-4 px-4 py-2 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-blue-500"
      :disabled="!canAddTeamMember"
      @click="addTeamMember"
    >
      + Add More
    </button>
  </div>
  <div class="bg-white shadow-md rounded-lg p-6 mb-4">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700"
        >Scope of Engagement</span
      >
      <div class="flex-grow h-0.5 bg-gray-400"></div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <BaseToggle
        v-model="data.scope.quantification"
        label="Quantification"
        name="quantification"
      />
      <BaseToggle
        v-model="data.scope.costEstimation"
        label="Cost Estimation"
        name="costEstimation"
      />
      <BaseToggle
        v-model="data.scope.probabilisticRiskAssessment"
        label="Probabilistic Risk Assessment"
        name="probabilisticRiskAssessment"
      />
      <BaseToggle
        v-model="data.scope.basisOfEstimateReport"
        label="Basis of Estimate Report"
        name="basisOfEstimateReport"
      />
      <BaseDateInput
        v-model="data.scope.estimatedCompletionDate"
        label="Estimated Completion Date"
        name="estimatedCompletionDate"
        :errors="props.error"
        :readonly="props.readonly"
      />

      <BaseSelect
        v-model.number="data.scope.numberOfMilestones"
        label="Number of Milestones"
        name="numberOfMilestones"
        :options="[1, 2, 3, 4, 5, 6, 7]"
      />

      <BaseSelect
        v-model.number="data.scope.remainsAccessibleForNDays"
        label="Remain Accessible for N Days"
        name="remainsAccessibleForNDays"
        :errors="props.error"
        :options="[7, 15, 30, 60, 90]"
      />
    </div>
  </div>
</template>
