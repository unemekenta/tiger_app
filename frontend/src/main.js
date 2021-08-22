import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueCookies from 'vue-cookies'

Vue.use(VueCookies)

Vue.config.productionTip = false

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
        next({
          path: '/signin'
        })
      });
  } else {
    next()
  }
})
export default router
