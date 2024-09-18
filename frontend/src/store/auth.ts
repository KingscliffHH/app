import { JwtPayload, jwtDecode } from "jwt-decode";
import { computed, ref, watch } from "vue";

import { User } from "@/types/user";

import router from "@/routes";

export const token = ref<string>("");
export const user = ref<User | null>(null);

type UserRole = "admin" | "member" | "client" | "unknown";

export const decodedToken = computed(() => {
  if (!token.value) return null;
  return jwtDecode<JwtPayload & { "https://ci.com.au/roles": UserRole[] }>(
    token.value
  );
});

export const userRole = computed<UserRole>(() => {
  if (!decodedToken.value) return "unknown";

  const roles = decodedToken.value["https://ci.com.au/roles"];
  if (roles.includes("admin")) return "admin";
  if (roles.includes("member")) return "member";
  if (roles.includes("client")) return "client";

  return "unknown";
});

watch(userRole, (role) => {
  console.info("w user role", role);
  if (role === "unknown") return;

  if (role !== "admin") {
    console.info("not admin", role, router);
    router.removeRoute("benchmarks");
    router.removeRoute("users");
    console.info("removed ", router.hasRoute("benchmarks"));

    if (role !== "member") {
      router.removeRoute("projects.edit");
      router.removeRoute("projects.metrics");
    }

    router.push(window.location.pathname);
  }
});
