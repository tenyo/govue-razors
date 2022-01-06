import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

const axiosClient = axios.create({
  baseURL: '/api/v1',
})

createApp(App).use(VueAxios, axiosClient).mount('#app')
