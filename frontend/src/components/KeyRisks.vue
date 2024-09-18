<script lang="ts" setup>
import { KeyRisk } from "@/types/project";

const props = defineProps<{ data: KeyRisk[] }>();
</script>

<template>
  <div>
    <span class="card-title">Key Risks</span>
    <div class="key-risks__card-content">
      <table>
        <thead>
          <tr>
            <th>Description</th>
            <th>Score</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in props.data"
            :key="index"
          >
            <td>{{ item.description }}</td>
            <td :class="item.score.toLowerCase()">{{ item.score }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.key-risks {
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
    border-radius: 10px 0 0 0;
  }

  & th:last-child {
    border-radius: 0 10px 0 0;
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
        text-transform: uppercase;
        &:nth-child(2) {
          color: white;
          font-weight: bold;
          text-align: center;

          &.high {
            background: rgb(242, 173, 173);
            background: linear-gradient(
              0deg,
              rgba(242, 173, 173, 1) 0%,
              rgba(233, 137, 137, 1) 100%
            );
            color: #9c0006;
          }
          &.medium {
            background: rgb(233, 225, 139);
            background: linear-gradient(
              0deg,
              rgba(233, 225, 139, 1) 0%,
              rgba(242, 217, 126, 1) 100%
            );
            color: #876b0d;
          }
          &.low {
            background: rgb(141, 214, 116);
            background: linear-gradient(
              0deg,
              rgba(141, 214, 116, 1) 0%,
              rgba(112, 177, 89, 1) 100%
            );
            color: #006100;
          }
        }
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
