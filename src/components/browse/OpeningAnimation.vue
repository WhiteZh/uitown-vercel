<script setup lang="ts">
import {onMounted, ref} from "vue";

const openingExist = ref(true);
const textExist = ref(true);

onMounted(() => {
  setTimeout(() => textExist.value = false, 900);
  setTimeout(() => openingExist.value = false, 1200);
})
</script>

<template>
  <Transition name="v-container">
    <div class="oa-container" v-if="openingExist">
      <Transition name="v-title" appear>
        <div class="openingElement" style="top: 25vh; font-size: 30vh;" v-if="textExist">
          New design
        </div>
      </Transition>
      <Transition name="v-title" appear>
        <div class="openingElement" style="bottom: 25vh; transition-delay: 0.1s; transition-duration: 0.5s; font-size: 8vh;" v-if="textExist">
          From UI Town
        </div>
      </Transition>
    </div>
  </Transition>
</template>

<style scoped>
.oa-container {
  position: absolute;
  z-index: 1001;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100%;
  background: linear-gradient(135deg, black, #737373);
}

.openingElement {
  position: absolute;
  text-align: center;
  color: white;
  width: 100%;
  font-family: Amsterdam, serif;
}

.v-container-leave-active {
  transition-property: opacity;
  transition-duration: 0.8s;
  transition-timing-function: linear;
}
.v-container-leave-to {
  opacity: 0;
}

.v-title-enter-active, .v-title-leave-active {
  transition-property: opacity;
  transition-duration: 0.6s;
  transition-timing-function: linear;
}
.v-title-enter-from, .v-title-leave-to {
  opacity: 0;
}
</style>