<script setup lang="ts">
import { computed, ref } from "vue";
import { toast } from "vue3-toastify";

import type {
  Member,
  ProjectTeam,
  TotalProjectCostPerMilestone
} from "@/types/project";

import CIInput from "@/components/CIInput.vue";
import CompletionDateDialog from "@/components/projects/CompletionDateDialog.vue";
import ProjectSpecificInputEstimate from "@/components/ProjectSpecificInputEstimate.vue";
import ProjectSpecificInputQuantity from "@/components/ProjectSpecificInputQuantity.vue";
import ProjectTeamModal from "@/components/ProjectTeamModal.vue";

import { useMutation, useQuery } from "@/hooks/fetch";
import services from "@/services";
import { userRole, user } from "@/store/auth";

const props = defineProps<{
  id: string;
}>();

const {
  isLoading,
  refetch,
  data: project
} = useQuery({
  queryFn: () => services.projects.get(props.id)
});

const tab = ref("ci_input");

const { isLoading: isMarkingAsCompleted, mutate: markAsCompleted } =
  useMutation({
    mutationFn: (id: string, date: Date) =>
      services.projects.markAsCompleted(id, date),
    onSuccess: (data) => {
      console.info("onSuccess", data);
      toast.success("Success!", {
        autoClose: 2000
      });
    },
    onError: (err) => {
      if ("message" in err) {
        toast.error(err.message, {
          autoClose: 5000
        });
        return;
      }

      toast.error("Error!", {
        autoClose: 2000
      });
      console.info("onError", err);
    }
  });

const onConfirmMarkAsCompleted = async (date: Date) => {
  console.info("date", date);
  await markAsCompleted(props.id, date);
  refetch();
};

const canEdit = computed(() => {
  return userRole.value === "admin";
});

const canUpdateMetrics = computed(() => {
  const isAdmin = userRole.value === "admin";
  const isProjectLead =
    userRole.value === "member" &&
    user.value?.id === project.value?.team.projectLead.id;

  return project.value?.status !== "completed" && (isAdmin || isProjectLead);
});

const canMarkAsCompleted = computed(() => {
  const isAdmin = userRole.value === "admin";
  const isProjectLead =
    userRole.value === "member" &&
    user.value?.id === project.value?.team.projectLead.id;

  return project.value?.status !== "completed" && (isAdmin || isProjectLead);
});

const team = computed(() => {
  const withPosition = (x: Member, position: string) => ({ ...x, position });

  if (!project.value) return [];

  return [
    withPosition(project.value.team.projectLead, "Project Lead"),
    ...project.value.team.teamMembers.map((x) => withPosition(x, "Team Member"))
  ];
});

const showCompletionDateDialog = ref(false);

const { data: availableBenchmarks } = useQuery({
  queryFn: () => services.benchmarks.getAll()
});

type TeamMember = ProjectTeam["projectLead"] & {
  position: string;
};
const selectedMember = ref<TeamMember>();
const shouldShowModal = ref<boolean>(false);

const onOpen = (member: TeamMember): void => {
  selectedMember.value = member;
  shouldShowModal.value = true;
};

const onClose = (): void => {
  shouldShowModal.value = false;
};

const currentMilestone = (
  costsPerMilestone: TotalProjectCostPerMilestone[]
): string => {
  for (const item of costsPerMilestone) {
    if (item.currentMilstone) {
      return item.levelOfDesign;
    }
  }

  return "";
};
</script>

