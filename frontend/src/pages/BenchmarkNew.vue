<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";

import BenchmarkForm from "@/components/BenchmarkForm.vue";

import { useMutation } from "@/hooks/fetch";
import services from "@/services";

const router = useRouter();

const data = ref({
  name: "",
  geographicLocation: "",
  totalProjectCostP90: 0,
  totalConstructionCostPerLaneKm: 0,
  cubicMetreRateForEarthworksPerM3: 0,
  squareMetreRateForPavementPerBridgePerM2: 0
});

const {
  // isLoading: saving,
  // isError,
  error,
  // isSuccess,
  mutate: save
} = useMutation({
  mutationFn: (payload: any) => services.benchmarks.create(payload),
  onSuccess: (data) => {
    console.info("onSuccess", data);
    // toast("Success!");
    router.push("/benchmarks");
    toast.success("Success!", {
      autoClose: 2000
    });
  },
  onError: (err) => {
    toast.error("Error!", {
      autoClose: 2000
    });
    console.info("onError", err);
  }
});

const cancel = () => {
  router.push("/benchmarks");
};

const submit = async () => {
  await save(data.value);
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">Benchmark</h1>
      <section class="flex gap-4">
        <button
          class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded"
          type="button"
          @click="cancel"
        >
          Cancel
        </button>
        <button
          class="px-4 py-1 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
          type="submit"
          @click="submit"
        >
          Submit
        </button>
      </section>
    </section>
    <BenchmarkForm
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
