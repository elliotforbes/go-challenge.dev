<template>
  <div id="app">
    <Navbar />

    <div class="flex">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";

import "./assets/style.css"

export default {
  name: "app",
  components: {
    Navbar
  }
};
</script>

<style>
</style>
