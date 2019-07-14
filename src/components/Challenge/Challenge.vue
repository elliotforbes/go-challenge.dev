<template>
  <div>
    <div class="flex-1 text-gray-700 px-6 py-6 m-2">
      <Editor />
    </div>
    <div class="flex-1 text-gray-700 px-6 py-6 m-2">
      <Problem />
    </div>
  </div>
</template>

<script>
import Problem from './Problem.vue'
import Editor from './Editor.vue'

export default {
    name: 'Challenge',
    components: {
        Editor,
        Problem
    }
}
</script>
