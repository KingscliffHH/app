<script setup lang="ts">
import { ref, computed } from "vue";
import { toast } from "vue3-toastify";

import { Project, TotalProjectCostPerMilestone } from "@/types/project";

import CompletionDateDialog from "@/components/projects/CompletionDateDialog.vue";

import { useMutation, useQuery } from "@/hooks/fetch";
import services from "@/services";
import { user, userRole } from "@/store/auth";

const {
  // loading,
  refetch,
  data: projects
} = useQuery({
  queryFn: () => services.projects.getAll(),
  onError: (err) => {
    console.info("onError", err);
    if ("message" in err) {
      toast.error(err.message, {
        autoClose: 5000
      });
      return;
    }

    toast.error("Error fetching projects!", {
      autoClose: 2000
    });
  }
});

const inProgressProjects = computed(() => {
  return projects.value?.filter((project) => project.status !== "completed");
});

const completedProjects = computed(() => {
  return projects.value?.filter((project) => project.status === "completed");
});
// console.info(projects.value);

const {
  // isLoading: saving,
  // isError,
  // error,
  // isSuccess,
  mutate: deleteProject
} = useMutation({
  mutationFn: (id: string) => services.projects.delete(id),
  onSuccess: (data) => {
    console.info("onSuccess", data);
    toast.success("Success!", {
      autoClose: 2000
    });
  },
  onError: (err) => {
    console.info("onError", err);
    if ("message" in err) {
      toast.error(err.message, {
        autoClose: 5000
      });
      return;
    }

    toast.error("Error!", {
      autoClose: 2000
    });
  }
});

const confirmDelete = async (item: Project) => {
  if (confirm("Are you sure you want to delete this item?")) {
    await deleteProject(item.id);
    refetch();
  }
};

const showCompletionDateDialog = ref(false);
const markingAsCompleted = ref<Project | null>(null);
const onConfirmMarkAsCompleted = async (date: Date) => {
  // if (confirm("Are you sure you want to mark this project as completed?")) {
  if (!markingAsCompleted.value) {
    console.error("No project selected");
    return;
  }

  await services.projects.markAsCompleted(markingAsCompleted.value.id, date);
  refetch();
  markingAsCompleted.value = null;
  // }
};

// const canMarkAsCompleted = computed(() => {
//   return (
//     userRole.value === "admin" ||
//     (userRole.value === "member" &&
//       user.value?.email === project.value.team.projectLead.email)
//   );
// });

const formatDate = (date: Date) => {
  return new Date(date).toLocaleDateString();
};

const activeTab = ref("in_progress");

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

const canEdit = computed(() => {
  return userRole.value === "admin";
});

const canUpdateMetrics = (project: Project) => {
  const isAdmin = userRole.value === "admin";
  const isProjectLead =
    userRole.value === "member" &&
    user.value?.id === project.team.projectLead.id;

  return project.status !== "completed" && (isAdmin || isProjectLead);
};

