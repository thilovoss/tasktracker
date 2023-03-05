<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import './styles/app.css'; // Here

</script>

<template >
  <div>
  <header>
    <div >
      <div class="p-2">
        <h1 class="mt-0 mb-2 text-5xl font-medium leading-tight text-primary">tasktracker</h1>
        <h3 class="mt-0 mb-2 text-3xl font-medium leading-tight text-primary">
          track your tasks
        </h3>
      </div>

      <!-- <nav>
        <RouterLink to="/">Home</RouterLink>
      </nav> -->
    </div>
  </header>

  <RouterView />

</div>
</template>

<style scoped>

</style>
