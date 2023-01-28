<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
    <nav class="container-fluid">
        <ul>
            <li><strong>Simple Start Page</strong></li>
        </ul>
        <ul>
            <li><RouterLink to="/">Home</RouterLink></li>
            <!-- <li><RouterLink to="/login">Login</RouterLink></li> -->
            <li><RouterLink to="/setting">Setting</RouterLink></li>
            <li><RouterLink to="/about">About</RouterLink></li>
        </ul>
    </nav>
</template>