<script lang="ts" setup>
import { computed } from "vue";

import { Benchmark } from "@/types/benchmark";
import { Project } from "@/types/project";

import KeyCostDrivers from "./KeyCostDrivers.vue";
import KeyRisks from "./KeyRisks.vue";
import OptionCostSummary from "./OptionCostSummary.vue";
import ProjectBenchmarkingM2RateForPavement from "./ProjectBenchmarkingM2RateForPavement.vue";
import ProjectBenchmarkingM3RateForEarthworks from "./ProjectBenchmarkingM3RateForEarthworks.vue";
import ProjectBenchmarkingTotalConstructionCost from "./ProjectBenchmarkingTotalConstructionCost.vue";
import ProjectBenchmarkingTotalProjectCost from "./ProjectBenchmarkingTotalProjectCost.vue";
import TotalProjectCostPerMilestoneOutturnCost from "./TotalProjectCostPerMilestoneOutturnCost.vue";
import TotalProjectCostPerMilestoneRiskContingency from "./TotalProjectCostPerMilestoneRiskContingency.vue";
import ValueManagement from "./ValueManagement.vue";

const props = defineProps<{
  project: Project;
  availableBenchmarks: Benchmark[];
}>();

const totalProjectCostPerMilestoneOutturnCost = computed(() => {
  return new Map(
    (props.project.metrics.totalProjectCostPerMilestone || []).map((x) => {
      return [
        x.date.toISOString().slice(0, 10),
        { baseValue: x.baseValue, p50: x.p50OutturnCost, p90: x.p90OutturnCost }
      ];
    })
  );
});

const totalProjectCostPerMilestoneRiskContingency = computed(() => {
  return new Map(
    (props.project.metrics.totalProjectCostPerMilestone || []).map((x) => {
      return [
        x.date.toISOString().slice(0, 10),
        { p50: x.p50RiskContingency, p90: x.p90RiskContingency }
      ];
    })
  );
});

const valueManagement = computed(() => {
  return (props.project.metrics.valueManagementOpportunities || []).map(
    (x) => ({ opportunity: x })
  );
});

const benchmarks = computed<Benchmark[]>(() => {
  if (!props.project.metrics.benchmarking.benchmarks) return [];
  if (!props.availableBenchmarks) return [];

  const res = [];

  for (const i in props.project.metrics.benchmarking.benchmarks) {
    const b = props.project.metrics.benchmarking.benchmarks[i];
    const found = props.availableBenchmarks.find(
      (ab) => ab.id === b.benchmarkId
    );

    if (!found) continue;

    if (!b.displayProjectName) found.name = `Project ${Number(i) + 1}`;
    res.push(found);
  }

  res.push({
    id: "current",
    name: props.project.name,
    geographicLocation: "Grange Rd",
    totalProjectCostP90: props.project.metrics.benchmarking.totalProjectCostP90,
    totalConstructionCostPerLaneKm:
      props.project.metrics.benchmarking.totalConstructionCostPerLaneKm,
    squareMetreRateForPavementPerBridgePerM2:
      props.project.metrics.benchmarking
        .squareMetreRateForPavementPerBridgePerM2,
    cubicMetreRateForEarthworksPerM3:
      props.project.metrics.benchmarking.cubicMetreRateForEarthworksPerM3
  });
  return res;
});

const colors = ["#3a6ba6", "#2C4C6E", "#00A1D7", "#7090B3", "#BDD8F6"];
const benchmarkLegends = computed<{ name: string; color: string }[]>(() => {
  return benchmarks.value.map((x, i) => ({ color: colors[i], name: x.name }));
});
//  = [
//   {
//     name: "Benchmark 1",
//     color: "#3a6ba6"
//   },
//   {
//     name: "Benchmark 2",
//     color: "#2C4C6E"
//   },
//   {
//     name: "Benchmark 3",
//     color: "#00A1D7"
//   },
//   {
//     name: "Benchmark 4",
//     color: "#7090B3"
//   },
//   {
//     name: "Benchmark 5",
//     color: "#BDD8F6"
//   }
// ];
</script>

<template>
  <div class="grid-container">
    <OptionCostSummary
      class="component component1"
      :data="props.project.metrics.optionOutturnCosts || []"
    />
    <KeyCostDrivers
      class="component component2"
      :data="{
        baseValue: props.project.metrics.keyCostDriversBaseValue,
        drivers: props.project.metrics.keyCostDrivers || []
      }"
    />
    <KeyRisks
      class="component component3"
      :data="props.project.metrics.keyRisks || []"
    />

    <div class="component component4">
      <span class="card-title">Preferred Option Cost (per Milestone)</span>
      <div class="inline-chart">
        <TotalProjectCostPerMilestoneOutturnCost
          :data="totalProjectCostPerMilestoneOutturnCost"
        />
        <TotalProjectCostPerMilestoneRiskContingency
          :data="totalProjectCostPerMilestoneRiskContingency"
        />
      </div>
    </div>

    <ValueManagement
      class="component component5"
      :data="valueManagement"
    />

    <div class="component component6">
      <span class="card-title">Preferred Option Benchmarks</span>

      <div class="legend">
        <div
          v-for="{ name, color } in benchmarkLegends"
          :key="name"
          class="legend__content"
        >
          <div
            class="legend__color"
            :style="{ backgroundColor: color }"
          />
          <span>{{ name }}</span>
        </div>
      </div>

      <div class="inline-chart">
        <ProjectBenchmarkingTotalProjectCost
          v-if="props.project.metrics.benchmarking.enableTotalProjectCost"
          :data="benchmarks"
        />
        <ProjectBenchmarkingM2RateForPavement
          v-if="
            props.project.metrics.benchmarking.enableSquareMetreRateForPavement
          "
          :data="benchmarks"
        />
        <ProjectBenchmarkingM3RateForEarthworks
          v-if="
            props.project.metrics.benchmarking.enableCubicMetreRateForEarthworks
          "
          :data="benchmarks"
        />
        <ProjectBenchmarkingTotalConstructionCost
          v-if="props.project.metrics.benchmarking.enableTotalConstructionCost"
          :data="benchmarks"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.grid-container {
  flex: 1;
  gap: 15px;
  display: grid;
  margin-top: 15px;
  grid-template-columns: repeat(4, 1fr);
  // breaking for small resolutions
  grid-template-rows: auto auto 25vh;
}

.component {
  display: flex;
  padding: 15px;
  border-radius: 18px;
  flex-direction: column;
  background-color: #ffffff;
  border: 2px solid #eff0f6;
}

.component1 {
  grid-area: 1 / 1 / 2 / 3;
}

.component2 {
  grid-area: 1 / 3 / 2 / 4;
}

.component3 {
  grid-area: 1 / 4 / 2 / 5;
}

.component4 {
  grid-area: 2 / 1 / 3 / 4;
}

.component5 {
  grid-area: 2 / 4 / 3 / 5;
}

.component6 {
  grid-area: 3 / 1 / 4 / 5;
}

.inline-chart {
  display: flex;
  flex: 1;
  padding: 10px;
  gap: 30px;
}

.legend {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 15px;

  &__content {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 5px;
    cursor: not-allowed;

    > div {
      width: 40px;
      height: 12px;
    }

    > span {
      color: #666666;
      font-size: 14px;
    }
  }
}
</style>
