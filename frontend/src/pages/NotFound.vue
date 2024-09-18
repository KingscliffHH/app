<script setup lang="ts">
import { useRouter } from "vue-router";

const $router = useRouter();

const back = () => {
  $router.back();
};
</script>

<template>
  <div class="flex h-full justify-center">
    <div>
      <h1 class="text-[10rem] font-bold text-blue-500">404</h1>
      <p class="text-2xl font-semibold">Page not found</p>
      <div class="actions flex gap-4">
        <button
          class="text-blue-500 hover:underline"
          @click="back"
        >
          Back
        </button>
        <router-link
          to="/"
          class="text-blue-500 hover:underline"
        >
          Home
        </router-link>
      </div>
    </div>
  </div>
</template>
