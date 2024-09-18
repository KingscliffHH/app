<script setup lang="ts">
import ChartJS from "chart.js/auto";
import { ref, onMounted } from "vue";

interface Props {
  milestone: {
    completed: number;
    total: number;
  };
}

const props = defineProps<Props>();

const chart = ref<HTMLCanvasElement | null>(null);

onMounted(() => {
  if (!chart.value) {
    return;
  }

  new ChartJS(chart.value, {
    type: "doughnut",
    data: {
      datasets: [
        {
          borderRadius: 10,
          data: [
            props.milestone.completed,
            props.milestone.total - props.milestone.completed
          ],
          backgroundColor: ["#6F8FB2", "#EDF2F7"]
        }
      ]
    },
    options: {
      maintainAspectRatio: false,
      cutout: "80%"
    }
  });
});
</script>

<template>
  <div>
    <span class="card-title"> Number Of Milestones </span>
    <div class="commercial-information__card-content">
      <div style="position: relative">
        <canvas
          ref="chart"
          style="width: 100%; height: 100%"
        />
        <span>{{ props.milestone.completed }}/{{ props.milestone.total }}</span>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.commercial-information {
  &__card-content {
    flex: 1;
    display: flex;
    padding: 10px;

    div > span {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      font-size: 45px;
      color: #2d4c6e;
    }
  }
}
</style>
