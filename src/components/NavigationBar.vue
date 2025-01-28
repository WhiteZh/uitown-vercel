<script setup lang="ts">
import Login from "@/components/Login.vue"
import {computed, inject, Ref, ref} from "vue";
import {RouterLink, useRouter} from "vue-router";
import {User} from "@/constants";
import {notifications, user} from "@/globs";

let openLogin = ref(false);

let router = useRouter();
</script>

<template>
  <div class="z-10 absolute h-screen w-screen opacity-30 bg-black" v-if="openLogin" @click="openLogin = !openLogin"></div>
  <Teleport to="#app">
    <Login v-if="openLogin" class="z-10" v-on:login="(v) => openLogin = v ? false : openLogin"/>
  </Teleport>
  <div class="px-4 pt-4 flex flex-row justify-between items-end" id="navbar">
    <RouterLink class="me-10" to="/"><img src="@/assets/logo.png" alt="UITOWN" class="h-12"/></RouterLink>
    <RouterLink :to="{name: 'browse', params: {category: ''}}" class="font-[Cooljazz] text-white hover:text-gray-300 text-sm italic me-7 -mb-0.5">Browse</RouterLink>
    <a href="#" class="font-[Cooljazz] text-white hover:text-gray-300 text-sm italic me-7 -mb-0.5">Information</a>
    <a href="#" class="font-[Cooljazz] text-white hover:text-gray-300 text-sm italic me-7 -mb-0.5">Recommendation</a>
    <span class="flex-grow"></span>
    <a href='#' class="text-sm self-center mx-1 py-1.5 px-3 rounded-full text-black font-bold bg-highlight1" @click.prevent="openLogin = !openLogin" v-if="!user">Join in the Town</a>
    <div v-if="user" class="bg-white self-center h-9 w-9 relative rounded-full mx-3">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16" style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); cursor: pointer;" @click="() => $router.push({name: 'user'})">
        <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6"/>
      </svg>
    </div>
    <RouterLink :to="user === undefined ? {} : {name: 'create'}" class="text-sm self-center mx-1 py-1.5 px-3 rounded-full text-black font-bold bg-highlight2" @click="user === undefined ? notifications.push({message:'Please login before creating new styles', color: 'yellow'}) : null">Create</RouterLink>
  </div>

</template>

<style scoped>
</style>