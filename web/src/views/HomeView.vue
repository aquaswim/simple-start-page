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
    <section class="services">
      <a target="_blank" :href="link.url" class="service-item" v-for="link in links" :key="link.url">
        <div class="service-icon">
          <i class="material-icons">{{ link.icon || "question_mark" }}</i>
        </div>
        <div class="service-heading">{{ link.name }}</div>
      </a>
    </section>
  </main>
</template>

<style scoped>
.services {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  grid-gap: 20px;
}

.service-item {
  background-color: #121e27;
  padding: 20px;
  border-radius: 5px;
  text-align: center;
  transition: background-color 0.3s;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-decoration: none;
  color: unset;
}

.service-item:hover {
  background-color: #293846;
}

.service-icon {
  font-size: 40px;
  margin-bottom: 10px;
}

.service-heading {
  font-size: 18px;
  margin-bottom: 5px;
}
</style>
