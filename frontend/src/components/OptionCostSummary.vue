<script setup lang="ts">
import ChartJS from "chart.js/auto";
import { ref, onMounted } from "vue";

import { OptionOutturnCost } from "@/types/project";

const props = defineProps<{
  data: OptionOutturnCost[];
}>();

const chart = ref<HTMLCanvasElement | null>(null);

interface Gradient {
  colors: {
    start: string;
    end: string;
  };
}

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

  new ChartJS(chart?.value, {
    type: "bar",
    data: {
      labels: props.data.map((_: unknown, i: number) => `Option ${i + 1}`),
      datasets: [
        {
          label: "Base",
          data: props.data.map((d: any) => d.base),
          borderRadius: 5,
          barThickness: 22,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#D8EFF8",
              end: "#BDD8F6"
            });
          }
        },
        {
          label: "P50",
          data: props.data.map((d: any) => d.p50),
          borderRadius: 5,
          barThickness: 22,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#94c8d9",
              end: "#00A1D7"
            });
          }
        },
        {
          label: "P90",
          data: props.data.map((d: any) => d.p90),
          borderRadius: 5,
          barThickness: 22,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#6990ba",
              end: "#2C4C6E"
            });
          }
        }
      ]
    },
    options: {
      indexAxis: "y",
      scales: {
        y: {
          stacked: true
        },
        x: {
          ticks: {
            callback: function (value) {
              return Number(value).toLocaleString("en-AU", {
                style: "currency",
                notation: "compact",
                compactDisplay: "short",
                // maximumFractionDigits: 2,
                currency: "AUD"
              });
            }
          }
        }
      },
      plugins: {
        tooltip: {
          mode: "y",
          intersect: false
        }
      },
      maintainAspectRatio: false
    }
  });
});
</script>

<template>
  <div>
    <span class="card-title">Option Cost Summary</span>
    <div class="key-cost-drivers__card-content">
      <canvas ref="chart" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.key-cost-drivers {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 10px;
  }
}
</style>
