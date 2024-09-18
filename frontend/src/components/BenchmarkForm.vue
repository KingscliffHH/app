<script setup lang="ts">
import { computed } from "vue";

import BaseInput from "./base/BaseInput.vue";
import BaseNumberInput from "./base/BaseNumberInput.vue";

const props = defineProps<{
  modelValue: any;
  error?: Record<string, string>;
  readonly: boolean;
}>();

const emits = defineEmits<{
  (e: "update:modelValue", value: any): void;
}>();

const data = computed({
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value)
});
</script>

<template>
  <div class="bg-white shadow-md rounded-lg p-6 mb-4">
    <div class="grid grid-cols-1 w-1/2 gap-4">
      <BaseInput
        v-model="data.name"
        label="Name"
        name="name"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />
      <BaseInput
        v-model="data.geographicLocation"
        label="Geographic Location"
        name="geographicLocation"
        type="text"
        :errors="props.error"
        :readonly="props.readonly"
      />
      <BaseNumberInput
        v-model="data.totalProjectCostP90"
        label="Total Project Cost (P90)"
        name="totalProjectCostP90"
        :min="0"
        prefix="$"
        :errors="props.error"
        :readonly="props.readonly"
      />
      <BaseNumberInput
        v-model="data.totalConstructionCostPerLaneKm"
        label="Total Construction Cost $/Lane km"
        name="totalConstructionCostPerLaneKm"
        :min="0"
        prefix="$"
        :errors="props.error"
        :readonly="props.readonly"
      />
      <BaseNumberInput
        v-model="data.cubicMetreRateForEarthworksPerM3"
        label="Cubic Metre Rate for Earthworks $/m3"
        name="cubicMetreRateForEarthworksPerM3"
        :min="0"
        prefix="$"
        :errors="props.error"
        :readonly="props.readonly"
      />
      <BaseNumberInput
        v-model="data.squareMetreRateForPavementPerBridgePerM2"
        label="Square Metre Rate for Pavement/Bridge $/m2"
        name="squareMetreRateForPavementPerBridgePerM2"
        :min="0"
        prefix="$"
        :errors="props.error"
        :readonly="props.readonly"
      />
    </div>
  </div>
</template>
