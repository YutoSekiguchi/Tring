import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import axios from 'axios' //追記
import VueAxios from 'vue-axios' //追記

Vue.config.productionTip = false

Vue.use(VueAxios, axios) // 追記

axios.defaults.baseURL = "http://localhost:7121" // add(apiのurlをここに記す．今回はDockerで7121ポートを使ってるからこう書く)

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
