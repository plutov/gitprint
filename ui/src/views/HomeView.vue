<script setup lang="ts">
import { ref, onMounted } from "vue";
import Form from "../components/Form.vue";
import { getPopularRepos } from "../lib/api";
import type { RepoInfo } from "../lib/types";

const repos = ref<RepoInfo[]>([]);
const loading = ref(true);

onMounted(async () => {
  try {
    const response = await getPopularRepos();
    if (response.status === 200 && response.data && response.data.repos) {
      repos.value = response.data.repos;
    }
  } catch (error) {
    console.error("Failed to fetch recent repos:", error);
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <div>
    <div class="pb-10 pt-10">
      <div class="mx-auto max-w-2xl px-4">
        <div class="card">
          <h1 class="text-lg font-semibold text-teal-400">
            &gt; Print your favorite Git repositories as PDF
          </h1>
          <p class="py-2 text-gray-300">
            Looking for a fun way to explore your favorite GitHub repositories? Tired of staring at
            the screen for hours on end? Or maybe want to print out a hard copy as a keepsake?
          </p>
          <p class="py-2 text-gray-300">
            Simply sign in with your GitHub account and start printing
            <span class="text-white font-semibold">public</span>
            repositories in a beautiful, easy-to-read format.
          </p>
          <p class="py-2 text-gray-300">
            It is currently in beta, so please be patient with us as we work. Feel free to request
            features or report bugs.
          </p>
          <p class="py-2 font-bold">Made by <a href="https://x.com/pliutau">@pliutau</a> with ❤️</p>
        </div>
      </div>
    </div>
    <div class="pb-10">
      <Form />
    </div>
    <div v-if="!loading && repos.length > 0" class="pb-10">
      <div class="mx-auto max-w-2xl px-4">
        <div class="flex flex-col text-center pb-4">
          <h2 class="text-lg font-semibold text-teal-400">Popular repositories</h2>
          <span class="text-sm text-gray-500">Click to download the PDF</span>
        </div>
        <div class="mb-4 grid grid-cols-2 gap-2">
          <a
            v-for="repo in repos"
            :key="repo.name"
            class="cursor-pointer rounded-md border border-gray-800 p-4 transition-all hover:border-teal-700/50 hover:bg-teal-950/20 no-underline"
            style="color: inherit"
            :href="repo.download_url"
          >
            <div class="text-sm font-semibold text-white">{{ repo.name }}</div>
            <div class="text-sm text-gray-400">{{ repo.version }}</div>
            <div class="text-xs text-gray-600 mt-1">{{ repo.size }}</div>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>
