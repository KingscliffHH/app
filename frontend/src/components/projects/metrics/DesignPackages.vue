<script setup lang="ts">
import { computed, onMounted, unref } from "vue";

import type { DesignPackages, DesignPackage, Package } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseSelect from "@/components/base/BaseSelect.vue";

const props = defineProps<{
  modelValue: DesignPackages | null;
  packages: Package[][] | null;
  milestones: number;
  disabled?: boolean;
}>();

const emits = defineEmits(["update:model-value", "update:packages"]);

const EmptyObject = (description: string = ""): DesignPackage => {
  return {
    description,
    milestones: Array.from({ length: props.milestones }, () => "")
  };
};

const limit = 15;
const data = computed({
  get: () =>
    props.modelValue || {
      levelOfDesign: EmptyObject("Level of Design"),
      packages: [
        EmptyObject("Package 1"),
        EmptyObject("Package 2"),
        EmptyObject("Package 3")
      ]
    },
  set: (value) => emits("update:model-value", value)
});

const packages = computed({
  get: () => props.packages || [],
  set: (value) => emits("update:packages", value)
});

const canAddMore = computed(() => data.value.packages.length < limit);
const addMore = () => {
  data.value.packages = [
    ...data.value.packages,
    {
      description: "",
      milestones: Array.from({ length: props.milestones }, () => "")
    }
  ];

  packages.value = [
    ...packages.value,
    Array.from({ length: props.milestones }, () => ({
      description: "",
      progress: 0,
      secondStageQA: false,
      finalQAReview: false,
      submitted: false
    }))
  ];
};

const removeItem = (index: number) => {
  console.info("remove item", index);
  data.value.packages = data.value.packages.filter((_, i) => i !== index);
  packages.value = packages.value.filter((_, i) => i !== index);
};

const updatePackage = (pkgIdx: number, milestoneIdx: number, value: string) => {
  console.info("updatePackage", pkgIdx, milestoneIdx, value);
  console.info(unref(packages));
  packages.value[pkgIdx][milestoneIdx].description = value;
};

onMounted(() => {
  if (!props.modelValue) {
    data.value = {
      levelOfDesign: EmptyObject("Level of Design"),
      packages: Array.from({ length: 3 }, EmptyObject)
    };

    packages.value = Array.from({ length: 3 }, () =>
      Array.from({ length: props.milestones }, () => ({
        description: "",
        progress: 0,
        secondStageQA: false,
        finalQAReview: false,
        submitted: false
      }))
    );
  }
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <h2 class="pr-4 text-lg font-medium text-gray-700">Design Packages:</h2>
    </div>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg mb-8">
      <table
        class="w-full table-fixed text-sm text-left rtl:text-right text-gray-500"
      >
        <thead class="text-xs text-gray-700 uppercase bg-gray-50">
          <tr>
            <th
              scope="col"
              class="sticky left-0 z-10 px-6 py-2 bg-gray-50 w-[160px]"
            ></th>
            <th
              v-for="i in props.milestones"
              :key="i"
              scope="col"
              class="px-6 py-3 w-[200px]"
            >
              Milestone {{ i }}
            </th>
            <th
              scope="col"
              class="px-6 py-3 w-[60px]"
            ></th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-if="data.levelOfDesign"
            class="odd:bg-white even:bg-gray-50 border-b"
          >
            <th
              scope="row"
              class="sticky left-0 z-10 px-6 py-2 bg-inherit"
            >
              {{ data.levelOfDesign.description }}
            </th>

            <td
              v-for="mIdx in props.milestones"
              :key="mIdx"
              class="px-6 py-2"
            >
              <BaseInput
                v-model="data.levelOfDesign.milestones[mIdx - 1]"
                name="milestone"
                :disabled="props.disabled"
              />
            </td>
            <td></td>
          </tr>
          <tr
            v-for="(pkg, i) in data.packages"
            :key="i"
            class="odd:bg-white even:bg-gray-50 border-b"
          >
            <th
              scope="row"
              class="sticky left-0 z-10 px-6 py-2 bg-inherit"
            >
              Package {{ i + 1 }}
            </th>

            <td
              v-for="mIdx in props.milestones"
              :key="mIdx"
              class="px-6 py-2"
            >
              <BaseSelect
                v-model="pkg.milestones[mIdx - 1]"
                name="milestone"
                :options="[
                  '',
                  'Architectural',
                  'Base Building',
                  'Bridges',
                  'Building Electrical',
                  'Building Structure',
                  'Canopy Structure',
                  'Communication Systems',
                  'Culverts and Stormwater Structures',
                  'Demolition',
                  'Earlyworks',
                  'Earthing, Bonding & Electrolysis',
                  'Earthworks',
                  'Electrical & Lighting',
                  'Fire Detection Systems',
                  'Fire Services',
                  'Fire Suppression Systems',
                  'Geometry/Cross Sections',
                  'Geotechnical',
                  'Heating, Ventilation and Air-Conditioning (HVAC)',
                  'Hydraulic/Plumbing',
                  'ITS',
                  'LV Electrical & Lighting System',
                  'Landscaping',
                  'Multi-Discipline Services (Mechanical)',
                  'Pavements & Kerbs',
                  'Piling',
                  'Property Works',
                  'Retaining Walls',
                  'Road Alignment',
                  'Road Furniture',
                  'Signage & Deliniation',
                  'Staging Plans',
                  'Stormwater',
                  'Structural Walls',
                  'Structures',
                  'TCS Design',
                  'Temporary Works',
                  'Tunnel Ventilation',
                  'Tunnels',
                  'Utilities'
                ]"
                :disabled="props.disabled"
                @update:model-value="updatePackage(i, mIdx - 1, $event)"
              />
            </td>
            <td class="sticky right-0 z-10 bg-inherit">
              <v-btn
                class="text-red-500"
                variant="outlined"
                icon="mdi-delete"
                :disabled="props.disabled"
                @click="removeItem(i)"
              >
                <i class="text-base material-icons-round">delete</i>
              </v-btn>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <button
      class="mt-4 px-4 py-2 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-blue-500"
      :disabled="!canAddMore || disabled"
      @click="addMore"
    >
      + Add More
    </button>
  </section>
</template>
