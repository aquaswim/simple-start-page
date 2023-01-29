<script setup lang="ts">
import { RouterView } from "vue-router";
import Nav from "./components/Nav.vue";
import Footer from "./components/Footer.vue";
import { SettingStore } from "./store";
import { onMounted } from "vue";
// todo refetch setting mechanism
onMounted(() => {
  SettingStore.fetchSetting()
})
</script>

<template>
  <Nav />
  <RouterView />
  <Footer />
</template>

<style scoped>
</style>
