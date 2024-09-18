<script setup lang="ts">
import { computed, ref } from "vue";

import type { Package } from "@/types/project";

import BaseInput from "@/components/base/BaseInput.vue";
import BaseTrueFalseInput from "@/components/base/BaseTrueFalseInput.vue";

const props = defineProps<{
  modelValue: Package[][] | null;
  disabled?: boolean;
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
const milestonePackages = computed(() => {
  if (selectedMilestone.value === -1) {
    return [];
  }

  return props.modelValue?.map((pkg) => pkg[selectedMilestone.value]) || [];
});
</script>

<template>
  <section class="mb-8">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">Packages:</span>
    </div>
    <div class="w-25 my-4">
      <v-select
        v-model="selectedMilestone"
        hide-details
        color="#1A3C5B"
        label="Select Milestone"
        :items="milestones"
      />
    </div>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left rtl:text-right text-gray-500">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50">
          <tr>
            <th
              scope="col"
              class="px-6 py-3"
            ></th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Progress %
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Second Stage (Senior QS) QA<br />Review Completed (50%)
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Final QA Review<br />Complated (90%)
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            >
              Submited (100%)
            </th>
            <th
              scope="col"
              class="px-6 py-3"
            ></th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(pkg, i) in milestonePackages"
            :key="i"
            class="odd:bg-white even:bg-gray-50 border-b"
          >
            <th
              scope="row"
              class="px-6 py-2"
            >
              <BaseInput
                v-model.number="pkg.description"
                name="description"
                class="bg-inherit border-none"
                readonly
                :disabled="props.disabled || pkg.description === ''"
              />
            </th>
            <td class="px-6 py-2">
              <BaseInput
                v-model.number="pkg.progress"
                type="number"
                min="0"
                max="100"
                name="progress"
                class="bg-inherit border-none"
                :disabled="props.disabled || pkg.description === ''"
              />
            </td>
            <td class="px-6 py-2">
              <BaseTrueFalseInput
                v-model="pkg.secondStageQA"
                name="secondStageQA"
                label-false="No"
                label-true="Yes"
                :disabled="props.disabled || pkg.description === ''"
              />
            </td>
            <td class="px-6 py-2">
              <BaseTrueFalseInput
                v-model="pkg.finalQAReview"
                name="finalQAReview"
                label-false="No"
                label-true="Yes"
                :disabled="props.disabled || pkg.description === ''"
              />
            </td>
            <td class="px-6 py-2">
              <BaseTrueFalseInput
                v-model="pkg.submitted"
                name="submitted"
                label-false="No"
                label-true="Yes"
                :disabled="props.disabled || pkg.description === ''"
              />
            </td>
            <td></td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
