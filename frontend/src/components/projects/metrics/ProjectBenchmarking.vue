<script setup lang="ts">
import { computed, onMounted } from "vue";

import { ascending } from "@/utils/sort";

import type { Benchmarking } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseNumberInput from "@/components/base/BaseNumberInput.vue";
import BaseSelect from "@/components/base/BaseSelect.vue";
import BaseToggle from "@/components/base/BaseToggle.vue";

import { useQuery } from "@/hooks/fetch";
import services from "@/services";

const props = defineProps<{
  modelValue: Benchmarking;
  projectName: string;
  disabled?: boolean;
  error?: Record<string, string>;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: Benchmarking): void;
}>();

const EmptyObject = () => {
  return {
    benchmarkId: "",
    displayProjectName: false
  };
};

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});

const benchmarks = computed({
  get: () => data.value.benchmarks || [],
  set: (value) => (data.value.benchmarks = value)
});

const limit = 3;
const { data: availableBenchmarks } = useQuery({
  queryFn: () => services.benchmarks.getAll()
});

const canAddMore = computed(() => benchmarks.value.length < limit);

const options = computed(() => {
  return (availableBenchmarks.value || [])
    .map((b) => ({
      label: b.name,
      value: b.id
    }))
    .sort((a, b) => ascending(a.label, b.label));
});

const addBenchmark = () => {
  benchmarks.value = [...benchmarks.value, EmptyObject()];
  console.info("addBenchmark", benchmarks.value);
};

onMounted(() => {
  if (
    !props.modelValue.benchmarks ||
    props.modelValue.benchmarks.length === 0
  ) {
    data.value.benchmarks = Array.from({ length: 3 }, () => EmptyObject());
  }
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <h2 class="pr-4 text-lg font-medium text-gray-700">Benchmarking:</h2>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 items-center mb-4">
      <div class="text-sm font-medium text-gray-700">
        {{ props.projectName }}
      </div>
      <!-- <BaseToggle
        v-model="data.enableGeographicLocation"
        label="Geographic Location"
        name="geographicLocation"
      /> -->
      <BaseToggle
        v-model="data.enableTotalProjectCost"
        label="Total Project Cost (P90)"
        name="totalProjectCost"
        :disabled="disabled"
      />

      <BaseToggle
        v-model="data.enableTotalConstructionCost"
        label="Total Construction Cost $/Lane Km"
        name="totalConstructionCost"
        :disabled="disabled"
      />

      <BaseToggle
        v-model="data.enableSquareMetreRateForPavement"
        label="M2 Rate for Pavement/bridge $/m2"
        name="squareMetreRate"
        :disabled="disabled"
      />
      <BaseToggle
        v-model="data.enableCubicMetreRateForEarthworks"
        label="M3 Rate for Earthworks $/m3"
        name="cubicMetreRate"
        :disabled="disabled"
      />
    </div>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 items-center mb-4">
      <BaseInput
        v-model="data.geographicLocation"
        label="Geographic Location"
        name="geographicLocation"
        :disabled="disabled"
      />

      <BaseNumberInput
        v-model="data.totalProjectCostP90"
        label="Total Project Cost (P90)"
        name="totalProjectCost"
        prefix="$"
        :precision="2"
        :readonly="!data.enableTotalProjectCost"
        :disabled="disabled"
      />

      <BaseNumberInput
        v-model="data.totalConstructionCostPerLaneKm"
        label="Total Construction Cost $/Lane Km"
        name="totalConstructionCost"
        prefix="$"
        :precision="2"
        :readonly="!data.enableTotalConstructionCost"
        :disabled="disabled"
      />

      <BaseNumberInput
        v-model="data.squareMetreRateForPavementPerBridgePerM2"
        label="M2 Rate for Pavement/bridge $/m2"
        name="squareMetreRate"
        prefix="$"
        :precision="2"
        :readonly="!data.enableSquareMetreRateForPavement"
        :disabled="disabled"
      />

      <BaseNumberInput
        v-model="data.cubicMetreRateForEarthworksPerM3"
        label="M3 Rate for Earthworks $/m3"
        name="cubicMetreRate"
        prefix="$"
        :precision="2"
        :readonly="!data.enableCubicMetreRateForEarthworks"
        :disabled="disabled"
      />
    </div>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 items-center mb-4">
      <v-card
        v-for="(benchmark, index) in benchmarks"
        :key="index"
        class="flex flex-col p-4"
      >
        <BaseSelect
          v-model="benchmark.benchmarkId"
          label="Benchmark"
          :name="`benchmark-${index}`"
          :errors="props.error"
          :options="options"
          :disabled="disabled"
        />
        <BaseToggle
          v-model="benchmark.displayProjectName"
          label="Display Project Name"
          name="displayProjectName"
          :disabled="disabled"
        />
        <button
          class="disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-gray-100"
          :disabled="disabled"
          @click="() => benchmarks.splice(index, 1)"
        >
          Remove
        </button>
      </v-card>
    </div>
    <button
      class="mt-4 px-4 py-2 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-blue-500"
      :disabled="!canAddMore || disabled"
      @click="addBenchmark"
    >
      Add More
    </button>
  </section>
</template>
