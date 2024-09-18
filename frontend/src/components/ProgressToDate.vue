<script setup lang="ts">
import ChartJS from "chart.js/auto";
import { ref, onMounted } from "vue";

import { ProgressToDate, ScopeOfEngagement } from "@/types/project";

const props = defineProps<{
  progressToDate: ProgressToDate;
  scope: ScopeOfEngagement;
}>();

interface Gradient {
  percentage: number;
  colors: {
    start: string;
    end: string;
  };
}

const chart = ref<HTMLCanvasElement | null>(null);

const mountGradient = (
  ctx: any,
  area: any,
  { colors, percentage }: Gradient
) => {
  const gradient = ctx.createLinearGradient(area.left, 0, area.right, 0);
  const endPosition = percentage / 100;
  gradient.addColorStop(0, colors.start);
  gradient.addColorStop(endPosition, colors.end);

  return gradient;
};

const fn = (context: any, colors: Gradient["colors"]) => {
  const { ctx, chartArea } = context.chart;
  const percentage = context.dataset.data[0];

  if (!chartArea) {
    return;
  }

  return mountGradient(ctx, chartArea, { percentage, colors });
};

const data = {
  labels: ["Percentage"],
  datasets: [] as any[]
};

if (props.scope.quantification) {
  data.datasets.push({
    label: "Quantification",
    data: [props.progressToDate.quantification],
    borderRadius: 10,
    barPercentage: 0.4,
    maxBarThickness: 25,
    backgroundColor: (context: any) => {
      return fn(context, {
        start: "#D8EFF8",
        end: "#BDD8F6"
      });
    }
  });
}

if (props.scope.costEstimation) {
  data.datasets.push({
    label: "Cost Estimation",
    data: [props.progressToDate.costEstimation],
    borderRadius: 10,
    barPercentage: 0.4,
    maxBarThickness: 25,
    backgroundColor: (context: any) => {
      return fn(context, {
        start: "#94c8d9",
        end: "#00A1D7"
      });
    }
  });
}

if (props.scope.probabilisticRiskAssessment) {
  data.datasets.push({
    label: "Probabilistic Risk Assessment",
    data: [props.progressToDate.probabilisticRiskAssessment],
    borderRadius: 10,
    barPercentage: 0.4,
    maxBarThickness: 25,
    backgroundColor: (context: any) => {
      return fn(context, {
        start: "#CBE2F8",
        end: "#7090B3"
      });
    }
  });
}

if (props.scope.basisOfEstimateReport) {
  data.datasets.push({
    label: "Basis Of Estimate Report",
    data: [props.progressToDate.basisOfEstimateReport],
    borderRadius: 10,
    barPercentage: 0.4,
    maxBarThickness: 25,
    backgroundColor: (context: any) => {
      return fn(context, {
        start: "#6990ba",
        end: "#2C4C6E"
      });
    }
  });
}

onMounted(() => {
  if (!chart.value) return;

  new ChartJS(chart?.value, {
    type: "bar",
    data,
    options: {
      indexAxis: "y",
      scales: {
        x: {
          min: 0,
          max: 100,
          ticks: {
            callback: function (value) {
              return `${value}%`;
            }
          }
        },
        y: {
          display: false
        }
      },
      plugins: {
        legend: {
          labels: {
            padding: 20,
            font: {
              size: 18
            }
          }
        }
      },
      maintainAspectRatio: false
    }
  });
});
</script>

<template>
  <div>
    <span class="card-title">Progress To Date</span>
    <div class="progress-to-date__card-content">
      <canvas ref="chart" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.progress-to-date {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 10px;
  }
}
</style>
