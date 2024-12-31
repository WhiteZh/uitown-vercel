<script setup lang="ts">
import { RouterView } from 'vue-router'
import {notifications} from "@/globs";
import {onMounted} from "vue";

onMounted(() => {
  if (window.screen.width < 704) {
    window.alert("This screen width is currently not supported");
  }
});
</script>

<template>
  <div class="absolute z-[1000] w-full">
    <div v-for="each in notifications" :style="{backgroundColor: each.color || 'lightblue'}" class="text-center py-1.5 px-0 opacity-80 relative">
      {{each.message}}
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x hover:text-[#333] cursor-pointer" viewBox="0 0 16 16" style="position: absolute; right: 0; height: 100%; top: 0; margin-right: 0.5rem;" @click="notifications.splice(notifications.indexOf(each), 1)">
        <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708"/>
      </svg>
    </div>
  </div>
  <RouterView :key="$route.fullPath"/>
</template>


<style scoped>
</style>
