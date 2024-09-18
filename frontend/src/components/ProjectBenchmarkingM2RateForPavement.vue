<script setup lang="ts">
import ChartJS from "chart.js/auto";
import { ref, onMounted } from "vue";

import { Benchmark } from "@/types/benchmark";

const props = defineProps<{
  data: Benchmark[];
}>();

const chart = ref<HTMLCanvasElement | null>(null);

interface Gradient {
  colors: {
    start: string;
    end: string;
  };
}

const colors: Gradient["colors"][] = [
  { start: "#7db5ff", end: "#3a6ba6" },
  { start: "#6990ba", end: "#2C4C6E" },
  { start: "#94c8d9", end: "#00A1D7" },
  { start: "#CBE2F8", end: "#7090B3" },
  { start: "#D8EFF8", end: "#BDD8F6" }
];

const mountGradient = (ctx: any, area: any, { colors }: Gradient) => {
  const gradient = ctx.createLinearGradient(area.left, 0, area.right, 0);
  gradient.addColorStop(0, colors.start);
  gradient.addColorStop(0.2, colors.end);

  return gradient;
};

const fn = (context: any, colors: Gradient["colors"]) => {
  const { ctx, chartArea } = context.chart;

  if (!chartArea) {
    return;
  }

  return mountGradient(ctx, chartArea, { colors });
};

onMounted(() => {
  if (!chart.value) {
    return;
  }

  new ChartJS(chart.value, {
    type: "bar",
    data: {
      labels: ["Benchmark"],
      datasets: [
        ...props.data.map((d, i) => {
          return {
            label: d.name,
            data: [d.squareMetreRateForPavementPerBridgePerM2],
            borderRadius: 5,
            barPercentage: 0.7,
            maxBarThickness: 25,
            backgroundColor: (context: any) => {
              const c = (colors[i] || {
                start: "#D8EFF8",
                end: "#BDD8F6"
              }) as Gradient["colors"];
              return fn(context, { ...c });
            }
          };
        })
      ]
    },
    options: {
      indexAxis: "y",
      maintainAspectRatio: false,
      scales: {
        y: {
          display: false
        },
        x: {
          ticks: {
            callback: function (value) {
              return Number(value).toLocaleString("en-AU", {
                currency: "AUD",
                style: "currency",
                notation: "compact",
                compactDisplay: "short"
              });
            }
          }
        }
      },
      plugins: {
        legend: {
          display: false
        }
      }
    }
  });
});
</script>

<template>
  <div class="project-benchmarking-m2-rate-for-pavement">
    <span class="card-subtitle">Rate for Pavement (Cost/M2)</span>
    <div class="project-benchmarking-m2-rate-for-pavement__card-content">
      <canvas ref="chart"></canvas>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.project-benchmarking-m2-rate-for-pavement {
  flex: 1;
  display: flex;
  flex-direction: column;

  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
  }
}
</style>
