import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import CallbackView from "../views/CallbackView.vue";
import DocsView from "../views/DocsView.vue";
import ContactView from "../views/ContactView.vue";

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: HomeView },
    { path: "/callback", component: CallbackView },
    { path: "/docs", component: DocsView },
    { path: "/contact", component: ContactView },
  ],
});
