import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueCookies from 'vue-cookies'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'

declare const window: Window

Vue.use(VueCookies)

Vue.config.productionTip = false

Vue.component('fa-icon', FontAwesomeIcon)
library.add(fas, far, fab)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

router.beforeEach(async (to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const data =''
  if (requiresAuth) {
    // todo サインインの有無判断
    if (window.$cookies.isKey('uuid')) {
      await axios
        .post(
          process.env.VUE_APP_API_BASE_URL + "/api/login_check",
          {
            uuid: window.$cookies.get('uuid')
          },
        )
        .then(() => {
          next();
        })
        .catch(() => {
          alert('ログインしてください')
          next({
            path: '/signin'
          });
        });

      await axios
        .post(
          process.env.VUE_APP_API_BASE_URL + '/api/auth/authCheck',
          data,
          {
            headers: {
              'Authorization': 'Bearer ' + this.jwtData,
            },
          })
        .then(() => {
          next();
        })
        .catch(() => {
          alert('権限がありません。');
          next({
            path: ''
          });
        });

    } else {
      alert('セッションが切れました。ログインしてください');
      next({
        path: '/signin'
      });
    }
  } else {
    next()
  }
})
export default router
