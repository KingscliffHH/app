<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  modelValue: any;
  labelTrue: string;
  labelFalse: string;
  name: string;
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
  <v-btn-toggle
    v-model="value"
    density="compact"
    rounded="xl"
    variant="outlined"
    color="primary"
    group
    mandatory
  >
    <v-btn :value="true">{{ props.labelTrue }}</v-btn>
    <v-btn :value="false">{{ props.labelFalse }}</v-btn>
  </v-btn-toggle>
</template>
