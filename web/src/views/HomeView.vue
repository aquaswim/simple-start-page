<script lang="ts">
import type { ILink } from "@/data";
import Greeting from "@/components/TimeGreeting.vue";
import { useFetch } from "@/utils";

export default {
  components: {
    Greeting,
  },
  setup() {
    const { isLoading, data, execute } = useFetch<ILink[]>("GET", "/api/links");
    return {
      isLoading,
      links: data,
      fetchLinks: execute,
    };
  },
  mounted() {
    this.fetchLinks();
  },
};
</script>
<template>
  <main class="container">
    <Greeting />
    <section id="applications" :aria-busy="isLoading">
      <div class="grid app-list">
        <a
          v-for="link in links"
          :key="link.url"
          :href="link.url || '#'"
          class="app"
          target="_blank"
        >
          <i class="material-icons">{{ link.icon || "question_mark" }}</i>
          <span>{{ link.name }}</span>
        </a>
      </div>
    </section>
  </main>
</template>

<style scoped>
.app-list {
  display: flex;
  flex-wrap: wrap;
}
.app {
  text-decoration: none;
  color: unset;
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 0.5rem;
  margin: 1rem;
}
.app i {
  font-size: 2rem;
}
</style>
