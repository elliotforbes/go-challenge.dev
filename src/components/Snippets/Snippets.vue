<template>
  <div>
    <div class="w-full p-8">
      <h2 class="text-lg">Golang Code Snippets</h2>
    </div>

    <div class="challenges min-h-screen">
      <div class="flex flex-wrap justify-center">
        <SnippetCard />
        <SnippetCard />
        <SnippetCard />
        <SnippetCard />
        <SnippetCard />
      </div>
    </div>
  </div>
</template>


<script>
import SnippetCard from './SnippetCard.vue'

export default {
    name: 'Snippets',
    components: {
        SnippetCard
    }
}
</script>

