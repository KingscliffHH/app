<script setup lang="ts">
import { ref } from "vue";

const props = defineProps<{
  url?: string;
  readonly: boolean;
}>();

const emits = defineEmits<{
  (e: "update:image", value: File): void;
}>();

const preview = ref<string>(props.url || "https://via.placeholder.com/400");

const imageInput = ref<HTMLInputElement | null>(null);

const showImageDialog = () => {
  if (imageInput.value) {
    imageInput.value.click();
  }
};

const loadPreview = () => {
  if (!imageInput.value) return;

  const file = imageInput.value.files?.[0];

  if (file) {
    emits("update:image", file);

    const reader = new FileReader();
    reader.onload = (e) => {
      const target = e.target as FileReader;

      preview.value = target.result as string;
    };
    reader.readAsDataURL(file);
  }
};
</script>

<template>
  <div
    class="field member relative inline-flex size-48 bg-gray-200 flex justify-center items-center overflow-hidden"
    :class="{
      'cursor-pointer': !readonly,
      'opacity-50 cursor-not-allowed': readonly
    }"
    @click="!readonly && showImageDialog()"
  >
    <img
      ref="imagePreview"
      class="max-w-full max-h-full"
      :src="preview"
      alt="Sample Image"
    />
    <input
      ref="imageInput"
      type="file"
      class="hidden"
      accept="image/*"
      @input="loadPreview()"
    />
  </div>
</template>
