<script setup lang="ts">
import { watch } from "vue";
import { CurrencyDisplay, useCurrencyInput } from "vue-currency-input";

import BaseInput from "./BaseInput.vue";

const props = defineProps<{
  modelValue: any;
  min?: number;
  max?: number;
}>();

watch(
  () => props.modelValue,
  (value) => {
    setValue(value);
  }
);

const { inputRef, formattedValue, setValue } = useCurrencyInput({
  currency: "AUD",
  currencyDisplay: CurrencyDisplay.hidden,
  hideCurrencySymbolOnFocus: false,
  hideGroupingSeparatorOnFocus: false,
  precision: 2,
  valueRange: {
    min: props.min || 0,
    max: props.max
  }
});
</script>

<template>
  <BaseInput
    ref="inputRef"
    v-model="formattedValue"
    class="number"
    name="teste"
  />
</template>
