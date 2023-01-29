<script setup lang="ts">
import { AuthStore } from '@/store';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import GeneralSetting from '@/components/Setting/GeneralSetting.vue'
import LinksSetting from '@/components/Setting/LinksSetting.vue'
import AuthSetting from '@/components/Setting/AuthSetting.vue'

const router = useRouter()

// note: idk why i cannot use guard in here
onMounted( async () => {
    if (AuthStore.authToken.value === "") {
        router.replace("/login")
    }
})

</script>
<template>
    <main class="container">
        <details>
            <summary>General Setting</summary>
            <GeneralSetting/>
        </details>
        <details>
            <summary>App List Setting</summary>
            <LinksSetting/>
        </details>
        <details>
            <summary>Auth Setting</summary>
            <AuthSetting/>
        </details>
    </main>
</template>