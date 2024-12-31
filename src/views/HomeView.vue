<script setup lang="ts">

import NavigationBar from "@/components/NavigationBar.vue";
import OpeningAnimation from "@/components/browse/OpeningAnimation.vue";
import {playedOA} from "@/globs";


let displayOA = !playedOA.value;
playedOA.value = true;
</script>

<template>
  <OpeningAnimation v-if="displayOA"/>
  <NavigationBar/>
  <div class="mx-auto mt-32 max-w-screen-2xl px-5 flex flex-row text-white">
    <div id="house-of-ui-img" class="lg:block hidden"></div>
    <div class="max-w-[48rem] px-10">
      <h5 class="font-thin tracking-widest mt-12 text-sm">HI I'M UI-TOWN</h5>
      <h1 class="font-thin text-7xl text-left text-[#D0C3F1] mt-16">A town of UI which brings your design to life.</h1>
      <h4 class="tracking-widest leading-9 text-lg max-w-lg ps-1.5 font-[Cooljazz] mt-16">"At UI-town, we offer a vast array of CSS and JS animations to elevate your web projects. Our dynamic animations and interactive effects make your user interfaces not just functional but truly captivating."</h4>
      <RouterLink :to="{name: 'browse'}" class="ms-0.5 w-40 mt-8 py-4 text-center rounded-full inline-block bg-[#D0C3F1] text-[rgb(85,26,139)] hover:text-white">Enter</RouterLink>
    </div>
  </div>
</template>

<style scoped>
#house-of-ui-img {
  @apply h-[45rem] rounded-2xl flex-grow min-w-[490px];
  background: url("@/assets/house_of_ui.png") no-repeat;
  background-size: auto 100%;
  background-position-x: center;
}
</style>