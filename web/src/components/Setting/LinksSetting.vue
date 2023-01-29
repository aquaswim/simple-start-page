<script setup lang="ts">
import type { ILink } from '@/data';
import { ref } from 'vue';

const linksListForm = ref<(ILink & {id: string})[]>([])
const addLinkForm = ref<ILink>({name: "", url: "", icon: ""})
const addLinkToList = () => {
    linksListForm.value.push({
        ...addLinkForm.value,
        id: Date.now().toString(16)
    })
    addLinkForm.value = {name: ""}
}
const deleteLinkFromList = (id: string) => {
    linksListForm.value = linksListForm.value.filter( l => l.id !== id)
}

</script>
<template>
    <div>
        <table role="grid">
            <thead>
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col" data-tooltip="icon name in material icon list">Icon</th>
                    <th scope="col">Url</th>
                    <th scope="col">Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="link in linksListForm" :key="link.id">
                    <td>{{ link.name }}</td>
                    <td>
                        <i class="material-icons">{{ link.icon }}</i>
                    </td>
                    <td>
                        <a :href="link.url" target="_blank">{{ link.url }}</a>
                    </td>
                    <th scope="row">
                        <button class="secondary" @click="deleteLinkFromList(link.id)">
                            <i class="material-icons">delete</i>
                        </button>
                    </th>
                </tr>
                <tr>
                    <td>
                        <input type="text" placeholder="My Url" v-model="addLinkForm.name"/>
                    </td>
                    <td>
                        <input type="text" placeholder="question_mark" v-model="addLinkForm.icon"/>
                    </td>
                    <td>
                        <input type="text" placeholder="https://example.com" v-model="addLinkForm.url"/>
                    </td>
                    <th scope="row">
                        <button class="contrast" @click="addLinkToList">
                            <i class="material-icons">done</i>
                        </button>
                    </th>
                </tr>
            </tbody>
        </table>
        <button>Save</button>
    </div>
</template>