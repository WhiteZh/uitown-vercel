<script setup lang="ts">

import {Ref, ref} from "vue";
import {notifications} from "@/globs";

const props = defineProps<{
  f: () => Promise<void>
}>();

const button = ref() as Ref<HTMLButtonElement>;

defineExpose(button);

async function submit() {
  button.value.setAttribute("disabled", "disabled");
  await props.f();
  button.value.removeAttribute("disabled");
}
</script>

<template>
  <button ref="button" @click="submit">
    <slot/>
  </button>
</template>

<style scoped>

</style>