<script setup lang="ts">
const props = defineProps<{
  icon?: "delete" | "edit" | "add";
  label?: string;
  size?: "sm" | "md" | "lg";
  rounded?: boolean;
}>();
</script>

<template>
  <button
    class="px-4 py-1 text-blue-700 font-semibold border border-blue-500 rounded hover:bg-blue-500 hover:text-white hover:border-transparent"
    :class="{
      'text-xs': props.size === 'sm',
      'text-sm': props.size === 'md',
      'text-lg': props.size === 'lg',
      'py-2': props.size === 'lg',
      'py-1': props.size === 'md',
      'py-0.5': props.size === 'sm'
    }"
  >
    <i class="text-base material-icons-round">{{ props.icon }}</i>
    <span>{{ props.label }}</span>
  </button>
</template>
