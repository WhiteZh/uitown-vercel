<script setup lang="ts">
import {inject, onMounted, Ref, ref} from "vue";
import DisplayCard from "@/components/DisplayCard.vue";
import {CSSCategory, CSSStyle, Notification} from "@/constants";
import {getCSSByIds, getValidCSSIds} from "@/api";
import {notifications} from "@/globs";

const props = withDefaults(defineProps<{
  contentType?: "css" | "js",
  category?: CSSCategory,
  hasSearcher?: boolean,
}>(), {
  hasSearcher: true
});

const list: Ref<CSSStyle[]> = ref([]);

onMounted(async function() {
  try {
    let ids = await getValidCSSIds({
      category: props.category
    });
    list.value = await getCSSByIds(ids);
  } catch (e) {
    notifications.push({message: String(e), color: 'red'});
    console.log(e);
  }
});
</script>

<template>
  <div class="flex-grow flex flex-col overflow-scroll [scrollbar-width:none]">
    <div class="flex flex-row-reverse justify-start italic text-white h-5 mb-3" v-if="hasSearcher">
      <div class="relative flex">
        <svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" fill="currentColor" viewBox="0 0 16 16" class="h-full absolute left-[0.4rem] text-black cursor-pointer">
          <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001q.044.06.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1 1 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0"/>
        </svg>
        <input class="my-[1px] rounded-full ps-[1.3rem] text-sm font-thin text-black outline-0"/>
      </div>
      <span class="text-sm font-[Cooljazz] italic self-end leading-tight me-5 tracking-widest"><span class="text-[0.5rem]">Sort : </span>Randomized</span>
    </div>
    <div class="grid grid-cols-[repeat(auto-fill,minmax(18rem,1fr))] grid-rows-[repeat(auto-fit,15rem))] gap-4 justify-items-center">
      <DisplayCard v-for="each in list" class="w-72 h-60" :name="each.name" :subscribed="0" :css="each.css" :html="each.html" :id="each.id"/>
    </div>
    <div class="footer">

    </div>
  </div>
</template>

<style scoped>
</style>