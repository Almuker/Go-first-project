<template>
  <nav class="bg-gray-700 py-4 flex">
    <RouterLink to="/">
      <button 
        class="bg-slate-200 px-4 py-2 ml-2 rounded-sm font-serif text-lg transition"
        :class="{'active': $route.name === 'home'}"
        >Home
      </button>
    </RouterLink>
    <RouterLink to="/about">
      <button 
        class="bg-slate-200 px-4 py-2 ml-2 rounded-sm font-serif text-lg transition"
        :class="{'active': $route.name === 'about'}"
        >About
      </button>
    </RouterLink>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>