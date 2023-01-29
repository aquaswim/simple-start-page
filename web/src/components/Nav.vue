<script setup lang="ts">
import { RouterLink } from "vue-router";
import { AuthStore, SettingStore } from "@/store";
</script>

<template>
    <nav class="container-fluid">
        <ul>
            <li><strong>{{ SettingStore.setting.value.siteTitle }}</strong></li>
        </ul>
        <ul>
            <li><RouterLink to="/">Home</RouterLink></li>
            <!-- <li><RouterLink to="/login">Login</RouterLink></li> -->
            <li><RouterLink to="/setting" v-if="!SettingStore.setting.value.hideSetting || AuthStore.authToken.value !== ''">Setting</RouterLink></li>
            <li><RouterLink to="/about">About</RouterLink></li>
        </ul>
    </nav>
</template>