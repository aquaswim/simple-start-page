<script setup lang="ts">
import { SettingStore } from "@/store";
</script>

<template>
  <main class="container">
    <p>{{ SettingStore.setting.value.siteAbout }}</p>
  </main>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
