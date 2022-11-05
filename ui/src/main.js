import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from "./router";
import store from "./store";
import request from "@/request/index";
import './assets/global.css';
import { MarkdownEditor } from 'markdown-it-editor'
import 'markdown-it-editor/lib/index.css'

Vue.use(ElementUI)
Vue.use(MarkdownEditor)

Vue.config.productionTip = false

Vue.prototype.request = request

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
