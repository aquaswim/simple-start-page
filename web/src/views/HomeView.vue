<script lang="ts">
import type { ILink } from "@/data"

const links: ILink[] = [
    {url: "http://google.com", icon: "cloud", name: "Adadeh"},
]
for (let index = 0; index < 12; index++) {
    links.push({name: `Test ${index + 1}`})
}

// todo greeting by time
const greeting = "Hello world"

export default {
    data() {
        return {
            links,
            greeting
        }
    }
}
</script>
<template>
    <main class="container">
        <h1><u>{{ greeting }}</u></h1>
        <section id="applications">
            <div class="grid app-list">
                <a v-for="(link, index) in links" :key="link.url" :href="link.url || '#'" class="app">
                    <i class="material-icons">{{ link.icon || 'question_mark' }}</i>
                    <span>{{ link.name }}</span>
                </a>
            </div>
        </section>
    </main>
</template>

<style scoped>
.app-list{
    display: flex;
    flex-wrap: wrap;
}
.app{
    text-decoration: none;
    color: unset;
    display: flex;
    flex-direction: column;
    align-items: center;
    row-gap: .5rem;
    margin: 1rem;
}
.app i {
    font-size: 2rem;
}
</style>