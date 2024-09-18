<script setup lang="ts">
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { toast } from "vue3-toastify";

import { User } from "@/types/user";

import UserForm from "@/components/UserForm.vue";

import { useQuery, useMutation } from "@/hooks/fetch";
import services from "@/services";

const props = defineProps<{
  id: string;
}>();

const router = useRouter();
const route = useRoute();

const {
  isLoading: fetching,
  reset,
  data
} = useQuery({
  queryFn: () => services.users.get(props.id)
});

const image = ref<File | null>(null);
const {
  isLoading: saving,
  // isError,
  error,
  // isSuccess,
  mutate: save
} = useMutation({
  mutationFn: async (id: string, payload: User) => {
    console.info("payload", payload);
    if (image.value) {
      try {
        const { imageUrl } = await services.users.uploadAvatar(image.value);
        payload.avatar = imageUrl;
      } catch (err) {
        throw new Error("Error uploading an image.");
      }
    }

    console.info("payload", payload);
    return services.users.update(id, payload);
  },
  onSuccess: (data) => {
    console.info("onSuccess", data);
    toast.success("Success!", {
      autoClose: 2000
    });
    router.push("/users");
  },
  onError: (err) => {
    console.info("onError", err.message);
    if (err instanceof Error) {
      toast.error(err.message, {
        autoClose: 5000
      });
      return;
    }

    toast.error("Error!", {
      autoClose: 2000
    });
  }
});
const state = ref<"view" | "edit">("edit" in route.query ? "edit" : "view");

const submit = async () => {
  await save(props.id, data.value);
};

const cancel = () => {
  reset();
  state.value = "view";
};

const goBack = () => {
  router.push("/users");
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">System User</h1>
      <section class="flex gap-4">
        <v-btn
          color="#2c4c6e"
          variant="tonal"
          :class="{ hidden: state === 'edit' }"
          @click="goBack"
        >
          <i class="material-icons-round">arrow_back</i>
          <v-tooltip
            activator="parent"
            location="start"
          >
            Back
          </v-tooltip>
        </v-btn>
        <button
          class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded disabled:opacity-50 disabled:cursor-not-allowed"
          :class="{ hidden: state === 'view' }"
          type="button"
          :disabled="saving"
          @click="cancel"
        >
          Cancel
        </button>
        <button
          class="px-4 py-1 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
          :class="{ hidden: state === 'edit' }"
          type="button"
          @click="state = 'edit'"
        >
          Edit
        </button>
        <button
          class="px-4 py-1 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed"
          :class="{ hidden: state === 'view' }"
          type="submit"
          :disabled="saving"
          @click="submit"
        >
          Submit
        </button>
      </section>
    </section>
    <UserForm
      v-if="!fetching"
      v-model="data"
      v-model:image="image"
      :error="error"
      :readonly="state === 'view'"
    />
  </main>
</template>

<style lang="scss">
.main {
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin-left: 80px;
  background-color: #f9f9f9;
  padding: 15px;
}
</style>
