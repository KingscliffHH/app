<script lang="ts" setup>
import ChartJS from "chart.js/auto";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { ref, onMounted } from "vue";

import { KeyCostDriver } from "@/types/project";

const props = defineProps<{
  data: {
    baseValue: number;
    drivers: KeyCostDriver[];
  };
}>();

const chart = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (!chart.value) {
    return;
  }

  console.info("chart", chart.value, props.data);

  const drivers = props.data?.drivers?.map((d: any) => d.driver) || [];
  const costs = props.data?.drivers?.map((d: any) => d.cost) || [];
  const totalCost = costs.reduce((acc: number, cost: number) => acc + cost, 0);

  new ChartJS(chart.value, {
    plugins: [ChartDataLabels],
    type: "doughnut",
    data: {
      labels: [...drivers, "Other"],
      datasets: [
        {
          label: "Cost",
          borderRadius: 8,
          spacing: 3,
          data: [...costs, props.data.baseValue - totalCost],
          backgroundColor: [
            "#CCCCCC",
            "#2C4C6E",
            "#00A1D7",
            "#656262",
            "#9FB7D3",
            "#6F8FB2"
          ]
        }
      ]
    },
    options: {
      plugins: {
        datalabels: {
          color: "#fff",
          formatter: (value: number) => {
            return `${Math.trunc((value / props.data.baseValue) * 100)} %`;
          }
        },
        legend: {
          display: true,
          position: "right"
        }
      },
      maintainAspectRatio: false
    }
  });
});
</script>

<template>
  <div>
    <span class="card-title">Key Cost Drivers</span>
    <div class="option-cost-summary__card-content">
      <canvas ref="chart" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.option-cost-summary {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 10px;
  }
}
</style>
