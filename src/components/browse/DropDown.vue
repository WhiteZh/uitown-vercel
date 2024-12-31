<script setup lang="ts">
import {RouterLink, useRouter} from "vue-router";
import {CSSCategory} from "@/constants";

type URL = string | {
  name: string,
  params?: {
    content_type?: 'css' | 'js',
    category?: CSSCategory,
  }
};

const props = defineProps<{
  list?: {
    name: string,
    url: URL,
    colors: string[],
  }[],
  self_url?: URL,
  extendWidth?: string,
}>();

const router = useRouter();

const gotoSelfURL = () => {
  if (props.self_url !== undefined) {
    router.push(props.self_url);
  }
}
</script>

<template>
  <div class="relative group">
    <div class="text-lg py-2 px-4 w-full inline-block cursor-pointer font-bold text-white text-left rounded-full group-hover:bg-[#272030]" @click="gotoSelfURL">
      <slot/>
    </div>
    <div class="group-hover:inline-block absolute hidden ps-2 group" v-if="list">
      <ul class="bg-[#272030] rounded-2xl py-0.5 px-2" :style="{width: extendWidth}">
        <li v-for="{name, colors, url} in list" class="my-3 text-white list-none hover:underline" :key="name">
          <RouterLink class="leading-tight text-white text-sm font-thin text-center block py-2 px-4 rounded-full" :to="url" :style="{ background: `linear-gradient(90deg, ${colors.join(',')})` }">{{name}}</RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
</style>