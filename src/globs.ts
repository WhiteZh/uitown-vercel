import {Notification, User} from "@/constants";
import {reactive, ref, Ref} from "vue";

const globs: {
    notifications: Notification[],
    user: Ref<User|undefined>,
    playedOA: boolean
} = {
    notifications: reactive([]),
    user: ref<User>(),
    playedOA: false,
}

export const notifications: Notification[] = reactive<Notification[]>([]);

export const user: Ref<User | undefined> = ref<User>();

export const playedOA: Ref<boolean> = ref<boolean>(false);