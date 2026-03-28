<template>
  <div>
    <div class="pb-10 pt-10">
      <div class="mx-auto max-w-2xl px-4 flex flex-col gap-4">
        <div class="card">
          <h1 class="text-lg font-semibold text-teal-400">&gt; Which files are printed?</h1>
          <p class="text-gray-300">
            To keep the size of the PDF small, only the whitelisted extensions are include. That
            includes programming languages, markup languages, and some other common file types.
            Files bigger than 100KB are not included in the PDF.
          </p>
        </div>
        <div class="card">
          <h1 class="text-lg font-semibold text-teal-400">&gt; Limitations</h1>
          <p class="text-gray-300">
            Because it is a free service, there are some limitations: you can generate only one PDF
            per hour.
          </p>
        </div>
        <div class="card">
          <h1 class="text-lg font-semibold text-teal-400">&gt; Known Issues</h1>
          <p class="text-gray-300">1. It may not work well on huge repositories.</p>
          <p class="text-gray-300">2. No syntax highlighting, but supports Markdown.</p>
        </div>
        <div class="card">
          <h1 class="text-lg font-semibold text-teal-400">&gt; CHANGELOG</h1>
          <p class="text-gray-300">2026-03-28: Parse full GitHub URL.</p>
          <p class="text-gray-300">2025-12-10: Upgrade JS packages.</p>
          <p class="text-gray-300">2024-10-11: Parse Markdown files.</p>
          <p class="text-gray-300">2024-10-11: Add exclude patterns.</p>
        </div>
      </div>
    </div>
  </div>
</template>
