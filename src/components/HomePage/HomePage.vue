<template>
  <div>
    <div class="w-full p-8">
      <h2 class="text-lg"></h2>
    </div>

    <div class="challenges min-h-screen">
      <div class="flex flex-wrap">
        <ChallengeCard />
        <ChallengeCard />
        <ChallengeCard />
        <ChallengeCard />
        <ChallengeCard />
      </div>
    </div>
  </div>
</template>

<script>
import ChallengeCard from "./ChallengeCard.vue";

export default {
  name: "HomePage",
  components: {
    ChallengeCard
  }
};
</script>

<style>
.challenges {
  background-color: #f2f3f7;
  width: 100%;
}
</style>

