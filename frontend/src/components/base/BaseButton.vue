<script setup lang="ts">
import { computed } from "vue";

type Icon = "delete" | "edit" | "add";
type ButtonColor = "standard" | "primary" | "secondary" | "delete";

const props = defineProps<{
  label?: string;
  icon: Icon;
  rounded?: boolean;
  primary?: boolean;
  color?: ButtonColor;
}>();

const class_ = computed(() => {
  return {
    "rounded-full": props.rounded,
    "bg-blue-500 text-white": props.color === "primary",
    "bg-gray-200 text-gray-800":
      props.color === "secondary" || props.color === "delete",
    "hover:bg-red-500 hover:text-white": props.color === "delete"
  };
});
</script>

<template>
  <button
    :class="class_"
    class="px-4 py-2 text-gray-800"
  >
    <i class="text-base material-icons-round">{{ props.icon }}</i>
    {{ props.label }}
  </button>
</template>
