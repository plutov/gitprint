<script setup lang="ts">
import { onMounted, onUnmounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getJWT } from "../lib/api";

const route = useRoute();
const router = useRouter();

let timeout: ReturnType<typeof setTimeout>;

onMounted(async () => {
  timeout = setTimeout(() => {
    router.push("/");
  }, 5000);

  const code = route.query.code as string;
  const state = route.query.state as string;

  if (code && state) {
    const res = await getJWT(code, state);
    if (!res.error && res.data.jwt_token) {
      localStorage.setItem("auth_jwt", res.data.jwt_token);
      clearTimeout(timeout);
      router.push("/");
    }
  }
});

onUnmounted(() => {
  clearTimeout(timeout);
});
</script>

<template>
  <div>
    <div class="pt-20">
      <div class="mx-auto max-w-xl px-4">
        <div class="flex flex-col rounded-md border p-8 text-center">
          <p class="py-2">Authenticating...</p>
        </div>
      </div>
      <div class="h-px w-full"></div>
    </div>
  </div>
</template>
