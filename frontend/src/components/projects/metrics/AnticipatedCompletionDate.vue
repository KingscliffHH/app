<script setup lang="ts">
import { computed } from "vue";

import type { AnticipatedCompletionDate } from "@/types/project";

import BaseDateInput from "@/components/base/BaseDateInput.vue";

const props = defineProps<{
  modelValue: AnticipatedCompletionDate;
  disabled?: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">
        Anticipated Completion Date:
      </span>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4">
      <BaseDateInput
        v-model="data.quantification"
        label="Quantification"
        name="quantification"
        :disabled="disabled"
      />
      <BaseDateInput
        v-model="data.costEstimation"
        label="Cost Estimation"
        name="costEstimation"
        :disabled="disabled"
      />
      <BaseDateInput
        v-model="data.probabilisticRiskAssessment"
        label="Probabilistic Risk Assessment"
        name="probabilisticRiskAssessment"
        :disabled="disabled"
      />
      <BaseDateInput
        v-model="data.basisOfEstimateReport"
        label="Basis of Estimate Report"
        name="basisOfEstimateReport"
        :disabled="disabled"
      />
    </div>
  </section>
</template>
