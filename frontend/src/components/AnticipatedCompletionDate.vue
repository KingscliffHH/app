<script lang="ts" setup>
import moment from "moment";

import { AnticipatedCompletionDate, ScopeOfEngagement } from "@/types/project";

interface TimelineItem {
  text: string;
  month: Date;
  color: string;
}

const props = defineProps<{
  data: AnticipatedCompletionDate & {
    startDate: Date;
    completionDate: Date;
  };
  scope: ScopeOfEngagement;
}>();
// const props = {
//   data: {
//     startDate: 1630454400000,
//     basisOfEstimateReport: 1633132800000,
//     costEstimation: 1635724800000,
//     probabilisticRiskAssessment: 1638403200000,
//     quantification: 1641081600000,
//     estimatedCompletionDate: 1643673600000
//   }
// };

const timeline: TimelineItem[] = [];

if (props.scope.basisOfEstimateReport) {
  timeline.push({
    text: "Basis of Estimate Report",
    month: props.data.basisOfEstimateReport,
    color: "#2C4C6E"
  });
}

if (props.scope.costEstimation) {
  timeline.push({
    text: "Cost Estimation",
    month: props.data.costEstimation,
    color: "#00A1D7"
  });
}

if (props.scope.probabilisticRiskAssessment) {
  timeline.push({
    text: "Probabilistic Risk Assessment",
    month: props.data.probabilisticRiskAssessment,
    color: "#7090B3"
  });
}

if (props.scope.quantification) {
  timeline.push({
    text: "Quantification",
    month: props.data.quantification,
    color: "#BDD8F6"
  });
}

timeline.sort((a, b) => a.month.getTime() - b.month.getTime());

timeline.unshift({
  text: "Start Date",
  month: props.data.startDate,
  color: "#c1c1c1"
});

timeline.push({
  text: "Completion Date",
  month: props.data.completionDate,
  color: "#c1c1c1"
});

const formatMonth = (month: Date) => moment(month).format("DD/MM/YYYY");
</script>

<template>
  <div>
    <span class="card-title">Anticipated Completion Date </span>
    <div class="commercial-information__card-content">
      <v-timeline
        side="start"
        class="scroll"
        direction="horizontal"
      >
        <v-timeline-item
          v-for="(item, i) in timeline"
          :key="i"
          size="small"
          :dot-color="item.color"
        >
          <template #opposite>
            <div v-text="formatMonth(item.month)" />
          </template>

          <span class="text">
            {{ item.text }}
          </span>
        </v-timeline-item>
      </v-timeline>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.commercial-information {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 10px;
    overflow-x: auto;

    & .text {
      color: #2c4c6e;
      font-size: 18px;
    }
  }
}
</style>
