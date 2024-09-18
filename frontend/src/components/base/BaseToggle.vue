<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  modelValue: any;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const value = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});
</script>

<template>
  <v-switch
    v-model="value"
    color="primary"
    label="on"
    hide-details
    density="compact"
  ></v-switch>
</template>
