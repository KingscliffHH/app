<script setup lang="ts">
import { computed, onMounted } from "vue";

import type { TotalProjectCostPerMilestone } from "@/types/project";

import BaseDateInput from "@/components/base/BaseDateInput.vue";
import BaseInput from "@/components/base/BaseInput.vue";
import BaseNumberInput from "@/components/base/BaseNumberInput.vue";

const props = defineProps<{
  modelValue: TotalProjectCostPerMilestone[] | null;
  disabled?: boolean;
}>();

const emits = defineEmits(["update:model-value"]);

const EmptyObject = (): TotalProjectCostPerMilestone => {
  return {
    date: new Date(),
    levelOfDesign: "",
    baseValue: 0,
    p90OutturnCost: 0,
    p90RiskContingency: 0,
    p50OutturnCost: 0,
    p50RiskContingency: 0,
    currentMilstone: false
  };
};

const limit = 10;
const data = computed({
  get: () => props.modelValue || [],
  set: (value) => emits("update:model-value", value)
});
const canAddMore = computed(() => data.value.length < limit);

const addMore = () => {
  data.value = [...data.value, EmptyObject()];
};

const removeItem = (index: number) => {
  data.value = data.value.filter((_, i) => i !== index);
};

const selectedMilestone = computed({
  get: () => {
    return data.value.map((x) => x.currentMilstone).indexOf(true);
  },
  set: (val) => {
    data.value.forEach((x) => {
      x.currentMilstone = false;
    });

    data.value[val].currentMilstone = true;
  }
});

onMounted(() => {
  if (!props.modelValue) {
    data.value = Array.from({ length: 3 }, () => EmptyObject());
  }
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">
        Total Project Cost -
        <small class="ms-2 font-semibold text-gray-500 dark:text-gray-400">
          (Percentage Completed)
        </small>
        :
      </span>
    </div>

    <div
      style="display: flex; align-items: center; margin-bottom: 1rem; gap: 10px"
    >
      <span>Milestones</span>
    </div>

    <div
      class="grid grid-cols-1 md:grid-cols-[auto_auto_1fr_1fr_1fr_1fr_1fr_1fr_1fr_auto] gap-4 items-center"
    >
      <template
        v-for="(item, i) in data"
        :key="i"
      >
        <input
          v-model="selectedMilestone"
          class="h-5 w-5"
          :value="i"
          type="radio"
          name="current-milestone"
        />

        <div>
          <span
            class="block text-sm font-medium text-gray-700 dark:text-gray-400"
          >
            #{{ i + 1 }}
          </span>
        </div>

        <BaseDateInput
          v-model="item.date"
          label="Date"
          name="date"
          type="date"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.baseValue"
          label="Base Value"
          name="baseValue"
          prefix="$"
          :precision="2"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p50OutturnCost"
          label="P50 Outturn Cost"
          name="p50OutturnCost"
          prefix="$"
          :precision="2"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p50RiskContingency"
          class="rtl"
          label="P50 Risk Contingency"
          name="p50RiskContingency"
          :min="0"
          :max="100"
          suffix="%"
          :precision="2"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p90OutturnCost"
          label="P90 Outturn Cost"
          name="p90OutturnCost"
          prefix="$"
          :precision="2"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p90RiskContingency"
          class="rtl"
          label="P90 Risk Contingency"
          name="p90RiskContingency"
          :min="0"
          :max="100"
          suffix="%"
          :precision="2"
          :disabled="disabled"
        />
        <BaseInput
          v-model="item.levelOfDesign"
          label="Level Of Design"
          name="levelOfDesign"
          :disabled="disabled"
        />

        <v-btn
          class="text-red-500"
          variant="outlined"
          icon="mdi-delete"
          :disabled="disabled"
          @click="removeItem(i)"
        />
      </template>
    </div>
    <button
      class="mt-4 px-4 py-2 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-blue-500"
      :disabled="!canAddMore || disabled"
      @click="addMore"
    >
      + Add More
    </button>
  </section>
</template>
