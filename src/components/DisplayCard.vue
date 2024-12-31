<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";
import {shadowContent} from '@/constants';
import {useRouter} from "vue-router";

const router = useRouter();

const props = defineProps<{
  name: string,
  subscribed: number,
  html: string,
  css: string,
  id: number
}>();

const cardContentRoot = ref() as Ref<HTMLDivElement>;
onMounted(() => {
  let shadowDOM = cardContentRoot.value.attachShadow({mode: 'open'});
  shadowDOM.innerHTML = shadowContent(props.html, props.css);
  const calcHeightRatio = (x: number) => cardContentRoot.value.offsetHeight * 0.9 / x;
  const calcWidthRatio = (x: number) => cardContentRoot.value.offsetWidth * 0.9 / x;
  let ratio = Infinity;
  for (let each of Array.from(shadowDOM.querySelectorAll<HTMLElement>('*')).filter((e) => getComputedStyle(e).position === 'absolute')) {
    ratio = Math.min(ratio, calcHeightRatio(each.offsetHeight));
    ratio = Math.min(ratio, calcWidthRatio(each.offsetWidth));
  }
  if (ratio < 1) {
    (shadowDOM.getElementById('the-id-of-the-shadow-root') as HTMLDivElement).style.transform += ` scale(${ratio}) `;
  }
})
</script>

<template>
  <div class="flex flex-col">
    <div class="flex-grow relative overflow-hidden group">
      <div class="rounded-3xl h-full relative bg-[#2b2a2a]" ref="cardContentRoot"/>
      <RouterLink class="card-btn absolute bottom-4 right-5 bg-white rounded-md h-8 w-20 hidden group-hover:inline-flex flex-row items-center justify-center gap-1.5 text-black cursor-pointer" :to="{name: 'view', params: {id: props.id}}">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-braces" viewBox="0 0 16 16">
          <path d="M2.114 8.063V7.9c1.005-.102 1.497-.615 1.497-1.6V4.503c0-1.094.39-1.538 1.354-1.538h.273V2h-.376C3.25 2 2.49 2.759 2.49 4.352v1.524c0 1.094-.376 1.456-1.49 1.456v1.299c1.114 0 1.49.362 1.49 1.456v1.524c0 1.593.759 2.352 2.372 2.352h.376v-.964h-.273c-.964 0-1.354-.444-1.354-1.538V9.663c0-.984-.492-1.497-1.497-1.6M13.886 7.9v.163c-1.005.103-1.497.616-1.497 1.6v1.798c0 1.094-.39 1.538-1.354 1.538h-.273v.964h.376c1.613 0 2.372-.759 2.372-2.352v-1.524c0-1.094.376-1.456 1.49-1.456V7.332c-1.114 0-1.49-.362-1.49-1.456V4.352C13.51 2.759 12.75 2 11.138 2h-.376v.964h.273c.964 0 1.354.444 1.354 1.538V6.3c0 .984.492 1.497 1.497 1.6"/>
        </svg>
        <span style="padding-top: 0.15rem;">Code</span>
      </RouterLink>
    </div>
    <div class="pt-1.5 px-3 flex flex-row justify-between items-center">
      <div class="text-white text-sm">
        {{name}}
      </div>
      <div class="text-white text-xs">
        subs: {{ subscribed }}
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>