<script setup lang="ts">
import ChartJS from "chart.js/auto";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { ref, onMounted, computed } from "vue";

import { type CommercialInformation } from "@/types/project";

const props = defineProps<{
  data: CommercialInformation;
}>();

const remainingFee = computed(() => {
  return (
    props.data.ciContractedValue +
    props.data.approvedVariationToDate -
    props.data.ciAccrualToDate
  );
});

const chart = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (!chart.value) {
    return;
  }

  const total =
    props.data.ciContractedValue + props.data.approvedVariationToDate;

  new ChartJS(chart.value, {
    plugins: [ChartDataLabels],
    type: "doughnut",
    data: {
      labels: ["Accrual to Date", "Remaining Fee"],
      datasets: [
        {
          borderRadius: 8,
          spacing: 3,
          data: [props.data.ciAccrualToDate, remainingFee.value],
          backgroundColor: ["#6f8fb2", "#2d4c6e"]
        }
      ]
    },
    options: {
      plugins: {
        datalabels: {
          color: "#fff",
          formatter: (value: number) => {
            return `${Math.trunc((value / total) * 100)} %`;
          }
        },
        legend: {
          display: true,
          position: "right",
          labels: {
            padding: 10,
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
  <v-card subtitle="C&I Contracted Value">
    <div class="px-2 pb-2 h-[170px]">
      <canvas ref="chart"></canvas>
    </div>
  </v-card>
</template>
