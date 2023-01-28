<script lang="ts">
export default {
  data(){
    return {
      exportData: "This is simple start page"
    }
  }
}
</script>

<template>
  <main class="container">
    <p>{{ exportData }}</p>
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
