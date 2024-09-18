<script setup lang="ts">
const props = defineProps<{
  value: number;
}>();

const format = (number: number) => {
  return new Intl.NumberFormat("en-AU", {
    style: "currency",
    currency: "AUD"
  })
    .format(number)
    .replace("$", "");
};
</script>

<template>
  <div>
    <span class="card-title">C&I Contracted Value</span>
    <div class="ci-contracted-value__card-content">
      <span>$</span>
      <span>{{ format(props.value) }}</span>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.ci-contracted-value {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;

    & span:nth-child(1) {
      font-size: 30px;
      font-weight: 400;
      color: #496895;
      margin-right: 10px;
    }

    span {
      font-size: 45px;
      font-weight: bold;
      color: #496895;
    }
  }
}
</style>
