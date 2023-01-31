<script setup lang="ts">
import { useFetch } from "@/utils";
import { reactive } from "vue";
import { AlertModalStore, AuthStore } from "@/store";
import type { ILoginResp } from "@/data";
import { useRouter } from "vue-router";

const router = useRouter();
const loginForm = reactive({ username: "", password: "" });
const {
  isLoading: loginIsLoading,
  data: loginResponse,
  error: loginError,
  execute: executeLogin,
} = useFetch<ILoginResp>("POST", "/api/auth/login");
async function login() {
  await executeLogin({
    username: loginForm.username,
    password: loginForm.password,
  });
  if (loginError.value != null) {
    AlertModalStore.showAlertModal(loginError.value.message, "Login Error");
  }

  if (loginResponse.value != null && loginResponse.value.token) {
    // process login
    AuthStore.authToken.value = loginResponse.value.token;
    router.push("/setting");
  } else {
    // alert("Unexpected error");
  }
}
</script>

<template>
  <main class="container">
    <article class="grid">
      <div>
        <hgroup>
          <h1>Sign in</h1>
          <h2>Please sign in to access setting page</h2>
        </hgroup>
        <form @submit.prevent="login">
          <input
            type="text"
            name="username"
            placeholder="Login"
            aria-label="Login"
            autocomplete="nickname"
            required
            v-model="loginForm.username"
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            aria-label="Password"
            autocomplete="current-password"
            required
            v-model="loginForm.password"
          />
          <button type="submit" :aria-busy="loginIsLoading">Login</button>
        </form>
      </div>
    </article>
  </main>
</template>
