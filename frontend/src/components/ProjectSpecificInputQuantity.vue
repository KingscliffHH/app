<script lang="ts" setup>
import { Project } from "@/types/project";

import DesignPackages from "./DesignPackages.vue";
import ProjectPackages from "./ProjectPackages.vue";

const props = defineProps<{
  project: Project;
}>();
</script>

<template>
  <div class="grid-container">
    <DesignPackages
      class="component component1"
      :data="props.project.metrics.designPackages"
      :milestones="props.project.scope.numberOfMilestones"
    />
    <ProjectPackages
      class="component component2"
      :data="props.project.metrics.packages || []"
      :milestones="props.project.scope.numberOfMilestones"
    />
  </div>
</template>

<style lang="scss" scoped>
.grid-container {
  flex: 1;
  gap: 15px;
  display: grid;
  margin-top: 15px;
  grid-template-columns: repeat(1, 1fr);
  grid-template-rows: auto;
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
  grid-area: 1 / 1 / 2 / 5;
}

.component2 {
  grid-area: 2 / 1 / 3 / 5;
}
</style>
