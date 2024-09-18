<script lang="ts" setup>
import ChartJS from "chart.js/auto";
import moment from "moment";
import { computed, onMounted, ref } from "vue";

// const props = defineProps<{ data: any }>();

const props = defineProps<{
  data: Map<
    string,
    {
      p50: number;
      p90: number;
    }
  >;
}>();

// const props = {
//   data: new Map([
//     ["2021-01-01", { p90: 15, p50: 12 }],
//     ["2021-02-01", { p90: 20, p50: 13 }],
//     ["2021-03-01", { p90: 30, p50: 15 }],
//     ["2021-04-01", { p90: 40, p50: 20 }],
//     ["2021-05-01", { p90: 50, p50: 35 }],
//     ["2021-06-01", { p90: 55, p50: 37 }],
//     ["2021-07-01", { p90: 59, p50: 49 }]
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

const dates = computed<string[]>(() => {
  return Array.from(props.data.keys());
});

const costs = computed<{ p90: number[]; p50: number[] }>(() => {
  const values = Array.from(props.data.values());

  return {
    p90: values.map((v: any) => v.p90),
    p50: values.map((v: any) => v.p50)
  };
});

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
          label: "P90",
          data: costs.value.p90,
          stack: "2",
          order: 2,
          borderRadius: 5,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#6990ba",
              end: "#2C4C6E"
            });
          }
        },
        {
          label: "P50",
          data: costs.value.p50,
          stack: "1",
          order: 1,
          borderRadius: 5,
          backgroundColor: (context: any) => {
            return fn(context, {
              start: "#94c8d9",
              end: "#00A1D7"
            });
          }
        }
      ]
    },
    options: {
      datasets: {
        bar: {
          barPercentage: 3,
          maxBarThickness: 30,
          categoryPercentage: 0.3
        }
      },
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
          suggestedMax: 60,
          beginAtZero: true,
          ticks: {
            stepSize: 20,
            callback: function (value) {
              return value + "%";
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
  <div class="total-project-cost-per-milestone-risk-contingency">
    <span class="card-subtitle">Risk (Contingency)</span>
    <div
      class="total-project-cost-per-milestone-risk-contingency__card-content"
    >
      <canvas ref="chart" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.total-project-cost-per-milestone-risk-contingency {
  flex: 1;
  display: flex;
  flex-direction: column;

  &__card-content {
    flex: 1;
  }
}
</style>
