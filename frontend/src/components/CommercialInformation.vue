<script setup lang="ts">
import ChartJS from "chart.js/auto";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { ref, onMounted } from "vue";

const props = defineProps<{ remainingFee: number; accrualToDate: number }>();

const chart = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (!chart.value) {
    return;
  }

  const total = props.remainingFee + props.accrualToDate;

  new ChartJS(chart.value, {
    plugins: [ChartDataLabels],
    type: "doughnut",
    data: {
      labels: ["Accrual to Date", "Remaining Fee"],
      datasets: [
        {
          borderRadius: 8,
          spacing: 3,
          data: [props.accrualToDate, props.remainingFee],
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
            padding: 4,
            font: {
              size: 16
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
    <span class="card-title"> Commercial Information </span>
    <div class="commercial-information__card-content">
      <div>
        <canvas
          ref="chart"
          style="width: 100%; height: 100%"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.commercial-information {
  &__card-content {
    flex: 1;
    display: flex;
  }
}
</style>
