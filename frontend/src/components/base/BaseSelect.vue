<script setup lang="ts">
import { computed, useAttrs } from "vue";

defineOptions({
  inheritAttrs: false
});

const props = defineProps<{
  modelValue?: any;
  options: number[] | string[] | { label: string; value: any }[];
  name: string;
  errors?: Record<string, string>;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const attrs = useAttrs();

const value = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const hasError = computed(() => {
  return !!props.errors && props.name in props.errors;
});

const hint = computed(() => {
  return hasError.value ? props.errors![props.name] : "";
});
</script>

<template>
  <v-select
    v-model="value"
    item-title="label"
    item-value="value"
    :items="options"
    :name="$props.name"
    variant="outlined"
    :error="hasError"
    :error-messages="hint"
    v-bind="attrs"
    hide-details="auto"
  ></v-select>
  <!-- <div :class="hasError ? 'text-red-500' : 'text-gray-700'">
    <label
      v-if="props.label"
      :for="id"
      class="block text-sm font-medium"
    >
      {{ props.label }}
    </label>
    <select
      :id="id"
      v-bind="attrs"
      v-model="value"
      :name="props.name"
      class="mt-1 block w-full h-11 border border-gray-300 bg-inherit rounded-md shadow-sm px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 disabled:cursor-not-allowed disabled:bg-gray-100"
      :class="{ 'border-none': borderless }"
    >
      <option
        v-for="option in displayOptions"
        :key="option.value"
        :value="option.value"
      >
        {{ option.label }}
      </option>
    </select>
    <p
      v-if="hint"
      class="text-sm font-bold opacity-50"
    >
      {{ hint }}
    </p>
  </div> -->
</template>
