import './assets/main.css'

import {createApp, reactive, ref} from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)

app.mount('#app')
