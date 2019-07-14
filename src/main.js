import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router'

import Login from "./components/Login.vue"
import HomePage from "./components/HomePage/HomePage.vue"
import Challenge from "./components/Challenge/Challenge.vue"
import Snippets from "./components/Snippets/Snippets.vue"
import Single from "./components/Snippets/Single.vue"

Vue.use(VueRouter)
Vue.config.productionTip = false

const routes = [
  { path: "/", component: HomePage },
  { path: "/login", component: Login },
  { path: "/challenge/:id", component: Challenge },
  { path: "/snippets", component: Snippets },
  { path: "/snippet/:id", component: Single },
];

const router = new VueRouter({
  routes
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
