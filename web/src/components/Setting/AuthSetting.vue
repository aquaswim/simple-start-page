<script setup lang="ts">
import { AlertModalStore, AuthStore } from "@/store";
import { useFetch } from "@/utils";
import { ref } from "vue";
const updAuthForm = ref({ username: "", password: "" });
const {
  isLoading,
  error,
  execute: executeUpdate,
} = useFetch("PUT", "/api/auth/login");
const authUpdate = async () => {
  await executeUpdate({ ...updAuthForm.value }, AuthStore.authToken.value);
  if (error.value && error.value.message) {
    AlertModalStore.showAlertModal(error.value.message, "Error update auth");
  }
};
</script>
<template>
  <article>
    <form @submit.prevent="authUpdate">
      <label for="input-auth-username">
        Username
        <input
          type="text"
          id="input-auth-username"
          name="input-auth-username"
          placeholder="new username"
          required
          v-model="updAuthForm.username"
        />
      </label>

      <label for="lastname">
        Password
        <input
          type="password"
          id="lastname"
          name="lastname"
          placeholder="new password"
          required
          v-model="updAuthForm.password"
        />
      </label>
      <button type="submit" :aria-busy="isLoading">Save</button>
    </form>
  </article>
</template>
