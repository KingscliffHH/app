<script lang="ts" setup>
import { ref } from "vue";
import { useRoute } from "vue-router";

import { user, userRole } from "@/store/auth";

const route = useRoute();
const version = APP_VERSION;

// TODO: improve onClick
const menuItems = ref([
  {
    icon: "home",
    name: "Home",
    path: "/projects",
    requiresPermission: [],
    isActive: () => route.path.indexOf("/projects") === 0,
    onClick: () => {}
  },
  {
    icon: "bar_chart",
    name: "Benchmarks",
    path: "/benchmarks",
    requiresPermission: ["admin"],
    isActive: () => route.path.indexOf("/benchmarks") === 0,
    onClick: () => {}
  },
  {
    icon: "person",
    name: "Users",
    path: "/users",
    requiresPermission: ["admin"],
    isActive: () => route.path.indexOf("/users") === 0,
    onClick: () => {}
  }
]);
</script>

<template>
  <div class="app-menu">
    <div class="app-menu__top-items">
      <router-link to="/projects">
        <img
          src="@/assets/logo.svg"
          alt="logo"
          class="app-menu__top-items--logo"
        />
      </router-link>

      <picture>
        <img
          :src="user?.avatar"
          :alt="user?.fullName"
          class="app-menu__top-items--user-photo"
        />

        <v-tooltip
          activator="parent"
          location="start"
        >
          {{ user?.fullName }}
        </v-tooltip>
      </picture>

      <template
        v-for="item in menuItems"
        :key="item.icon"
      >
        <router-link
          v-if="
            !item.requiresPermission.length ||
            item.requiresPermission.indexOf(userRole) !== -1
          "
          :to="item.path"
        >
          <v-btn variant="flat">
            <i
              :active="item.isActive()"
              class="material-icons-round"
            >
              {{ item.icon }}
            </i>
            <v-tooltip
              activator="parent"
              location="start"
            >
              {{ item.name }}
            </v-tooltip>
          </v-btn>
        </router-link>
      </template>
    </div>

    <div class="flex flex-col align-center">
      <router-link to="/logout">
        <i class="material-icons-round"> logout </i>
      </router-link>
      <span
        v-if="userRole === 'admin'"
        class="font-bold text-gray-300"
        >v{{ version }}</span
      >
    </div>
  </div>
</template>

<style lang="scss">
.app-menu {
  position: fixed;
  height: 100vh;
  width: 80px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding-block: 30px;
  background-color: white;
  box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;

  i {
    &[active="true"] {
      color: #1a3c5b;
    }

    cursor: pointer;
    font-size: 34px;
    color: grey;
  }

  &__top-items {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;

    &--logo {
      cursor: pointer;
      width: 40px;
      height: 55px;
      margin-bottom: 20px;
    }

    &--user-photo {
      width: 50px;
      height: 50px;
      border-radius: 50%;
      margin-bottom: 20px;
    }
  }
}
</style>
