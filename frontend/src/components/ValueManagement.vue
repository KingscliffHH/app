<script lang="ts" setup>
const props = defineProps<{ data: { opportunity: string }[] }>();

// const props = {
//   data: [
//     { opportunity: "Recycle site won material" },
//     { opportunity: "Replace concrete sharepath with asphalt solution" },
//     { opportunity: "Encapsulate ACM on site" },
//     { opportunity: "Recycle site won material" },
//     { opportunity: "Replace concrete sharepath with asphalt solution" }
//   ]
// };
</script>

<template>
  <div>
    <span class="card-title">Value Management Opportunities</span>
    <div class="value-management__card-content">
      <table>
        <thead>
          <tr>
            <th>Opportunity</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in props.data"
            :key="index"
          >
            <td>{{ item.opportunity }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.value-management {
  &__card-content {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 10px;
  }
}

table {
  flex: 1;

  & th:first-child {
    border-radius: 10px 10px 0 0;
  }

  & thead > tr {
    background-color: #f2f2f2;
  }

  & th {
    padding: 10px 20px;
    font-weight: 500;
    text-align: center;
    &:first-child {
      text-align: left;
    }
  }

  & td {
    padding: 10px 20px;
    font-weight: 400;
    text-align: center;
    &:first-child {
      text-align: left;
    }
  }

  & tbody {
    tr {
      border: 1px solid #e0e0e0;

      &:nth-child(even) {
        background-color: #f9fafb;
      }

      &:nth-child(odd) {
        background-color: #ffffff;
      }
    }
  }
}

@media (max-height: 1080px) or (max-width: 1920px) {
  table tbody tr td {
    padding: 5px 10px;
  }

  table thead tr th {
    font-size: 14px;
  }

  table tbody tr td {
    font-size: 12px;
  }
}
</style>
