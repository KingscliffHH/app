<script lang="ts" setup>
import ChartJS from "chart.js/auto";
import moment from "moment";
import { computed, onMounted, ref } from "vue";

const props = defineProps<{
  data: Map<
    string,
    {
      baseValue: number;
      p50: number;
      p90: number;
    }
  >;
}>();

// const props = {
//   data: new Map([
//     ["2021-01-01", { p50: 100_000_000, p90: 120_000_000 }],
//     ["2021-02-01", { p50: 110_000_000, p90: 130_000_000 }],
//     ["2021-03-01", { p50: 120_000_000, p90: 140_000_000 }],
//     ["2021-04-01", { p50: 130_000_000, p90: 150_000_000 }],
//     ["2021-05-01", { p50: 140_000_000, p90: 160_000_000 }],
//     ["2021-06-01", { p50: 140_000_000, p90: 160_000_000 }],
//     ["2021-07-01", { p50: 140_000_000, p90: 160_000_000 }]
//   ])
// };

const chart = ref<HTMLCanvasElement | null>(null);

interface Gradient {
  colors: {
    start: string;
    end: string;
  };
}

const mountGradient = (ctx: any, area: any, { colors }: Gradient) => {
  const gradient = ctx.createLinearGradient(0, area.bottom, 0, area.top);
  gradient.addColorStop(0, colors.start);
  gradient.addColorStop(0.3, colors.end);

  return gradient;
};

const fn = (context: any, colors: Gradient["colors"]) => {
  const { ctx, chartArea } = context.chart;

  if (!chartArea) {
    return;
  }

  return mountGradient(ctx, chartArea, { colors });
};

const dates = computed<string[]>(() => {
  return Array.from(props.data.keys());
});

const costs = computed<{ baseValue: number[]; p90: number[]; p50: number[] }>(
  () => {
    const values = Array.from(props.data.values());

    return {
      baseValue: values.map((v) => v.baseValue),
      p90: values.map((v) => v.p90),
      p50: values.map((v) => v.p50)
    };
  }
);

const mountChart = () => {
  if (!chart.value) {
    return;
  }

  new ChartJS(chart.value, {
    type: "bar",
    data: {
      labels: dates.value,
      datasets: [
        {
          label: "Base Value",
          data: costs.value.baseValue,
          borderRadius: 5,
          barThickness: 35,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#D8EFF8",
              end: "#BDD8F6"
            });
          }
        },
        {
          label: "P50",
          data: costs.value.p50,
          borderRadius: 5,
          barThickness: 35,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#94c8d9",
              end: "#00A1D7"
            });
          }
        },
        {
          label: "P90",
          data: costs.value.p90,
          borderRadius: 5,
          barThickness: 35,
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
      plugins: {
        legend: {
          display: true
        }
      },
      maintainAspectRatio: false,
      scales: {
        x: {
          stacked: true,
          ticks: {
            callback: function (i) {
              const value: string = this.getLabelForValue(i as number);
              return moment(value).format("DD MMM YY");
            }
          }
        },
        y: {
          min: 0,
          beginAtZero: true,
          ticks: {
            stepSize: 10_000_000,
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
      }
    }
  });
};

onMounted(mountChart);
</script>

<template>
  <div class="total-project-cost-per-milestone-outturn-cost">
    <span class="card-subtitle">Total Project Outturn Cost</span>
    <div class="total-project-cost-per-milestone-outturn-cost__card-content">
      <canvas ref="chart" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.total-project-cost-per-milestone-outturn-cost {
  flex: 1;
  display: flex;
  flex-direction: column;

  &__card-content {
    flex: 1;
  }
}
</style>
