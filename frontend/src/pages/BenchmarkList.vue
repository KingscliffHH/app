<script setup lang="ts">
import { toast } from "vue3-toastify";

import type { Benchmark } from "@/types/benchmark";

import BaseButtonOutlined from "@/components/base/BaseButtonOutlined.vue";

import { useMutation, useQuery } from "@/hooks/fetch";
import services from "@/services";

const onError = (err: any): void => {
  if ("message" in err) {
    toast.error(err.message, { autoClose: 5000 });
  } else {
    toast.error("Error fetching benchmarks!", { autoClose: 2000 });
  }
};

const { refetch, data: benchmarks } = useQuery({
  queryFn: () => services.benchmarks.getAll(),
  onError
});

const { mutate: deleteBenchmark } = useMutation({
  mutationFn: (id: string) => services.benchmarks.delete(id),
  onSuccess: () => toast.success("Success!", { autoClose: 2000 }),
  onError
});

const confirmDelete = async (item: Benchmark) => {
  if (confirm("Are you sure you want to delete this item?")) {
    await deleteBenchmark(item.id);
    refetch();
  }
};

const format = (number: number) => {
  return new Intl.NumberFormat("en-AU", {
    style: "currency",
    currency: "AUD"
  })
    .format(number)
    .replace("$", "");
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">Benchmarks</h1>
      <router-link to="/benchmarks/new">
        <BaseButtonOutlined
          color="primary"
          label="+ New"
        />
      </router-link>
    </section>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left rtl:text-right text-gray-500">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50">
          <tr>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Project Name
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Geographic Location
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-right"
            >
              Total Project Cost (P90)
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-right"
            >
              Total Construction Cost $/Lane Km
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in benchmarks"
            :key="item.id"
            class="odd:bg-white even:bg-gray-50 border-b"
          >
            <td
              scope="row"
              class="px-6 py-2"
            >
              {{ item.name }}
            </td>
            <td class="px-6 py-2">{{ item.geographicLocation }}</td>
            <td class="px-6 py-2 text-right">
              ${{ format(item.totalProjectCostP90) }}
            </td>
            <td class="px-6 py-2 text-right">
              ${{ format(item.totalConstructionCostPerLaneKm) }}
            </td>
            <td class="px-6 py-2 flex items-center space-x-2">
              <router-link :to="`/benchmarks/${item.id}`">
                <BaseButtonOutlined
                  label="View Details"
                  size="sm"
                />
              </router-link>
              <router-link :to="`/benchmarks/${item.id}?edit`">
                <button
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                >
                  <i class="text-base material-icons-round">edit</i>
                </button>
              </router-link>

              <button
                class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-red-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                @click="() => confirmDelete(item)"
              >
                <i class="text-base material-icons-round">delete</i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
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
