<script setup lang="ts">
import numeral from "numeral";
import { computed } from "vue";

import { type CommercialInformation } from "@/types/project";

const props = defineProps<{
  data: CommercialInformation;
}>();

const remainingFee = computed(() => {
  return (
    props.data?.ciContractedValue +
    props.data?.approvedVariationToDate -
    props.data?.ciAccrualToDate
  );
});
</script>

<template>
  <div class="grid grid-cols-2">
    <v-card
      subtitle="C&I Contracted Value"
      class="pb-2"
    >
      <p class="text-4xl font-bold px-4 text-right text-sky-700">
        {{ numeral(props.data?.ciContractedValue).format("$ 0,0.00") }}
      </p>
    </v-card>
    <v-card
      subtitle="Remaining Fee"
      class="pb-2"
    >
      <p class="text-4xl font-bold px-4 text-right text-sky-700">
        {{ numeral(remainingFee).format("$ 0,0.00") }}
      </p>
    </v-card>
    <v-card
      subtitle="Accrual To Date"
      class="pb-2"
    >
      <p class="text-4xl font-bold px-4 text-right text-sky-700">
        {{ numeral(props.data?.ciAccrualToDate).format("$ 0,0.00") }}
      </p>
    </v-card>
    <v-card
      v-if="props.data?.approvedVariationToDate > 0"
      subtitle="Approved Variation To Date"
      class="pb-2"
    >
      <p class="text-4xl font-bold px-4 text-right text-sky-700">
        {{ numeral(props.data?.approvedVariationToDate).format("$ 0,0.00") }}
      </p>
    </v-card>
  </div>
  <!-- <section class="mb">
    <div class="flex items-center mb-6">
      <span class="pr-4 text-lg font-medium text-gray-700">
        Commercial Information:
      </span>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 p-6">
        <div
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-md flex flex-col justify-between"
        >
          <h5 class="text-sm font-medium text-gray-700">
            C&I Contracted Value
          </h5>
          <p class="text-4xl font-semibold text-gray-900 text-right">
            $
            {{ data?.ciContractedValue }}
          </p>
        </div>
        <div
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-md flex flex-col justify-between"
        >
          <h5 class="text-sm font-medium text-gray-700">Accrual To Date</h5>
          <p class="text-4xl font-semibold text-gray-900 text-right">
            ${{ data?.ciAccrualToDate }}
          </p>
        </div>
        <div
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-md flex flex-col justify-between"
        >
          <h5 class="text-sm font-medium text-gray-700">Remaining Fee:</h5>
          <p class="text-4xl font-semibold text-gray-900 text-right">
            ${{ remainingFee }}
          </p>
        </div>
        <div
          v-if="data?.approvedVariationToDate !== 0"
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-md flex flex-col justify-between"
          :class="{ 'opacity-0': data?.approvedVariationToDate === 0 }"
        >
          <h5 class="text-sm font-medium text-gray-700">
            Approved Variation To Date
          </h5>
          <p class="text-4xl font-semibold text-gray-900 text-right">
            ${{ data?.approvedVariationToDate }}
          </p>
        </div>
      </div>
    </div>
  </section> -->
</template>
