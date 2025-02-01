<script setup lang="ts">
import NavigationBar from "@/components/NavigationBar.vue"
import CodeDisplay from "@/components/code/CodeDisplay.vue";
import {computed, nextTick, onMounted, ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";
import {createCSSStyle, deleteCSSStyle, getCSSByIds, updateCSSStyle} from "@/api";
import {notifications, user} from "@/globs";
import {cssCategories, CSSCategory} from "@/constants";
import SubmitButton from "@/components/SubmitButton.vue";
import {isMatching, match, P} from "ts-pattern";

const route = useRoute();
const router = useRouter();

const mode = computed<'create' | 'view' | undefined>(() => match(route.meta.mode)
    .with(P.union('create', 'view'), it => it)
    .otherwise(() => undefined));
const codeID = computed<number>(() => parseInt(route.params.id as string));

const nameInput = ref<HTMLInputElement>();
const categoryInput = ref<HTMLSelectElement>();

const html = ref('');
const css = ref('');
const authorID = ref<number>();
const name = ref('');
const category = ref<CSSCategory | 'category'>('category');

if (mode.value === "create" && user.value === undefined) {
  router.push({name: 'browse', params: {category: ''}});
  notifications.push({
    message: 'Please login before creating new styles',
    color: 'yellow'
  });
}

// fetch data for 'view' mode
watch(() => route.fullPath, async () => {
  match(mode.value)
      .with('view', async () => {
        try {
          if (Array.isArray(route.params.id)) {
            notifications.push({message: 'Something wrong happened', color: 'yellow'});
            console.log(`\`route.params.id\` is an array, full path: ${route.fullPath}`);
            return;
          }

          let styles = await getCSSByIds([parseInt(route.params.id)]);
          if (styles.length > 0) {
            html.value = styles[0].html;
            css.value = styles[0].css;
            authorID.value = styles[0].author_id;
            name.value = styles[0].name;
            category.value = styles[0].category;
          } else {
            notifications.push({message: 'Id does not exist', color: 'yellow'});
            await router.push({name: 'browse', params: {category: ''}});
            return;
          }
        } catch (e) {
          console.log(e);
          notifications.push({message: `Failed to fetch data from server ${String(e)}`, color: 'red'});
        }
      })
      .with('create', () => {
        html.value = '';
        css.value = '';
        authorID.value = undefined;
        name.value = '';
        category.value = 'category';
      })
      .otherwise(() => undefined);
}, {immediate: true});

async function submit() {
  console.log(user);
  if (!user.value) {
    notifications.push({message: 'Not logged in', color: 'yellow'});
    return;
  }
  try {
    if (mode.value === "create") {
      if (!isMatching(P.union(...cssCategories), category.value)) {
        notifications.push({message: "Please select a category before submit", color: "yellow"});
        return;
      }
      await createCSSStyle(
          user.value.id,
          user.value.password_hashed,
          name.value,
          category.value,
          html.value,
          css.value
      );
      notifications.push({message: 'Successfully created a new style'});
    } else {
      if (!isMatching(P.union(...cssCategories), category.value)) {
        notifications.push({message: "Please select a category before submit", color: "yellow"});
        return;
      }
      let err = await updateCSSStyle(codeID.value!, user.value.password_hashed, {
        name: name.value,
        html: html.value,
        css: css.value,
        category: category.value,
      });

      if (err) {
        throw err.message;
      }

      notifications.push({message: 'Successfully updated the style'});
    }
  } catch (e) {
    console.log(e);
    notifications.push({message: `Upload failed: ${String(e)}`, color: 'yellow'});
  }
}

const del = async () => {
  if (!user.value) {
    notifications.push({message: 'Login first', color: 'yellow'});
    return;
  }
  if (codeID.value === undefined) {
    notifications.push({message: "Cannot delete a style that hasn't been uploaded", color: "yellow"});
    return;
  }
  try {
    await deleteCSSStyle(codeID.value, user.value.password_hashed);
    notifications.push({message: 'successfully deleted'});
  } catch (e) {
    notifications.push({message: 'deletion failed'});
    console.log(e);
  }
}

</script>

<template>
  <NavigationBar/>
  <div class="m-8 mx-auto max-w-screen-2xl px-3">
    <!--    <RouterLink :to="{name: 'browse'}" class="text-black bg-white py-3 inline-flex w-28 text-center rounded-full flex-row items-center justify-center">-->
    <!--      <svg xmlns="http://www.w3.org/2000/svg" height="100%" fill="currentColor" class="bi bi-arrow-left  h-4 me-1" viewBox="0 0 16 16">-->
    <!--        <path fill-rule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8"/>-->
    <!--      </svg>-->
    <!--      <span class="me-0.5">Go Back</span>-->
    <!--    </RouterLink>-->
    <div
        class="bg-neutral-800 rounded-full px-6 h-20 -mx-0.5 flex flex-row justify-start items-center gap-4"
        v-if="mode === 'create' || user !== undefined && authorID === user.id"
    >
      <div class="flex justify-start items-center">
        <input
            class="h-12 rounded-full w-52 text-sm px-3 text-center font-mono"
            ref="nameInput"
            v-model="name"
            placeholder="name"/>
      </div>
      <div class="flex justify-end items-center">
        <select
            ref="categoryInput"
            class="h-12 rounded-full w-40 px-3 text-center font-mono text-sm"
            :class="{'text-gray-400': category === 'category'}"
            v-model="category"
        >
          <option disabled selected>category</option>
          <option v-for="cate in cssCategories">{{ cate }}</option>
        </select>
      </div>
      <SubmitButton
          class="bg-highlight1 disabled:bg-zinc-500 px-4 h-12 flex items-center rounded-full text-white tracking-widest text-base"
          :f="submit"
      >
        SUBMIT
      </SubmitButton>
      <SubmitButton
          class="bg-red-800 disabled:bg-zinc-500 px-4 h-12 flex items-center rounded-full text-white tracking-widest text-base"
          :f="del"
          v-if="user !== undefined && authorID !== undefined && user.id === authorID"
      >
        DELETE
      </SubmitButton>
    </div>

    <CodeDisplay
        class="rounded-2xl overflow-hidden h-[70svh] mt-4 mx-auto"
        v-model:html="html" v-model:css="css"/>
  </div>
</template>

<style scoped>
</style>