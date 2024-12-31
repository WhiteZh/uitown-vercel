<script setup lang="ts">
import {ref} from "vue";
import NavigationBar from "@/components/NavigationBar.vue";
import DropDown from "@/components/browse/DropDown.vue";
import DisplayMenu from "@/components/DisplayMenu.vue";
import {useRoute, useRouter} from "vue-router";
import {CSSCategory, isCSSCategory} from "@/constants";
import {notifications} from "@/globs";

let route = useRoute();
let router = useRouter();

let contentType: 'css' | 'js' | undefined = undefined;
if (route.params['content_type'] === 'css' || route.params['content_type'] === 'js') {
  contentType = route.params['content_type'];
} else if (route.params['content_type'] !== '' && route.params['content_type'] !== undefined) {
  notifications.push({message: 'Path does not exist', color: 'red'});
  console.log(route.params['content_type']);
  router.push({name: 'browse'});
}
let category: CSSCategory | undefined = undefined;
if (isCSSCategory(route.params['category'])) {
  category = route.params['category'];
} else if (route.params['category'] !== '' && route.params['category'] !== undefined) {
  notifications.push({message: 'Illegal Path', color: 'red'});
  console.log(route.params['category']);
  router.push({name: 'browse'});
}

let CSS_attribute = ref<{
  name: string,
  url: string | {
    name: string,
    params?: {
      content_type?: 'css' | 'js',
      category?: CSSCategory,
    }
  },
  colors: string[],
}[]>([
  {name: 'Buttons', url: {name: 'browse', params: {content_type: 'css', category: 'button'}}, colors: ['#8c52ff', '#ff914d']},
  {name: 'Checkboxes', url: {name: 'browse', params: {content_type: 'css', category: 'checkbox'}}, colors: ['#ff5757', '#8c52ff']},
  {name: 'Toggle Switches', url: {name: 'browse', params: {content_type: 'css', category: 'toggle switch'}}, colors: ['#8c52ff', '#5ce1e6']},
  {name: 'Loaders', url: {name: 'browse', params: {content_type: 'css', category: 'loader'}}, colors: ['#8c52ff', '#00bf63']},
  {name: 'Cards', url: {name: 'browse', params: {content_type: 'css', category: 'card'}}, colors: ['#5170ff', '#ff66c4']},
  {name: 'Inputs', url: {name: 'browse', params: {content_type: 'css', category: 'input'}}, colors: ['#5de0e6', '#004aad']},
  {name: 'Transitions', url:{name: 'browse', params: {content_type: 'css', category: 'transition'}}, colors: ['#004add', '#cb6ce6']},
  {name: 'Special Effects', url: {name: 'browse', params: {content_type: 'css', category: 'special effect'}}, colors: ['#0097b2', '#7ed957']}
]);
let JS_attribute = ref([
  {name: 'Mouse Effect', url: {name: 'browse'}, colors: ['#8c52ff', '#ff914d']},
  {name: 'Background', url: {name: 'browse'}, colors: ['#ff5757', '#8c52ff']},
  {name: 'Menu', url: {name: 'browse'}, colors: ['#8c52ff', '#5ce1e6']},
  {name: 'Visible Chart', url: {name: 'browse'}, colors: ['#8c52ff', '#00bf63']},
]);

</script>

<template>
  <NavigationBar/>
  <div class="py-0 px-60 text-[#D0C3F1]">
    <h3 class="text-2xl mt-7 mb-1.5 mx-0 font-bold">To Select</h3>
    <h6 class="font-[Cooljazz] tracking-[0.2rem] font-thin italic indent-8 my-3 mx-0 text-xs">Choose the code of your choice</h6>
  </div>
  <div class="flex flex-row flex-grow">
    <div class="pe-36 flex flex-col ms-5 items-stretch">
      <DropDown :self_url="{name: 'browse'}">All</DropDown>
      <DropDown :list="CSS_attribute" :self_url="{name: 'browse', params: {content_type: 'css'}}">CSS</DropDown>
      <DropDown :list="JS_attribute" :self_url="{name: 'browse'}">JavaScript</DropDown>
    </div>
    <DisplayMenu :contentType="contentType" :category="category" class="pe-20"/>
  </div>

</template>

<style scoped>
</style>