const canMarkAsCompleted = (project: Project): boolean => {
  const isAdmin = userRole.value === "admin";
  const isProjectLead =
    userRole.value === "member" &&
    user.value?.id === project.team.projectLead.id;

  return project.status !== "completed" && (isAdmin || isProjectLead);
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">Projects</h1>
      <router-link
        v-if="userRole === 'admin'"
        class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 py-1 border border-blue-500 hover:border-transparent rounded"
        to="/projects/new"
        hx-replace-url="true"
        hx-swap="innerHTML"
        hx-target="#main"
        hx-trigger="click"
      >
        + New
      </router-link>
    </section>
    <div
      id="tabs"
      class="flex flex-col items-left w-full"
    >
      <!-- Tab Headers -->
      <div class="flex border-b-2 border-gray-700">
        <button
          :class="activeTab === 'in_progress' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'in_progress'"
        >
          In Progress
        </button>

        <button
          :class="activeTab === 'completed' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'completed'"
        >
          Completed
        </button>
      </div>
      <div
        class="w-full bg-white shadow-md rounded-b-lg overflow-hidden"
        :class="activeTab === 'in_progress' ? '' : 'hidden'"
      >
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Lead
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Client
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Region
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Progress
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Start Date
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Completion Date
                <small class="ms-2 font-semibold text-gray-500">(est)</small>
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in inProgressProjects"
              :key="item.id"
              class="odd:bg-white even:bg-gray-50 border-b"
            >
              <td
                scope="row"
                class="px-6 py-2"
              >
                {{ item.name }}
              </td>
              <td class="px-6 py-2">{{ item.team.projectLead.fullName }}</td>
              <td class="px-6 py-2">{{ item.client }}</td>
              <td class="px-6 py-2">{{ item.region }}</td>
              <td class="px-6 py-2">
                {{
                  currentMilestone(
                    item.metrics.totalProjectCostPerMilestone || []
                  )
                }}
              </td>
              <td class="px-6 py-2">{{ formatDate(item.startDate) }}</td>
              <td class="px-6 py-2">
                {{ formatDate(item.scope.estimatedCompletionDate) }}
              </td>
              <td class="px-6 py-2 flex items-center space-x-2">
                <router-link :to="`/projects/${item.id}`">
                  <button
                    class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 border border-blue-500 hover:border-transparent rounded"
                  >
                    View Details
                  </button>
                </router-link>

                <button
                  v-if="canMarkAsCompleted(item)"
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-700 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  title="Mark as Completed"
                  @click="
                    markingAsCompleted = item;
                    showCompletionDateDialog = true;
                  "
                >
                  <i class="text-base material-icons-round">check</i>
                </button>

                <router-link
                  v-if="canEdit"
                  :to="`/projects/${item.id}/edit`"
                >
                  <button
                    class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-700 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                    title="Edit"
                  >
                    <span class="text-base material-icons-round">edit</span>
                  </button>
                </router-link>
                <router-link
                  v-if="canUpdateMetrics(item)"
                  :to="`/projects/${item.id}/metrics`"
                >
                  <button
                    class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-700 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                    hx-get='{ "/projects/" + item.ID.Hex() + "/metrics" }'
                    hx-replace-url="true"
                    hx-swap="innerHTML"
                    hx-target="#main"
                    hx-trigger="click"
                    title="Metrics"
                  >
                    <span class="text-base material-icons-round"
                      >stacked_bar_chart</span
                    >
                  </button>
                </router-link>
                <button
                  v-if="userRole === 'admin'"
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-red-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  title="Delete"
                  @click="confirmDelete(item)"
                >
                  <span class="text-base material-icons-round">delete</span>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div
        class="w-full bg-white shadow-md rounded-b-lg overflow-hidden"
        :class="activeTab === 'completed' ? '' : 'hidden'"
      >
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Lead
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Client
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Region
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Progress
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Start Date
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Completion Date
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in completedProjects"
              :key="item.id"
              class="odd:bg-white even:bg-gray-50 border-b"
            >
              <td
                scope="row"
                class="px-6 py-2"
              >
                {{ item.name }}
              </td>
              <td class="px-6 py-2">{{ item.team.projectLead.fullName }}</td>
              <td class="px-6 py-2">{{ item.client }}</td>
              <td class="px-6 py-2">{{ item.region }}</td>
              <td class="px-6 py-2">50%</td>
              <td class="px-6 py-2">{{ formatDate(item.startDate) }}</td>
              <td class="px-6 py-2">
                {{ formatDate(item.completionDate) }}
              </td>
              <td class="px-6 py-2 flex items-center space-x-2">
                <router-link :to="`/projects/${item.id}`">
                  <button
                    class="hover:bg-blue-500 text-blue-700 font-semibold hover:text-white px-4 border border-blue-500 hover:border-transparent rounded"
                  >
                    View Details
                  </button>
                </router-link>

                <button
                  v-if="userRole === 'admin'"
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-red-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  title="Delete"
                  @click="confirmDelete(item)"
                >
                  <span class="text-base material-icons-round">delete</span>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
  <CompletionDateDialog
    v-model="showCompletionDateDialog"
    @on-confirm="onConfirmMarkAsCompleted"
  />
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
