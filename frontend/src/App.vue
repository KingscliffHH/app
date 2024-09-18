<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";
import { watch } from "vue";
import { useRouter } from "vue-router";

import services from "@/services";
import { token, user } from "@/store/auth";

const auth0 = useAuth0();
const router = useRouter();

watch(auth0.isAuthenticated, async (newVal) => {
  console.info(newVal);
  if (!newVal) {
    router.push("/login");
    return;
  }

  try {
    token.value = await auth0.getAccessTokenSilently();
    user.value = await services.users.me();
  } catch (err) {
    console.error("main", err);
    router.push("/login");
  }
});
</script>

<template>
  <!-- <div class="flex flex-col h-screen"> -->
  <router-view />
  <!-- </div> -->
</template>
