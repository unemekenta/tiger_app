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
  var data =''
  if (requiresAuth) {
    // todo サインインの有無判断
    await axios.post('http://localhost:8000/auth/api/authCheck', data,
    {
      headers: {'Authorization': 'Bearer ' + window.$cookies.get('jwt')}
    })
      .then(() => {
        next()
      })
      .catch(error => {
        alert(error)
        console.log(window.$cookies.get('jwt'))
        next({
          path: '/signin'
        })
      });
  } else {
    next()
  }
})
export default router
