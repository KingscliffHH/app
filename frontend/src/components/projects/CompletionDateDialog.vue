<script setup lang="ts">
import { ref, computed } from "vue";

const props = defineProps<{
  modelValue: boolean;
}>();

const emits = defineEmits<{
  (e: "update:model-value", val: boolean): void;
  (e: "on-confirm", date: Date): void;
}>();

const dialog = computed({
  get: () => props.modelValue,
  set: (val: boolean) => emits("update:model-value", val)
});

const date = ref(new Date());

const onConfirm = () => {
  emits("on-confirm", date.value);
  dialog.value = false;
};
</script>

<template>
  <v-row justify="center">
    <v-dialog
      v-model="dialog"
      width="350"
    >
      <v-card>
        <v-card-title class="text-h5">Completion Date:</v-card-title>

        <v-date-picker
          v-model="date"
          min="2024-04-04"
        ></v-date-picker>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            color="green darken-1"
            @click="dialog = false"
          >
            Cancel
          </v-btn>

          <v-btn
            color="green darken-1"
            @click="onConfirm"
          >
            Confirm
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>
