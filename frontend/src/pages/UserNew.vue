<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";

import { type User, EmptyUser } from "@/types/user";

import UserForm from "@/components/UserForm.vue";

import { useMutation } from "@/hooks/fetch";
import services from "@/services";

const MAX_IMAGE_SIZE_IN_MB = 1;
const router = useRouter();

const data = ref<User>({ ...EmptyUser });
const image = ref<File | null>(null);

const {
  isLoading: saving,
  // isError,
  error,
  // isSuccess,
  mutate: save
} = useMutation({
  mutationFn: async (payload: User) => {
    if (!image.value) {
      throw new Error("Please upload an image.");
    }

    const { imageUrl } = await services.users.uploadAvatar(image.value);

    payload.avatar = imageUrl;
    return services.users.create(payload);
  },
  onSuccess: (data) => {
    console.info("onSuccess", data);
    // toast("Success!");
    router.push("/users");
    toast.success("Success!", {
      autoClose: 2000
    });
  },
  onError: (err) => {
    console.info("onError", err);
    if ("message" in err) {
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

const cancel = () => {
  router.push("/users");
};

const submit = async () => {
  if (!image.value) {
    toast.error("Please upload an image.", { autoClose: 5000 });
    return;
  }

  if (!validateImageSize(image.value)) {
    return;
  }

  await save(data.value);
};

function bytes2Size(byteVal: number) {
  var units = ["Bytes", "KB", "MB", "GB", "TB"];
  var kounter = 0;
  var kb = 1000;
  var div = byteVal / 1;
  while (div >= kb) {
    kounter++;
    div = div / kb;
  }
  return div.toFixed(1) + " " + units[kounter];
}

const validateImageSize = (file: File): boolean => {
  const maxSize = MAX_IMAGE_SIZE_IN_MB * 1000 * 1000;
  if (file.size > maxSize) {
    const formattedSize = bytes2Size(file.size);
    toast.error(`Image size (${formattedSize}) must be less than 1MB`, {
      autoClose: 5000
    });
    return false;
  }
  return true;
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">System User</h1>
      <section class="flex gap-4">
        <button
          class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded disabled:opacity-50 disabled:cursor-not-allowed"
          type="button"
          :disabled="saving"
          @click="cancel"
        >
          Cancel
        </button>
        <button
          class="px-4 py-1 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed"
          type="submit"
          :disabled="saving"
          @click="submit"
        >
          Submit
        </button>
      </section>
    </section>
    <UserForm
      v-model="data"
      v-model:image="image"
      :error="error"
      :readonly="false"
      @update:image="validateImageSize"
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
