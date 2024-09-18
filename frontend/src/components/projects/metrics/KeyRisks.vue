<script setup lang="ts">
import { computed, onMounted } from "vue";

import type { KeyRisk } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseSelect from "@/components/base/BaseSelect.vue";

const props = defineProps<{
  modelValue: KeyRisk[] | null;
  disabled?: boolean;
}>();

const emits = defineEmits(["update:model-value"]);

const EmptyObject = (): KeyRisk => {
  return {
    description: "",
    score: "",
    trend: ""
  };
};

const limit = 5;
const data = computed({
  get: () => props.modelValue || [],
  set: (value) => emits("update:model-value", value)
});

const canAddMore = computed(() => data.value.length < limit);

const addAndFocus = () => {
  data.value = [
    ...data.value,
    {
      description: "",
      score: "",
      trend: ""
    }
  ];
};

const removeItem = (index: number) => {
  console.info("remove item", index);
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
      <span class="pr-4 text-lg font-medium text-gray-700"> Key Risks: </span>
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
            >Risk {{ i + 1 }}</span
          >
        </div>
        <BaseInput
          v-model="item.description"
          label="Description"
          name="description"
          :disabled="disabled"
        />
        <BaseSelect
          v-model="item.score"
          label="Score"
          name="score"
          placeholder="High, Medium, Low"
          :options="[
            { label: 'High', value: 'high' },
            { label: 'Medium', value: 'medium' },
            { label: 'Low', value: 'low' }
          ]"
          :disabled="disabled"
        />
        <BaseSelect
          v-model="item.trend"
          label="Trend"
          name="trend"
          :options="[
            { label: 'Orange', value: 'orange' },
            { label: 'Yellow', value: 'yellow' },
            { label: 'Green', value: 'green' }
          ]"
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
      @click="addAndFocus"
    >
      + Add More
    </button>
  </section>
</template>
