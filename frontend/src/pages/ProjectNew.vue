<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";

import { type Project, EmptyProject } from "@/types/project";

import ProjectForm from "@/components/ProjectForm.vue";

import { useMutation } from "@/hooks/fetch";
import services from "@/services";

const router = useRouter();
const data = ref<Project>(JSON.parse(JSON.stringify(EmptyProject)));

const {
  // isLoading: saving,
  error,
  // isSuccess,
  mutate: save
} = useMutation({
  mutationFn: (project: any) => services.projects.create(project),
  onSuccess: (data) => {
    console.info("onSuccess", data);
    toast.success("Success!", {
      autoClose: 2000
    });
    router.push("/projects");
  },
  onError: (err) => {
    toast.error("Error!", {
      autoClose: 2000
    });
    console.info("onError", err);
  }
});

const submit = async () => {
  await save(data.value);
};

const cancel = () => {
  router.push("/projects");
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">New Project</h1>
      <section>
        <button
          class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded"
          type="button"
          hx-get="/projects"
          hx-replace-url="true"
          hx-swap="innerHTML"
          hx-target="#main"
          hx-trigger="click"
          @click="cancel"
        >
          Cancel
        </button>
        <button
          class="px-4 py-1 ml-2 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
          type="submit"
          @click="submit"
        >
          Submit
        </button>
      </section>
    </section>
    <ProjectForm
      v-model="data"
      :error="error"
      :readonly="false"
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
