<script setup lang="ts">
import { RouterView } from "vue-router";
import Nav from "./components/MainNav.vue";
import Footer from "./components/MainFooter.vue";
import AlertModal from "@/components/AlertModal.vue";
import { SettingStore } from "./store";
import { onMounted } from "vue";
// todo refetch setting mechanism
onMounted(() => {
  SettingStore.fetchSetting();
});
</script>

<template>
  <AlertModal />
  <Nav />
  <RouterView />
  <Footer />
</template>

<style scoped></style>
