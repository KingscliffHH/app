<script setup lang="ts">
import { computed, ref } from "vue";

import type { Package } from "@/types/project";

const props = defineProps<{
  data: Package[][];
  milestones: number;
}>();

const milestones = computed(() => {
  const options = [...Array(props.milestones).keys()].map((_, i) => ({
    title: `Milestone ${i + 1}`,
    value: i
  }));
  return [{ title: "---", value: -1 }, ...options];
});

const selectedMilestone = ref<number>(-1);

const filteredPackages = computed(() => {
  if (selectedMilestone.value === -1) {
    return [];
  }

  return (
    props.data
      .map((pkg) => pkg[selectedMilestone.value])
      .filter((pkg) => pkg.description !== "") || []
  );
});
</script>

<template>
  <div>
    <span class="card-title">Packages</span>
    <div class="w-25 my-4">
      <v-select
        v-model="selectedMilestone"
        hide-details
        color="#1A3C5B"
        label="Select Milestone"
        :items="milestones"
      />
    </div>
    <div class="project-packages__card-content">
      <table>
        <thead>
          <tr>
            <th />
            <th>Progress %</th>
            <th>Second Stage (Senior QS) QA<br />Review Completed (50%)</th>
            <th>Final QA Review<br />Complated (90%)</th>
            <th>Submited (100%)</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(pkg, index) in filteredPackages"
            :key="index"
          >
            <td>{{ pkg.description }}</td>
            <td>{{ pkg.progress }}%</td>
            <td>{{ pkg.secondStageQA ? "Yes" : "No" }}</td>
            <td>{{ pkg.finalQAReview ? "Yes" : "No" }}</td>
            <td>{{ pkg.submitted ? "Yes" : "No" }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.project-packages {
  &__card-content {
    flex: 1;
    display: flex;
    padding: 10px;
  }
}

table {
  flex: 1;

  & th {
    padding: 10px 20px;
    font-weight: 500;
    text-align: center;

    &:first-child {
      border-radius: 10px 0 0 0;
    }

    &:last-child {
      border-radius: 0 10px 0 0;
    }
  }

  & thead > tr {
    background-color: #f2f2f2;
  }

  & td {
    padding: 10px 20px;
    font-weight: 400;

    &:first-child {
      text-align: left;
    }
  }

  & tbody {
    tr {
      border: 1px solid #e0e0e0;

      &:nth-child(even) {
        background-color: #f9fafb;
      }

      &:nth-child(odd) {
        background-color: #ffffff;
      }

      td {
        color: #4d4d4d;

        &:not(:nth-child(1)) {
          text-align: center;
        }
      }
    }
  }
}

@media (max-height: 1080px) or (max-width: 1920px) {
  td {
    padding: 5px 10px;
  }

  table thead tr th {
    font-size: 14px;
  }

  table tbody tr td {
    font-size: 12px;
  }
}
</style>
