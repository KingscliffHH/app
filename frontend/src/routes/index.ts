import { authGuard } from "@auth0/auth0-vue";
import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw
} from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    component: () => import("@/pages/AppLogin.vue")
  },
  {
    path: "/logout",
    component: () => import("@/pages/AppLogout.vue")
  },
  {
    path: "/",
    component: () => import("@/components/layouts/BaseLayout.vue"),
    beforeEnter: authGuard,
    redirect: "/projects",
    children: [
      {
        path: "/projects",
        children: [
          {
            path: "",
            component: () => import("@/pages/ProjectList.vue")
          },
          {
            path: "new",
            name: "projects.new",
            component: () => import("@/pages/ProjectNew.vue")
          },
          {
            path: ":id",
            component: () => import("@/pages/ProjectDashboard.vue"),
            props: true
          },
          {
            path: ":id/edit",
            name: "projects.edit",
            // beforeEnter: myguard(["admin"]),
            component: () => import("@/pages/ProjectEdit.vue"),
            props: true
          },
          {
            path: ":id/metrics",
            name: "projects.metrics",
            // beforeEnter: myguard(["admin"]),
            component: () => import("@/pages/ProjectMetrics.vue"),
            props: true
          }
        ]
      },
      {
        path: "/benchmarks",
        name: "benchmarks",
        // beforeEnter: myguard(["admin"]),
        children: [
          {
            path: "",
            name: "benchmark.list",
            component: () => import("@/pages/BenchmarkList.vue")
          },
          {
            path: "new",
            name: "benchmark.new",
            component: () => import("@/pages/BenchmarkNew.vue")
          },
          {
            path: ":id",
            name: "benchmark.edit",
            component: () => import("@/pages/BenchmarkEdit.vue"),
            props: true
          }
        ]
      },
      {
        path: "/users",
        name: "users",
        children: [
          {
            path: "",
            name: "user.list",
            component: () => import("@/pages/UserList.vue")
          },
          {
            path: "new",
            name: "user.new",
            component: () => import("@/pages/UserNew.vue")
          },
          {
            path: ":id",
            name: "user.edit",
            component: () => import("@/pages/UserEdit.vue"),
            props: true
          }
        ]
      },
      // Not Found
      {
        path: "/:pathMatch(.*401)*",
        component: () => import("@/pages/ErrorUnauthorized.vue")
      },
      {
        path: "/:pathMatch(.*)*",
        component: () => import("@/pages/NotFound.vue")
      }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
