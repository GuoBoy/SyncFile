import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import zhCn from "element-plus/es/locale/lang/zh-cn";
import "element-plus/dist/index.css";

import "bootstrap-icons/font/bootstrap-icons.css";

const app = createApp(App);

app
  .use(createPinia())
  .use(ElementPlus, {
    locale: zhCn,
  })
  .use(router)
  .mount("#app");
