<script setup lang="ts">
import { computed, onMounted } from "vue";

import type { OptionOutturnCost } from "@/types/project";

import BaseNumberInput from "@/components/base/BaseNumberInput.vue";

const props = defineProps<{
  modelValue: OptionOutturnCost[] | null;
  disabled?: boolean;
}>();

const emits = defineEmits(["update:model-value"]);

const EmptyObject = (): OptionOutturnCost => {
  return {
    p90: 0,
    p50: 0,
    base: 0
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
        Option Outturn Cost:
      </span>
    </div>
    <div
      class="grid grid-cols-1 md:grid-cols-[auto_1fr_1fr_1fr_auto] gap-4 items-center"
    >
      <template
        v-for="(item, i) in data"
        :key="i"
      >
        <div>
          <span
            class="block text-sm font-medium text-gray-700 dark:text-gray-400"
            >Option {{ i + 1 }}</span
          >
        </div>
        <BaseNumberInput
          v-model="item.base"
          label="Base"
          name="base"
          prefix="$"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p50"
          label="P50"
          name="p50"
          prefix="$"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.p90"
          label="P90"
          name="p90"
          prefix="$"
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
