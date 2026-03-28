<script setup lang="ts">
import { ref, onMounted } from "vue";
import { parseJwt } from "../lib/jwt";
import { download, generate } from "../lib/api";
import type { User } from "../lib/types";

const apiAddr = import.meta.env.VITE_API_ADDR as string;

enum ProgressStatus {
  None = "none",
  Downloading = "downloading",
  Generating = "generating",
  Done = "done",
  Error = "error",
}

const user = ref<User | undefined>();
const token = ref<string | undefined>();
const repo = ref("");
const refVal = ref("");
const advanced = ref(false);
const exclude = ref("");
const exportId = ref("");
const progress = ref<ProgressStatus>(ProgressStatus.None);
const log = ref<string[]>([]);

onMounted(() => {
  const jwtToken = localStorage.getItem("auth_jwt");
  if (!jwtToken) return;

  const claims = parseJwt(jwtToken);
  if (!claims) return;

  token.value = jwtToken;
  user.value = claims.user;
});

function logout() {
  localStorage.removeItem("auth_jwt");
  user.value = undefined;
}

function parseRepo(input: string): string {
  try {
    const url = new URL(input);
    if (url.hostname === "github.com") {
      const parts = url.pathname.replace(/^\//, "").split("/");
      if (parts.length >= 2 && parts[0] && parts[1]) {
        return `${parts[0]}/${parts[1]}`;
      }
    }
  } catch {}
  return input;
}

function onRepoInput(e: Event) {
  repo.value = parseRepo((e.target as HTMLInputElement).value);
}

function isValidRepo(r: string) {
  const parts = r.split("/");
  return parts.length === 2 && parts[0] && parts[1];
}

async function start() {
  progress.value = ProgressStatus.Downloading;
  log.value = ["Downloading repository files..."];

  const res = await download(token.value as string, repo.value, refVal.value, exclude.value);
  if (res.error) {
    progress.value = ProgressStatus.Error;
    log.value = [...log.value, "ERROR: " + res.error];
  } else {
    progress.value = ProgressStatus.Generating;
    log.value = [...log.value, "DONE", "Generating PDF..."];

    exportId.value = res.data.exportID;
    const generateRes = await generate(
      token.value as string,
      repo.value,
      refVal.value,
      res.data.exportID,
    );
    if (generateRes.error) {
      progress.value = ProgressStatus.Error;
      log.value = [...log.value, "ERROR: " + generateRes.error];
    } else {
      progress.value = ProgressStatus.Done;
      log.value = [...log.value, "DONE"];
    }
  }
}
</script>

<template>
  <div class="mx-auto max-w-2xl px-4">
    <div
      class="flex flex-col gap-2 rounded-md border border-gray-800 p-4 justify-center items-center"
      style="background: rgba(0, 9, 29, 0.5)"
    >
      <p v-if="user">
        Signed in as
        <span class="font-semibold text-teal-400">{{ user.username }}</span>
        (<a href="#" @click.prevent="logout">Log out</a>)
      </p>
      <a v-else :href="`${apiAddr}/github/auth/url`">Sign in with GitHub</a>
    </div>
    <div
      v-if="user"
      class="flex flex-col gap-2 rounded-md border border-gray-800 p-4 mt-2"
      style="background: rgba(0, 9, 29, 0.5)"
    >
      <div>
        <label for="repository" class="block text-sm text-white font-semibold">Repository</label>
        <label for="repository" class="block text-xs text-gray-500 mt-0.5">
          Must be a public GitHub repository.
        </label>
        <div class="mt-2">
          <input
            type="text"
            id="repository"
            class="input"
            placeholder="owner/repo"
            :value="repo"
            @input="onRepoInput"
          />
        </div>
      </div>
      <div class="mt-2">
        <label for="ref" class="block text-sm text-white font-semibold">[Optional] Ref</label>
        <label for="ref" class="block text-xs text-gray-500 mt-0.5">
          Could be a branch, tag, or commit SHA.
        </label>
        <div class="mt-2">
          <input type="text" id="ref" class="input" placeholder="main" v-model="refVal" />
        </div>
      </div>
      <div class="mt-2">
        <button type="button" class="btn inline" @click="advanced = !advanced">
          Advanced Options
        </button>
      </div>
      <div v-if="advanced" class="mt-2">
        <label for="exclude" class="block text-sm text-white font-semibold"
          >[Optional] Exclude</label
        >
        <label for="exclude" class="block text-xs text-gray-500 mt-0.5">
          Comma separated list of patterns to exclude.
        </label>
        <div class="mt-2">
          <input
            type="text"
            id="exclude"
            class="input"
            placeholder="tests/*,docs/*"
            v-model="exclude"
          />
        </div>
      </div>
      <div class="mt-2">
        <button type="submit" class="btn" :disabled="!isValidRepo(repo)" @click="start">
          Generate PDF
        </button>
      </div>
    </div>
    <div
      v-if="progress !== 'none'"
      class="rounded-md border border-gray-800 p-4 mt-2"
      style="background: rgba(0, 9, 29, 0.7)"
    >
      <div class="text-xs text-gray-500 mb-2 border-b border-gray-800 pb-2 flex items-center gap-2">
        <span
          class="w-2 h-2 rounded-full"
          :class="
            progress === 'error'
              ? 'bg-red-500'
              : progress === 'done'
                ? 'bg-teal-500'
                : 'bg-yellow-500 animate-pulse'
          "
        ></span>
        LOG OUTPUT
      </div>
      <pre class="text-sm text-teal-300 whitespace-pre-wrap">{{ log.join("\n") }}</pre>
      <p v-if="progress === 'done'" class="mt-3 pt-3 border-t border-gray-800">
        <a :href="`${apiAddr}/files?export_id=${exportId}&ext=pdf`">Download PDF</a>
      </p>
    </div>
  </div>
</template>
