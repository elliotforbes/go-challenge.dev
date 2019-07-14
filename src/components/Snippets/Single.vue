<template>
  <div class="flex flex-justify">
    <div class="flex-1 text-gray-700 px-6 py-6 m-2">
      <h1 class="text-lg">Hello World - Go Snippet</h1>
      <pre><code class="go">package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

</code></pre>
    </div>
    <div class="flex-1 text-gray-700 px-6 py-6 m-2">
      <div>
        <h1 class="text-lg">Problem Statement</h1>
        <div class="problem mb-8 p-8">
          <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Unde veritatis, modi aliquam aspernatur illo, exercitationem odit enim iusto voluptas iste est corrupti repudiandae in! Qui illum dolore fugiat accusamus ipsa.</p>
          <pre><code class="go">package main

import "fmt"</code></pre>

          <p>Print something using the Println method:</p>

          <pre><code class="go">fmt.Println("Hello Go Challengers")</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Single"
};
</script>

<style>
.hljs {
  padding: 20px 30px;
  border-radius: 10px;
}

.problem {
  background-color: #f3f8fb;
  border-radius: 8px;
}
</style>

