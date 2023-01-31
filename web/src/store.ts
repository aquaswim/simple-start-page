import { ref } from "vue";
import type { ISetting } from "./data";
import { fetchApi } from "./utils";

export const SettingStore = {
  setting: ref<ISetting>({
    siteTitle: "Simple Start Page",
    siteAbout: "This is simple start page",
    hideSetting: false,
  }),
  fetchSetting: async function () {
    const result = await fetchApi<ISetting>("GET", "/api/setting", null);
    if (result && result.data) {
      this.setting.value.hideSetting = result.data.hideSetting;
      this.setting.value.siteAbout = result.data.siteAbout;
      this.setting.value.siteTitle = result.data.siteTitle;
    }
  },
};

export const AuthStore = {
  authToken: ref(""),
};

export const AlertModalStore = {
  state: ref({
    show: false,
    title: "",
    message: "",
  }),
  showAlertModal: function (message: string, title: string) {
    this.state.value.title = title;
    this.state.value.message = message;
    this.state.value.show = true;
  },
};
