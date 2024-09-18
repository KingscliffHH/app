<script setup lang="ts">
import { computed, ref } from "vue";
import { toast } from "vue3-toastify";

import { type User } from "@/types/user";

import BaseButtonOutlined from "@/components/base/BaseButtonOutlined.vue";

import { useMutation, useQuery } from "@/hooks/fetch";
import services from "@/services";

const {
  // isLoading,
  data: users
} = useQuery({
  queryFn: () => services.users.getAll()
});
// console.info(projects.value);

const members = computed(() => {
  return users.value?.filter((user) => user.type === "member");
});

const clients = computed(() => {
  return users.value?.filter((user) => user.type === "client");
});

const activeTab = ref("members");

const {
  // isLoading: saving,
  // isError,
  // error,
  // isSuccess,
  mutate: deleteUser
} = useMutation({
  mutationFn: (id: string) => services.users.delete(id),
  onSuccess: (data) => {
    console.info("onSuccess", data);
    toast.success("Marked for deletion!", {
      autoClose: 2000
    });
  },
  onError: () => {
    console.info("onError");
    toast.error("Error deleting user!", {
      autoClose: 2000
    });
  }
});

const confirmDelete = async (user: User) => {
  if (confirm("Are you sure you want to delete this item?")) {
    await deleteUser(user.id);
    users.value = users.value.filter(({ id }) => id !== user.id);
  }
};
</script>

<template>
  <main class="main">
    <section class="flex justify-between pb-4">
      <h1 class="text-xl font-bold">System Users</h1>
      <router-link to="/users/new">
        <BaseButtonOutlined
          color="primary"
          label="+ New"
        />
      </router-link>
    </section>
    <div
      id="tabs"
      class="flex flex-col items-left w-full"
    >
      <!-- Tab Headers -->
      <div class="flex border-b-2 border-gray-700">
        <button
          :class="activeTab === 'members' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'members'"
        >
          C&I
        </button>

        <button
          :class="activeTab === 'clients' ? 'text-white bg-gray-700' : ''"
          class="tab-header border border-b-0 border-gray-700 focus:outline-none px-4 py-2 rounded-t-lg"
          @click="activeTab = 'clients'"
        >
          Clients
        </button>
      </div>
      <!-- Tab Contents -->
      <div
        class="tab-content w-full bg-white shadow-md rounded-b-lg overflow-hidden"
        :class="activeTab === 'members' ? '' : 'hidden'"
      >
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Avatar
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Email
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Last Access
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in members"
              :key="item.id"
              class="odd:bg-white even:bg-gray-50 border-b"
            >
              <td
                scope="row"
                class="px-6 py-2"
              >
                <picture>
                  <img
                    :src="item.avatar"
                    :alt="item.fullName"
                    class="size-8 rounded-full"
                  />
                </picture>
              </td>
              <td class="px-6 py-2">{{ item.fullName }}</td>
              <td class="px-6 py-2">{{ item.email }}</td>
              <td class="px-6 py-2">{{ item.lastAccess }}</td>
              <td class="px-6 py-2 flex items-center space-x-2">
                <router-link :to="`/users/${item.id}`">
                  <BaseButtonOutlined
                    label="View Details"
                    size="sm"
                  />
                </router-link>
                <router-link :to="`/users/${item.id}?edit`">
                  <button
                    class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  >
                    <i class="text-base material-icons-round">edit</i>
                  </button>
                </router-link>

                <button
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-red-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  @click="confirmDelete(item)"
                >
                  <i class="text-base material-icons-round">delete</i>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div
        class="tab-content w-full bg-white shadow-md rounded-b-lg overflow-hidden"
        :class="activeTab === 'clients' ? '' : 'hidden'"
      >
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Avatar
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Email
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Organisation
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Last Access
              </th>
              <th
                scope="col"
                class="px-6 py-3"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in clients"
              :key="item.id"
              class="odd:bg-white even:bg-gray-50 border-b"
            >
              <td
                scope="row"
                class="px-6 py-2"
              >
                <picture>
                  <img
                    :src="item.avatar"
                    :alt="item.fullName"
                    class="size-8 rounded-full"
                  />
                </picture>
              </td>
              <td class="px-6 py-2">{{ item.fullName }}</td>
              <td class="px-6 py-2">{{ item.email }}</td>
              <td class="px-6 py-2">{{ item.organisation }}</td>
              <td class="px-6 py-2">{{ item.lastAccess }}</td>
              <td class="px-6 py-2 flex items-center space-x-2">
                <router-link :to="`/users/${item.id}`">
                  <BaseButtonOutlined
                    label="View Details"
                    size="sm"
                  />
                </router-link>
                <router-link :to="`/users/${item.id}?edit`">
                  <button
                    class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-blue-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  >
                    <i class="text-base material-icons-round">edit</i>
                  </button>
                </router-link>

                <button
                  class="flex items-center justify-center size-5 rounded-full bg-gray-200 hover:bg-red-500 hover:text-white text-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50"
                  @click="confirmDelete(item)"
                >
                  <i class="text-base material-icons-round">delete</i>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<style lang="scss">
.main {
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin-left: 80px;
  background-color: #f9f9f9;
  padding: 15px;
}
</style>
