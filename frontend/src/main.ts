import {createApp} from 'vue'
import App from './app/app.vue'
import './app/assets/styles/index.css'
import { createAppRouter } from './app/routes'

const app = createApp(App)
app.use(createAppRouter())
app.mount('#app')
