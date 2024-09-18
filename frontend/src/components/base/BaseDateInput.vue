<script setup lang="ts">
import { computed, useAttrs, ref } from "vue";

defineOptions({
  inheritAttrs: false
});

const props = defineProps<{
  modelValue: Date;
  name: string;
  errors?: Record<string, string>;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const attrs = useAttrs();

const shouldOpenMenu = ref<boolean>(false);

const formatDate = (date: Date) => {
  const offset = date.getTimezoneOffset();
  const localDate = new Date(date.getTime() - offset * 60 * 1000);
  return localDate.toISOString().slice(0, 10);
};

const value = computed({
  get: () => (props.modelValue ? formatDate(new Date(props.modelValue)) : null),
  set: (value) => emits("update:modelValue", value ? new Date(value) : null)
});

const date = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const hasError = computed(() => {
  return !!props.errors && props.name in props.errors;
});

const hint = computed(() => {
  return hasError.value ? props.errors![props.name] : "";
});

const onCloseMenu = (): void => {
  shouldOpenMenu.value = false;
};
</script>

<template>
  <v-text-field
    v-model="value"
    :name="props.name"
    variant="outlined"
    v-bind="attrs"
    :error-messages="hint"
    :error="hasError"
    type="date"
    hide-details="auto"
  >
    <template #append-inner>
      <v-menu
        v-model="shouldOpenMenu"
        :close-on-content-click="false"
      >
        <template #activator="{ props: mProps }">
          <v-btn
            flat
            size="small"
            icon="mdi-calendar"
            v-bind="mProps"
          />
        </template>
        <v-card>
          <v-card-text class="pa-0">
            <v-date-picker
              v-model="date"
              color="primary"
              @update:model-value="onCloseMenu"
            ></v-date-picker>
          </v-card-text>
        </v-card>
      </v-menu>
    </template>
  </v-text-field>
  <!-- <div>
    <label
      :for="id"
      class="block text-sm font-medium text-gray-700 dark:text-gray-400"
      >{{ props.label }}</label
    >
    <input
      :id="id"
      v-model="value"
      type="date"
      :name="props.name"
      class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 read-only:cursor-not-allowed read-only:bg-gray-100"
    />
  </div> -->
</template>

<style>
input[type="date"]::-webkit-inner-spin-button,
input[type="date"]::-webkit-calendar-picker-indicator {
  display: none;
  -webkit-appearance: none;
}
</style>
