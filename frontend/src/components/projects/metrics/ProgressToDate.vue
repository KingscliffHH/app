<script setup lang="ts">
import { computed } from "vue";

import type { ProgressToDate, ScopeOfEngagement } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseSelect from "@/components/base/BaseSelect.vue";

const props = defineProps<{
  modelValue: ProgressToDate;
  scope: ScopeOfEngagement;
  maxMilestone: number;
  disabled?: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: ProgressToDate): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const milestoneOptions: number[] = Array.from(
  { length: props.maxMilestone + 1 },
  (_, i) => i
);
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700"
        >Progress To Date -
        <small class="ms-2 font-semibold text-gray-500 dark:text-gray-400"
          >Percentage Completed</small
        >:</span
      >
    </div>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4">
      <BaseInput
        v-if="props.scope.quantification"
        v-model.number="data.quantification"
        class="rtl"
        label="Quantification"
        name="quantification"
        type="number"
        min="0"
        max="100"
        suffix="%"
        :disabled="disabled"
      />
      <BaseInput
        v-if="props.scope.costEstimation"
        v-model.number="data.costEstimation"
        class="rtl"
        label="Cost Estimation"
        name="costEstimation"
        type="number"
        min="0"
        max="100"
        suffix="%"
        :disabled="disabled"
      />
      <BaseInput
        v-if="props.scope.probabilisticRiskAssessment"
        v-model.number="data.probabilisticRiskAssessment"
        class="rtl"
        label="Probabilistic Risk Assessment"
        name="probabilisticRiskAssessment"
        type="number"
        min="0"
        max="100"
        suffix="%"
        :disabled="disabled"
      />
      <BaseInput
        v-if="props.scope.basisOfEstimateReport"
        v-model.number="data.basisOfEstimateReport"
        class="rtl"
        label="Basis of Estimate Report"
        name="basisOfEstimateReport"
        type="number"
        min="0"
        max="100"
        suffix="%"
        :disabled="disabled"
      />
      <BaseSelect
        v-model.number="data.numberOfMilestones"
        label="Number of Milestones"
        name="numberOfMilestones"
        :disabled="disabled"
        :options="milestoneOptions"
      />
    </div>
  </section>
</template>
