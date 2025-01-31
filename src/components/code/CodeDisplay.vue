<script setup lang="ts">
import {EditorView, basicSetup} from "codemirror";
import { oneDark } from "@codemirror/theme-one-dark";
import {EditorState, Extension} from "@codemirror/state";
import { html } from '@codemirror/lang-html';
import { css } from '@codemirror/lang-css';
import { iframeContent } from '@/constants'
import {onMounted, Ref, ref, watch} from "vue";
import SubmitButton from "@/components/SubmitButton.vue";

const props = withDefaults(defineProps<{
  html: string,
  css: string,
  deletion?: () => Promise<void>,
}>(), {
  html: '',
  css: '',
});

const emit = defineEmits(['update:html', 'update:css']);

const activeTab: Ref<"html" | "css"> = ref('html');

const htmlEditor = ref() as Ref<HTMLDivElement>;
const cssEditor = ref() as Ref<HTMLDivElement>;

let iframeValue = ref(iframeContent('', ''));

onMounted(() => {
  let htmlView: EditorView | null = null;
  let cssView: EditorView | null = null;

  // number 12 is autocompletion
  let setupExtensions: Extension[] = [
      ...(basicSetup as Extension[]).filter((_, i) => ![4, 12].includes(i)),
      oneDark,
      EditorView.domEventHandlers({
        change: () => {
          if (htmlView === null || cssView === null) return;

          let html = htmlView.state.doc.toString();
          let css = cssView.state.doc.toString();
          iframeValue.value = iframeContent(html, css);
          emit('update:html', html);
          emit('update:css', css);
        }
      })
  ];
  const htmlExtensions = [setupExtensions, html()];
  const cssExtensions = [setupExtensions, css()];

  const htmlState = EditorState.create({
    doc: props.html,
    extensions: htmlExtensions,
  });

  htmlView = new EditorView({
    state: htmlState,
    parent: htmlEditor.value,
  });

  const cssState = EditorState.create({
    doc: props.css,
    extensions: cssExtensions,
  });

  cssView = new EditorView({
    state: cssState,
    parent: cssEditor.value,
  });

  [{src: () => props.html, view: htmlView}, {src: () => props.css, view: cssView}].forEach(({src, view}) => {
    watch(src, () => {
      if (src() === view.state.doc.toString()) {
        return;
      }
      // keep cursor location even if view got dispatched/updated
      // let cursor = Math.min(view.state.selection.main.head, src().length);
      view.dispatch({
        changes: {
          from: 0,
          to: view.state.doc.length,
          insert: src(),
        },
        // selection: {
        //   anchor: cursor,
        // }
      });
    });
  });

  let lastHTML = '';
  let lastCSS = '';
  setInterval(() => {
    let html = htmlView.state.doc.toString();
    let css = cssView.state.doc.toString();
    if (html !== lastHTML || css !== lastCSS) {
      iframeValue.value = iframeContent(html, css);
      emit('update:html', html);
      emit('update:css', css);
      lastHTML = html;
      lastCSS = css;
    }
  }, 500);
});


</script>

<template>
  <div class="flex flex-row">
    <iframe class="w-1/2" :srcdoc="iframeValue"/>
    <div class="w-1/2 flex flex-col">
      <div class="min-h-20 max-h-20 bg-[#545454] flex flex-row justify-start items-center px-8 gap-6 text-white text-[1.2rem] font-[Cooljazz] italic tracking-[0.3rem]">
        <button class="w-32 h-16 rounded-full cursor-pointer bg-[#7ed957]" @click="activeTab = 'html'">HTML</button>
        <button class="w-32 h-16 rounded-full cursor-pointer bg-[#ff66c4]" @click="activeTab = 'css'">CSS</button>
        <div style="flex-grow: 1"/>
        <SubmitButton class="w-32 h-16 rounded-full bg-[#550000] disabled:bg-zinc-500" :f="deletion" v-if="deletion !== undefined">Delete</SubmitButton>
      </div>
      <div class="flex-grow bg-[#272727] pt-1 overflow-scroll" ref="htmlEditor" :style="{display: activeTab === 'html' ? 'block' : 'none'}"></div>
      <div class="flex-grow bg-[#272727] pt-1 overflow-scroll" ref="cssEditor" :style="{display: activeTab === 'css' ? 'block' : 'none'}"></div>
    </div>
  </div>
</template>

<style scoped>
</style>