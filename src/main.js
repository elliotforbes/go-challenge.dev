import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router'

import Login from "./components/Login.vue"
import HomePage from "./components/HomePage/HomePage.vue"
import Challenge from "./components/Challenge/Challenge.vue"

Vue.use(VueRouter)
Vue.config.productionTip = false

const routes = [
  { path: "/", component: HomePage },
  { path: "/login", component: Login },
  { path: "/challenge/:id", component: Challenge }
];

const router = new VueRouter({
  routes
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
