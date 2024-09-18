<script lang="ts" setup>
import { Project } from "@/types/project";

import AccrualToDate from "./AccrualToDate.vue";
import AnticipatedCompletionDate from "./AnticipatedCompletionDate.vue";
import ApprovedVariationsToDate from "./ApprovedVariationsToDate.vue";
import CIContractedValue from "./CIContractedValue.vue";
import CommercialInformation from "./CommercialInformation.vue";
import NumberOfMilestones from "./NumberOfMilestones.vue";
import ProgressToDate from "./ProgressToDate.vue";
import RemainingFee from "./RemainingFee.vue";

const props = defineProps<{
  project: Project;
}>();
</script>

<template>
  <div class="grid-container">
    <CIContractedValue
      class="component component1"
      :value="props.project.metrics.commercialInformation.ciContractedValue"
    />
    <RemainingFee
      class="component component2"
      :value="
        props.project.metrics.commercialInformation.ciContractedValue +
        props.project.metrics.commercialInformation.approvedVariationToDate -
        props.project.metrics.commercialInformation.ciAccrualToDate
      "
    />
    <AccrualToDate
      class="component component3"
      :value="props.project.metrics.commercialInformation.ciAccrualToDate"
    />
    <ApprovedVariationsToDate
      class="component component4"
      :value="
        props.project.metrics.commercialInformation.approvedVariationToDate
      "
    />

    <CommercialInformation
      class="component component5"
      :remaining-fee="
        props.project.metrics.commercialInformation.ciContractedValue +
        props.project.metrics.commercialInformation.approvedVariationToDate -
        props.project.metrics.commercialInformation.ciAccrualToDate
      "
      :accrual-to-date="
        props.project.metrics.commercialInformation.ciAccrualToDate
      "
    />

    <NumberOfMilestones
      class="component component6"
      :milestone="{
        completed: props.project.metrics.progressToDate.numberOfMilestones,
        total: props.project.scope.numberOfMilestones
      }"
    />

    <AnticipatedCompletionDate
      class="component component7"
      :data="{
        ...props.project.metrics.anticipatedCompletionDate,
        startDate: props.project.startDate,
        completionDate:
          props.project.status === 'completed'
            ? props.project.completionDate
            : props.project.scope.estimatedCompletionDate
      }"
      :scope="props.project.scope"
    />
    <ProgressToDate
      class="component component8"
      :progress-to-date="props.project.metrics.progressToDate"
      :scope="props.project.scope"
    />
  </div>
</template>

<style lang="scss" scoped>
.grid-container {
  flex: 1;
  gap: 15px;
  display: grid;
  margin-top: 15px;
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: 12vh 12vh 220px auto;
}

.component {
  display: flex;
  padding: 15px;
  border-radius: 18px;
  flex-direction: column;
  background-color: #ffffff;
  border: 2px solid #eff0f6;
}

.component1 {
  grid-area: 1 / 1 / 2 / 2;
}
.component2 {
  grid-area: 1 / 2 / 2 / 3;
}
.component3 {
  grid-area: 2 / 1 / 3 / 2;
}
.component4 {
  grid-area: 2 / 2 / 3 / 3;
}
.component5 {
  grid-area: 1 / 3 / 3 / 4;
}
.component6 {
  grid-area: 1 / 4 / 3 / 5;
}
.component7 {
  grid-area: 3 / 1 / 3 / 5;
}
.component8 {
  grid-area: 4 / 1 / 4 / 5;
}
</style>
