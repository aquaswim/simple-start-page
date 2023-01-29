<script lang="ts">
import type { ILink } from "@/data"
import Greeting from "@/components/Greeting.vue"

const links: ILink[] = [
    {url: "http://google.com", icon: "cloud", name: "Adadeh"},
]
for (let index = 0; index < 12; index++) {
    links.push({name: `Test ${index + 1}`})
}

export default {
    components: {
        Greeting
    },
    data() {
        return {
            links,
        }
    }
}
</script>
<template>
    <main class="container">
        <Greeting/>
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