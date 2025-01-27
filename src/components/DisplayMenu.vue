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
  <div class="flex flex-col">
    <div class="grid grid-cols-[repeat(auto-fill,minmax(16rem,1fr))] gap-4">
      <DisplayCard v-for="each in list"
                   class="aspect-[6/5] max-w-96"
                   :name="each.name" :subscribed="0" :css="each.css" :html="each.html" :id="each.id" :key="each.id"/>
    </div>
  </div>
</template>

<style scoped>
</style>