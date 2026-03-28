<template>
  <div class="flex flex-col min-h-screen">
    <header
      class="flex h-16 w-full items-center justify-center border-b border-teal-900/50 sticky top-0 z-10 backdrop-blur-md"
      style="background: rgba(0, 9, 29, 0.85)"
    >
      <div class="flex justify-center p-4 text-sm">
        <span class="mx-2"><RouterLink to="/">home</RouterLink></span>
        <span class="text-gray-600">•</span>
        <span class="mx-2"><RouterLink to="/docs">docs</RouterLink></span>
        <span class="text-gray-600">•</span>
        <span class="mx-2"><RouterLink to="/contact">contact</RouterLink></span>
        <span class="text-gray-600">•</span>
        <span class="mx-2"><a href="https://github.com/plutov/gitprint">github</a></span>
      </div>
    </header>
    <main class="flex flex-col flex-1">
      <RouterView />
    </main>
    <footer class="flex w-full items-center justify-center border-t border-teal-900/30">
      <div class="flex justify-center p-4 text-sm text-gray-500">
        <span class="mx-2">v0.1.5</span>
        <span class="text-gray-700">•</span>
        <span class="mx-2"
          >Made by <a rel="nofollow" href="https://x.com/pliutau">@pliutau</a></span
        >
        <span class="text-gray-700">•</span>
        <span class="mx-2"><RouterLink to="/docs">How it works</RouterLink></span>
      </div>
    </footer>
  </div>
</template>
