<script setup lang="ts">
import {computed, ref} from "vue";
import NavigationBar from "@/components/NavigationBar.vue";
import DisplayMenu from "@/components/DisplayMenu.vue";
import {useRoute, useRouter} from "vue-router";
import {cssCategories, CSSCategory, isCSSCategory} from "@/constants";
import {notifications} from "@/globs";
import {match, P} from "ts-pattern";
import StringUtilsWord from "string-utils-ts/lib/word/utils";

let route = useRoute();
let router = useRouter();

let category = computed<CSSCategory | undefined>(() => match(route.params['category'])
    .with("", () => undefined)
    .with(P.union(...cssCategories), it => it)
    .otherwise(it => {
      notifications.push({message: `Illegal path: ${route.fullPath}. Redirecting to /browse`, color: 'red'});
      console.log(`Illegal path: ${route.fullPath}`);
      console.log(`route.params['category'] = ${it}`)
      router.push({name: 'browse'});
      return undefined;
    }));

</script>

<template>
  <div class="flex flex-col items-stretch justify-stretch">
    <NavigationBar/>
    <div class="flex flex-row items-stretch pt-5 flex-grow px-5 gap-7 mb-4">
      <div class="pt-16 min-w-48 flex flex-col items-stretch gap-2 tracking-wide" ref="">
        <RouterLink v-for="cssCategory in [undefined, ...cssCategories]" :key="cssCategory ?? ''"
                    :to="{name: 'browse', params: {category: cssCategory ?? ''}}"
                    class="text-white text-md rounded-lg hover:bg-neutral-750 px-2 py-1.5 font-light"
                    :class="{'bg-neutral-800': cssCategory === category}"
        >{{ StringUtilsWord.formatWords(cssCategory ?? 'all') }}</RouterLink>
      </div>
      <div class="flex-grow flex flex-col items-stretch">
        <h1 class="text-3xl my-7 font-medium text-primary px-2">Browsing {{StringUtilsWord.formatWords(category ?? 'all')}}</h1>
        <DisplayMenu :category="category" class="flex-grow"/>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
