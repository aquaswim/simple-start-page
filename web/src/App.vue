<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import Nav from "./components/Nav.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <Nav />
  <RouterView />
  <Footer />
</template>

<style scoped>
</style>
