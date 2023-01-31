<script setup lang="ts">
import { AlertModalStore } from "@/store";
</script>

<template>
  <dialog :open="AlertModalStore.state.value.show">
    <article>
      <h3>{{ AlertModalStore.state.value.title || "Error" }}</h3>
      <p>
        {{ AlertModalStore.state.value.message || "Unknown Error" }}
      </p>
      <footer>
        <button @click.stop="AlertModalStore.state.value.show = false">
          OK
        </button>
      </footer>
    </article>
  </dialog>
</template>
