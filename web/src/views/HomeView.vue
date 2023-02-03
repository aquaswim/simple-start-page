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
        <div>
          <h3>
            {{ link.name }}
          </h3>
          <small>
            {{ link.url }}
          </small>
        </div>
      </a>
    </div>
  </section>
</main>
</template>

<style scoped>
.app-list {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
}
.app {
  display: flex;
  flex-direction: row;
  margin-bottom: 1rem;
}
a.app:is(:hover, :active, :focus) {
  text-decoration: none;
}
.app div h3{
  margin-bottom: 0;
}
i.material-icons {
  vertical-align: bottom;
  font-size: 3.4rem;
}
</style>
