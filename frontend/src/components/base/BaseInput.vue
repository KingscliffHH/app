<script setup lang="ts">
import { computed, useAttrs } from "vue";

defineOptions({
  inheritAttrs: false
});

const props = defineProps<{
  modelValue: any;
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

// const id = `${getCurrentInstance()!.uid}`;

const hasError = computed(() => {
  return !!props.errors && props.name in props.errors;
});

const hint = computed(() => {
  return hasError.value ? props.errors![props.name] : "";
});
</script>

<template>
  <v-text-field
    v-model="value"
    :name="props.name"
    variant="outlined"
    hide-details="auto"
    v-bind="attrs"
    :error-messages="hint"
    :error="hasError"
  ></v-text-field>
</template>

<style>
.rtl input {
  text-align: right;
}
</style>
