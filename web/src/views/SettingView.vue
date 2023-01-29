<script setup lang="ts">
import type { ISetting } from '@/data';
import { AuthStore, SettingStore } from '@/store';
import { useFetch } from '@/utils';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

// todo: use component for this general setting
const sgFormData = ref<ISetting>({
    siteTitle: "",
    siteAbout: "",
    hideSetting: false
})
const {isLoading: sgUpdIsLoading, error: sgUpdError, execute: sgUpdExecute} = useFetch("PUT", "/api/setting")
async function sgUpdate() {
    await sgUpdExecute({
        siteTitle: sgFormData.value.siteTitle,
        siteAbout: sgFormData.value.siteAbout,
        hideSetting: sgFormData.value.hideSetting
    } as ISetting, AuthStore.authToken.value)
    if (sgUpdError && sgUpdError.value) {
        alert(sgUpdError.value.message)
    }
    await SettingStore.fetchSetting()
}
// end of general setting

// note: idk why i cannot use guard in here
onMounted( async () => {
    if (AuthStore.authToken.value === "") {
        router.replace("/login")
    }
    await SettingStore.fetchSetting()
    sgFormData.value.hideSetting = SettingStore.setting.value.hideSetting
    sgFormData.value.siteAbout = SettingStore.setting.value.siteAbout
    sgFormData.value.siteTitle = SettingStore.setting.value.siteTitle
})

</script>
<template>
    <main class="container">
        <form @submit.prevent="sgUpdate">
            <label for="input-site-title">Text</label>
            <input type="text" id="input-site-title" name="site_title" placeholder="My Site" v-model="sgFormData.siteTitle">
            
            <label for="input-about">About page content</label>
            <textarea name="about" id="input-about" rows="10" placeholder="This is my simple site" v-model="sgFormData.siteAbout"></textarea>
            
            <fieldset>
                <label for="input-hide-setting">
                    <input type="checkbox" id="input-hide-setting" name="hide_setting" role="switch" v-model="sgFormData.hideSetting">
                    Hide setting page
                </label>
            </fieldset>
            
            <input type="submit" value="Save" :aria-busy="sgUpdIsLoading">
        </form>
    </main>
</template>