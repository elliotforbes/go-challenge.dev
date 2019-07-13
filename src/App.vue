<template>
  <div id="app">
    <Navbar />
    <div class="flex">
        <div class="flex-1 text-gray-700 px-6 py-6 m-2">
          <Editor />
        </div>
        <div class="flex-1 text-gray-700 px-6 py-6 m-2">
          <Problem />
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Editor from "./components/Editor.vue";
import Login from "./components/Login.vue";
import Problem from "./components/Problem.vue";

import "./assets/style.css"

export default {
  name: "app",
  components: {
    Navbar,
    Editor,
    Login,
    Problem
  }
};
</script>

<style>
</style>
