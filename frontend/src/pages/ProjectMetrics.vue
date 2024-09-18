<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";

import { Metrics } from "@/types/project";

import AnticipatedCompletionDate from "@/components/projects/metrics/AnticipatedCompletionDate.vue";
import CommercialInformation from "@/components/projects/metrics/CommercialInformation.vue";
import DesignPackages from "@/components/projects/metrics/DesignPackages.vue";
import KeyCostDrivers from "@/components/projects/metrics/KeyCostDrivers.vue";
import KeyRisks from "@/components/projects/metrics/KeyRisks.vue";
import OptionOutturnCost from "@/components/projects/metrics/OptionOutturnCost.vue";
import ProgressToDate from "@/components/projects/metrics/ProgressToDate.vue";
import ProjectBenchmarking from "@/components/projects/metrics/ProjectBenchmarking.vue";
import ProjectPackages from "@/components/projects/metrics/ProjectPackages.vue";
import TotalProjectCost from "@/components/projects/metrics/TotalProjectCost.vue";
import ValueManagementOpportunities from "@/components/projects/metrics/ValueManagementOpportunities.vue";

import { useMutation, useQuery } from "@/hooks/fetch";
import services from "@/services";

const props = defineProps<{
  id: string;
}>();

const router = useRouter();

const {
  isLoading: fetching,
  data: project
  // refetch
} = useQuery({
  queryFn: () => services.projects.get(props.id)
});

const {
  // isLoading: saving,
  // isError,
  error,
  // isSuccess,
  mutate: save
} = useMutation({
  mutationFn: (id: string, metrics: Metrics) => {
    const validBenchmark = (b: any) => b.benchmarkId !== "";
    const benchmarking =
      metrics.benchmarking.benchmarks?.filter(validBenchmark);

    return services.projects.updateMetrics(id, {
      ...metrics,
      benchmarking: {
        ...metrics.benchmarking,
        benchmarks: benchmarking
      }
    });
  },
  onSuccess: (data) => {
    console.info("onSuccess", data);
    // toast.success(".");
    toast.success("Success!", {
      autoClose: 2000
    });
    router.push(`/projects/${props.id}`);
  },
  onError: (err) => {
    toast.error("Error!", {
      autoClose: 2000
    });
    console.info("onError", err);
  }
});

const activeTab = ref("content1");

const submit = async () => {
  save(props.id, {
    progressToDate: project.value.metrics.progressToDate,
    anticipatedCompletionDate: project.value.metrics.anticipatedCompletionDate,
    commercialInformation: project.value.metrics.commercialInformation,
    optionOutturnCosts: project.value.metrics.optionOutturnCosts,
    totalProjectCostPerMilestone:
      project.value.metrics.totalProjectCostPerMilestone,
    keyCostDriversBaseValue: project.value.metrics.keyCostDriversBaseValue,
    keyCostDrivers: project.value.metrics.keyCostDrivers,
    keyRisks: project.value.metrics.keyRisks,
    valueManagementOpportunities:
      project.value.metrics.valueManagementOpportunities,
    benchmarking: project.value.metrics.benchmarking,
    designPackages: project.value.metrics.designPackages,
    packages: project.value.metrics.packages
  });

  // router.push(`/projects/${props.id}`);
};

const onCancel = () => {
  router.back();
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">Update Metrics - {{ project?.name }}</h1>
      <div>
        <button
          class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded"
          type="button"
          @click="onCancel"
        >
          Cancel
        </button>
        <button
          class="px-4 py-1 ml-2 bg-blue-500 border border-blue-500 text-white font-semibold rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-blue-500"
          type="button"
          :disabled="fetching || project?.status === 'completed'"
          @click="submit"
        >
          Submit
        </button>
      </div>
    </section>
    <div
      id="tabs"
      class="flex flex-col items-left w-full"
    >
      <!-- Tab Headers -->
      <div class="flex border-b-2 border-gray-700">
        <button
          data-tab="content1"
          :class="activeTab === 'content1' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'content1'"
        >
          C&I Input
        </button>
        <button
          data-tab="content2"
          :class="activeTab === 'content2' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'content2'"
        >
          Proj. Specific Input
          <small class="ms-2 font-semibold text-gray-500 dark:text-gray-400"
            >(Est)</small
          >
        </button>
        <button
          data-tab="content3"
          :class="activeTab === 'content3' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'content3'"
        >
          Proj. Specific Input
          <small class="ms-2 font-semibold text-gray-500 dark:text-gray-400"
            >(Qty)</small
          >
        </button>
      </div>
      <!-- Tab Contents -->
      <form
        action="javascript:void(0);"
        autocomplete="off"
      >
        <div
          class="tab-content w-full bg-white shadow-md rounded-b-lg p-6 mb-4"
          :class="activeTab === 'content1' ? '' : 'hidden'"
        >
          <ProgressToDate
            v-if="!fetching"
            v-model="project.metrics.progressToDate"
            :scope="project.scope"
            :max-milestone="project.scope.numberOfMilestones"
            :disabled="project.status === 'completed'"
          />
          <AnticipatedCompletionDate
            v-if="!fetching"
            v-model="project.metrics.anticipatedCompletionDate"
            :disabled="project.status === 'completed'"
          />
          <CommercialInformation
            v-if="!fetching"
            v-model="project.metrics.commercialInformation"
            :disabled="project.status === 'completed'"
          />
        </div>
        <div
          class="tab-content w-full bg-white shadow-md rounded-b-lg p-6 mb-4"
          :class="activeTab === 'content2' ? '' : 'hidden'"
        >
          <OptionOutturnCost
            v-if="!fetching"
            v-model="project.metrics.optionOutturnCosts"
            :disabled="project.status === 'completed'"
          />
          <TotalProjectCost
            v-if="!fetching"
            v-model="project.metrics.totalProjectCostPerMilestone"
            :disabled="project.status === 'completed'"
          />
          <KeyCostDrivers
            v-if="!fetching"
            v-model:drivers="project.metrics.keyCostDrivers"
            v-model:base-value="project.metrics.keyCostDriversBaseValue"
            :disabled="project.status === 'completed'"
          />
          <KeyRisks
            v-if="!fetching"
            v-model="project.metrics.keyRisks"
            :disabled="project.status === 'completed'"
          />
          <ValueManagementOpportunities
            v-if="!fetching"
            v-model="project.metrics.valueManagementOpportunities"
            :disabled="project.status === 'completed'"
          />
          <ProjectBenchmarking
            v-if="!fetching"
            v-model="project.metrics.benchmarking"
            :project-name="project.name"
            :error="error"
            :disabled="project.status === 'completed'"
          />
        </div>
        <div
          class="tab-content w-full bg-white shadow-md rounded-b-lg p-6 mb-4"
          :class="activeTab === 'content3' ? '' : 'hidden'"
        >
          <DesignPackages
            v-if="!fetching"
            v-model="project.metrics.designPackages"
            v-model:packages="project.metrics.packages"
            :milestones="project.scope.numberOfMilestones"
            :disabled="project.status === 'completed'"
          />
          <ProjectPackages
            v-if="!fetching"
            v-model="project.metrics.packages"
            :milestones="project.scope.numberOfMilestones"
            :disabled="project.status === 'completed'"
          />
        </div>
      </form>
    </div>
  </main>
</template>

<style lang="scss">
.main {
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin-left: 80px;
  background-color: #f9f9f9;
  padding: 15px;
}
</style>
