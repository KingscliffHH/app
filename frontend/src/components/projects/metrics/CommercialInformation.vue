<script setup lang="ts">
import { computed } from "vue";

import type { CommercialInformation } from "@/types/project";

import BaseNumberInput from "@/components/base/BaseNumberInput.vue";

const props = defineProps<{
  modelValue: CommercialInformation;
  disabled?: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const remainingFee = computed(() => {
  return (
    data.value.ciContractedValue +
    data.value.approvedVariationToDate -
    data.value.ciAccrualToDate
  );
});
</script>

<template>
  <section>
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700"
        >Commercial Information:</span
      >
    </div>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4">
      <BaseNumberInput
        v-model="data.ciContractedValue"
        label="C&I Contracted Value"
        name="ci_contracted_value"
        :min="0"
        prefix="$"
        :disabled="disabled"
      />
      <BaseNumberInput
        v-model="data.approvedVariationToDate"
        label="Approved Variation to Date"
        name="approved_variation_to_date"
        :min="0"
        prefix="$"
        :disabled="disabled"
      />
      <BaseNumberInput
        v-model="data.ciAccrualToDate"
        label="Accrual to Date"
        name="accrual_to_date"
        :min="0"
        prefix="$"
        :disabled="disabled"
      />
      <BaseNumberInput
        v-model="remainingFee"
        label="Remaining Fee"
        name="remaining_fee"
        readonly
        prefix="$"
        :disabled="disabled"
      />
    </div>
  </section>
</template>
