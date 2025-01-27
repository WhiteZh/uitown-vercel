<script setup lang="ts">
import {ref, inject, onMounted, Ref} from "vue";
import {sha256} from "js-sha256";
import {getUserById, getUserIdByLoginInfo} from "@/api";
import {notifications, user} from "@/globs";
import SubmitButton from "@/components/SubmitButton.vue";

const emailInput = ref() as Ref<HTMLInputElement>;
const passwordInput = ref() as Ref<HTMLInputElement>;
const loginButton = ref() as Ref<HTMLButtonElement>;

const emits = defineEmits(['login']);

const login = async () => {
  let email = emailInput.value.value;
  let password_hashed = sha256(passwordInput.value.value);
  let userId: number|undefined = undefined;
  try {
    userId = await getUserIdByLoginInfo(email, password_hashed);
  } catch (e) {
    console.log(e);
    notifications.push({
      message: `Failed to login (400): ${String(e)}`,
      color: 'yellow'
    });
    emits('login', false);
  }
  if (userId) {
    let res = await getUserById(userId, password_hashed);
    if (res instanceof Error) {
      notifications.push({message: res.message});
    } else {
      user.value = res;
      notifications.push({message: 'Successfully logged in'});
      emits('login', true);
    }
  } else {
    notifications.push({
      message: "Username or password incorrect",
      color: 'red'
    });
    emits('login', false);
  }
}

onMounted(() => {
  emailInput.value.focus();
})
</script>

<template>
  <div class="h-[368px] w-[560px] absolute bg-[linear-gradient(90deg,#8c52ff,#5ce1e6)] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-[2rem] flex flex-col">
    <div class="font-['Zhi_Mang_Xing'] text-5xl text-white py-4 self-center mb-5">UITOWN</div>
    <div class="flex flex-col">
      <div class="flex flex-row mb-3 justify-center">
        <input placeholder="username/email"
               class="h-11 w-80 rounded-full my-1.5 ms-1.5 outline-0 ps-0.5 font-mono text-xs text-center"
               ref="emailInput" @keydown.enter="() => passwordInput.focus()"/>
      </div>
      <div class="flex flex-row mb-3 justify-center">
        <input placeholder="password" type="password" class="h-11 w-80 rounded-full my-1.5 ms-1.5 outline-0 ps-0.5 font-mono text-xs text-center" ref="passwordInput" @keydown.enter="() => loginButton.click()">
      </div>
      <SubmitButton
          class="mb-7 h-10 px-5 rounded-full bg-stress disabled:bg-zinc-500 text-white font-bold text-sm self-center my-2"
          :f="login" ref="loginButton"
      >
        Login/Registration
      </SubmitButton>
      <div class="buttons flex flex-row justify-center *:text-stress *:w-8 *:h-8 *:cursor-pointer *:mx-0.5">
        <button>
          <i class="bi bi-github"></i>
        </button>
        <button>
          <i class="bi bi-twitter-x"></i>
        </button>
        <button>
          <i class="bi bi-instagram"></i>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
















