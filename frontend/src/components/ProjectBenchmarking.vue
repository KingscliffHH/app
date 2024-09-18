<script lang="ts" setup>
import { computed } from "vue";

import { type Benchmark } from "@/types/benchmark";
import { type Benchmarking } from "@/types/project";

import ProjectBenchmarkingChart from "@/components/ProjectBenchmarkingChart.vue";

import { useQuery } from "@/hooks/fetch";
import services from "@/services";

const props = defineProps<{ data: Benchmarking }>();

// const benchmarks = computed(() => {
//   return [...(props.data.benchmarks || [])];
// });

const { data: availableBenchmarks } = useQuery({
  queryFn: () => services.benchmarks.getAll()
});

const benchmarks = computed<Benchmark[]>(() => {
  if (!props.data.benchmarks) return [];
  if (!availableBenchmarks.value) return [];

  const res = [];

  for (const i in props.data.benchmarks) {
    const b = props.data.benchmarks[i];
    const found = availableBenchmarks.value.find(
      (ab) => ab.id === b.benchmarkId
    );

    if (!found) continue;

    if (!b.displayProjectName) found.name = `Project ${Number(i) + 1}`;
    res.push(found);
  }

  res.push({
    id: "current",
    name: "Current Project",
    geographicLocation: "Grange Rd",
    totalProjectCostP90: props.data.totalProjectCostP90,
    totalConstructionCostPerLaneKm: props.data.totalConstructionCostPerLaneKm,
    squareMetreRateForPavementPerBridgePerM2:
      props.data.squareMetreRateForPavementPerBridgePerM2,
    cubicMetreRateForEarthworksPerM3:
      props.data.cubicMetreRateForEarthworksPerM3
  });
  return res;
});

const colors = ["#3a6ba6", "#2C4C6E", "#00A1D7", "#7090B3", "#BDD8F6"];
const benchmarkLegends = computed(() => {
  return benchmarks.value.map((x, i) => ({ color: colors[i], name: x.name }));
});
</script>

<template>
  <v-card subtitle="Preferred Option Benchmarks">
    <div class="flex justify-center gap-2">
      <div
        v-for="{ name, color } in benchmarkLegends"
        :key="name"
        class="flex items-center gap-1 cursor-not-allowed"
      >
        <div
          class="h-4 w-10"
          :style="{ backgroundColor: color }"
        ></div>
        <span>{{ name }}</span>
      </div>
    </div>
    <div class="grid grid-cols-4">
      <ProjectBenchmarkingChart
        v-if="benchmarks.length"
        :data="benchmarks"
        title="Outturn Cost"
        field="totalProjectCostP90"
      />
      <ProjectBenchmarkingChart
        v-if="benchmarks.length"
        :data="benchmarks"
        title="Rate for Pavement (Cost/M2)"
        field="squareMetreRateForPavementPerBridgePerM2"
      />
      <ProjectBenchmarkingChart
        v-if="benchmarks.length"
        :data="benchmarks"
        title="Earthworks (Cost/M3)"
        field="cubicMetreRateForEarthworksPerM3"
      />
      <ProjectBenchmarkingChart
        v-if="benchmarks.length"
        :data="benchmarks"
        title="Cost/Lane KM (Total Construction)"
        field="totalConstructionCostPerLaneKm"
      />
      <!-- <TotalProjectCostPerMilestoneRiskContingency :data="props.data" />
      <TotalProjectCostPerMilestoneOutturnCost :data="props.data" />
      <TotalProjectCostPerMilestoneRiskContingency :data="props.data" /> -->
    </div>
  </v-card>
</template>
