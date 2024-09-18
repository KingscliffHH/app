<script setup lang="ts">
import { computed, onMounted } from "vue";

import type { KeyCostDriver } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseNumberInput from "@/components/base/BaseNumberInput.vue";

const props = defineProps<{
  drivers: KeyCostDriver[] | null;
  baseValue: number;
  disabled?: boolean;
}>();

const emits = defineEmits<{
  (e: "update:drivers", value: KeyCostDriver[]): void;
  (e: "update:baseValue", value: number): void;
}>();

const EmptyObject = (): KeyCostDriver => {
  return {
    driver: "",
    cost: null as any
  };
};
const limit = 5;

const drivers = computed({
  get: () => props.drivers || Array.from({ length: 3 }, EmptyObject),
  set: (value) => emits("update:drivers", value)
});

const baseValue = computed<number>({
  get: () => props.baseValue,
  set: (value) => emits("update:baseValue", value)
});

const canAddMore = computed(() => drivers.value.length < limit);

const addMore = () => {
  drivers.value = [...drivers.value, EmptyObject()];
};

const removeItem = (index: number) => {
  console.info("remove item", index);
  drivers.value = drivers.value.filter((_, i) => i !== index);
};

onMounted(() => {
  if (!props.drivers) {
    drivers.value = Array.from({ length: 3 }, EmptyObject);
  }
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">
        Key Cost Drivers:
      </span>
    </div>
    <!-- Input for Base Value -->
    <div class="mb-4 w-1/4">
      <BaseNumberInput
        v-model="baseValue"
        label="Base Value"
        name="baseValue"
        prefix="$"
        :disabled="disabled"
      />
    </div>
    <div
      class="grid grid-cols-1 md:grid-cols-[auto_1fr_1fr_auto] gap-4 items-center"
    >
      <template
        v-for="(item, i) in drivers"
        :key="i"
      >
        <div>
          <span
            class="block text-sm font-medium text-gray-700 dark:text-gray-400"
            >Cost Driver {{ i + 1 }}</span
          >
        </div>
        <BaseInput
          v-model="item.driver"
          label="Driver"
          name="driver"
          :disabled="disabled"
        />
        <BaseNumberInput
          v-model="item.cost"
          label="Cost"
          name="cost"
          :min="0"
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