<template>
  <main class="main flex flex-col h-full ml-[80px] p-4">
    <div class="border-2 border-gray-200 rounded-lg bg-white">
      <div class="flex justify-between p-4">
        <div class="flex align-center gap-4">
          <span class="text-3xl font-bold text-blue-950">
            {{ project?.name }}
          </span>
          <picture
            v-for="member in team"
            :key="member?.email"
            class="size-12 rounded-full overflow-hidden cursor-pointer"
            @click="onOpen(member)"
          >
            <img
              :src="member?.avatar"
              :alt="member?.fullName"
            />
            <v-tooltip
              activator="parent"
              location="bottom"
            >
              {{ member?.fullName }}
            </v-tooltip>
          </picture>
        </div>

        <div>
          <v-tabs
            v-model="tab"
            align-tabs="center"
            color="#2c4c6e"
          >
            <v-tab value="ci_input">C&I Input</v-tab>
            <v-tab value="psi_estimate">Proj Specific Input (Est)</v-tab>
            <v-tab value="psi_quantity">Proj Specific Input (Qty)</v-tab>
          </v-tabs>
        </div>

        <div class="flex align-center gap-4">
          <router-link
            v-if="canEdit"
            :to="`/projects/${props.id}/edit`"
          >
            <v-btn
              color="#2c4c6e"
              variant="tonal"
              size="large"
            >
              <i class="material-icons-round">edit</i>
              <v-tooltip
                activator="parent"
                location="start"
              >
                Edit
              </v-tooltip>
            </v-btn>
          </router-link>
          <router-link
            v-if="canUpdateMetrics"
            :to="`/projects/${props.id}/metrics`"
          >
            <v-btn
              variant="tonal"
              size="large"
            >
              <i class="material-icons-round">stacked_bar_chart</i>
              <v-tooltip
                activator="parent"
                location="start"
              >
                Update Metrics
              </v-tooltip>
            </v-btn>
          </router-link>
          <v-btn
            :color="project?.status === 'completed' ? 'green' : ''"
            variant="tonal"
            size="large"
            :loading="isMarkingAsCompleted"
            :disabled="!canMarkAsCompleted"
            @click="showCompletionDateDialog = true"
          >
            <i class="material-icons-round">check_circle</i>
            <v-tooltip
              activator="parent"
              location="start"
            >
              Mark as Completed
            </v-tooltip>
          </v-btn>
        </div>
      </div>

      <div class="flex justify-between border-t-2 border-gray-2 p-4">
        <div>
          <span class="text-md mr-1">Region:</span>
          <span class="text-sm text-slate-500">{{ project?.region }}</span>
        </div>

        <div>
          <span class="text-md mr-1">Client:</span>
          <span class="text-sm text-slate-500">{{ project?.client }}</span>
        </div>

        <div>
          <span class="text-md mr-1">C&I Project Nº:</span>
          <span class="text-sm text-slate-500">{{
            project?.ciProjectNumber
          }}</span>
        </div>

        <div>
          <span class="text-md mr-1">Client Project Nº:</span>
          <span class="text-sm text-slate-500">{{
            project?.clientProjectNumber
          }}</span>
        </div>

        <div>
          <span class="text-md mr-1">Milestone:</span>
          <span class="text-sm text-slate-500">
            {{
              currentMilestone(
                project?.metrics.totalProjectCostPerMilestone || []
              )
            }}
          </span>
        </div>

        <div>
          <span class="text-md mr-1">Client Representative:</span>
          <span class="text-sm text-slate-500">{{
            project?.clientRepresentative?.fullName
          }}</span>
        </div>

        <div>
          <span class="text-md mr-1">Role:</span>
          <span class="text-sm text-slate-500">{{
            project?.clientRepresentative?.clientRole
          }}</span>
        </div>
      </div>
    </div>

    <v-window
      v-model="tab"
      style="height: 100%; overflow: auto"
    >
      <v-window-item value="ci_input">
        <CIInput
          v-if="!isLoading"
          :project="project"
        />
      </v-window-item>

      <v-window-item value="psi_estimate">
        <ProjectSpecificInputEstimate
          v-if="!isLoading && availableBenchmarks !== null"
          :project="project"
          :available-benchmarks="availableBenchmarks"
        />
      </v-window-item>

      <v-window-item value="psi_quantity">
        <ProjectSpecificInputQuantity
          v-if="!isLoading"
          :project="project"
        />
      </v-window-item>
    </v-window>
  </main>
  <CompletionDateDialog
    v-model="showCompletionDateDialog"
    @on-confirm="onConfirmMarkAsCompleted"
  />
  <ProjectTeamModal
    :member="selectedMember"
    :is-open="shouldShowModal"
    :on-close="onClose"
  />
</template>

<style lang="scss">
.main {
  background-color: #f9f9f9;
}

.v-window-item.v-window-item--active {
  height: 100%;
  display: flex;
}
</style>
