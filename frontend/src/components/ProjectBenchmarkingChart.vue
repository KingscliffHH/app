<script setup lang="ts">
import ChartJS from "chart.js/auto";
import { ref, onMounted } from "vue";

import { Benchmark } from "@/types/benchmark";

const props = defineProps<{
  data: Benchmark[];
  field: keyof Benchmark;
  title: string;
}>();

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

const chart = ref<HTMLCanvasElement | null>(null);
onMounted(() => {
  if (!chart.value) {
    return;
  }

  new ChartJS(chart.value, {
    type: "bar",
    data: {
      labels: ["Benchmark"],
      datasets: props.data.map((d, i) => {
        return {
          label: d.name,
          data: [d[props.field]],
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
  <div class="px-2 pb-2 h-[150px]">
    <span class="text-sm text-gray-400">{{ props.title }}</span>
    <canvas ref="chart"></canvas>
  </div>
</template>
