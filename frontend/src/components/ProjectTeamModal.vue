<script lang="ts" setup>
import { computed } from "vue";

import type { ProjectTeam } from "@/types/project";

type TeamMember = ProjectTeam["projectLead"] & {
  position: string;
};
const props = defineProps<{
  member: TeamMember | undefined;
  isOpen: boolean;
  onClose: () => void;
}>();

const isProjectLead = computed(() => props.member?.position === "Project Lead");
</script>

<template>
  <div v-if="props.isOpen">
    <!-- Overlay -->
    <div class="fixed inset-0 bg-black opacity-50 z-10" />

    <!-- Modal -->
    <div class="fixed z-10 inset-0 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen">
        <div class="relative bg-white rounded-lg w-1/2">
          <div class="p-10 flex gap-10">
            <!-- Avatar -->
            <section class="w-[200px] min-w-[200px]">
              <div class="h-full w-auto flex items-center justify-center">
                <img
                  class="rounded"
                  :src="props.member?.avatar"
                  alt="Avatar Image"
                />
              </div>
            </section>

            <!-- Member Information -->
            <div class="flex flex-col gap-4">
              <div v-if="isProjectLead">
                <div
                  class="text-sm inline-block bg-gray-700 text-white px-4 py-1 rounded-full"
                >
                  Project Lead
                </div>
              </div>

              <section>
                <p class="font-semibold text-gray-900">
                  {{ props.member?.fullName }}
                </p>

                <p class="font-light text-gray-600">
                  {{ props.member?.email }}
                </p>
              </section>

              <p class="text-gray-900">
                {{ props.member?.bio }}
              </p>
            </div>
          </div>

          <button
            class="absolute top-0 right-0 my-4 mx-6"
            @click="props.onClose"
          >
            &times;
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
