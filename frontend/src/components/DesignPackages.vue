<script setup lang="ts">
import type { DesignPackages } from "@/types/project";

const props = defineProps<{
  data: DesignPackages | null;
  milestones: number;
}>();
</script>

<template>
  <div>
    <span class="card-title">Design Packages</span>
    <div class="design-packages__card-content">
      <table>
        <thead>
          <tr>
            <th />
            <th
              v-for="i in milestones"
              :key="i"
            >
              Milestone {{ i }}
            </th>
          </tr>
        </thead>
        <tbody v-if="props.data">
          <tr>
            <td>
              {{ props.data.levelOfDesign.description }}
            </td>
            <td
              v-for="mIdx in milestones"
              :key="mIdx"
            >
              {{ props.data.levelOfDesign.milestones[mIdx - 1] }}
            </td>
          </tr>
          <tr
            v-for="(pkg, i) in props.data.packages"
            :key="i"
          >
            <td>
              {{ `Packages ${i + 1}` }}
            </td>
            <td
              v-for="mIdx in milestones"
              :key="mIdx"
            >
              {{ pkg.milestones[mIdx - 1] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.design-packages {
  &__card-content {
    flex: 1;
    display: flex;
    padding: 10px;
  }
}

table {
  flex: 1;

  & th {
    padding: 10px 20px;
    font-weight: 500;
    text-align: center;

    &:first-child {
      border-radius: 10px 0 0 0;
    }

    &:last-child {
      border-radius: 0 10px 0 0;
    }
  }

  & thead > tr {
    background-color: #f2f2f2;
  }

  & td {
    padding: 10px 20px;
    font-weight: 400;

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

      td {
        color: #4d4d4d;

        &:not(:nth-child(1)) {
          text-align: center;
        }
      }
    }
  }
}

@media (max-height: 1080px) or (max-width: 1920px) {
  td {
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
