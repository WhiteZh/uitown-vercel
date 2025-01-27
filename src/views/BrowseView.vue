<script setup lang="ts">
import {ref} from "vue";
import NavigationBar from "@/components/NavigationBar.vue";
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
  // {name: 'Transitions', url:{name: 'browse', params: {content_type: 'css', category: 'transition'}}, colors: ['#004add', '#cb6ce6']},
  {name: 'Special Effects', url: {name: 'browse', params: {content_type: 'css', category: 'special effect'}}, colors: ['#0097b2', '#7ed957']}
]);

</script>

<template>
  <div class="flex flex-col items-stretch justify-stretch">
    <NavigationBar/>
    <div class="flex flex-row items-stretch pt-5 flex-grow">
      <div class="pt-5 min-w-48 px-5">
        <RouterLink :to="{name: 'browse'}" class="text-white text-lg font-thin">All</RouterLink>
      </div>
      <div class="flex-grow px-5 flex flex-col items-stretch">
        <h1 class="text-3xl my-7 font-medium text-primary px-2">Browsing All</h1>
        <DisplayMenu :contentType="contentType" :category="category" class="flex-grow"/>
      </div>
    </div>
  </div>

</template>

<style scoped>
</style>
