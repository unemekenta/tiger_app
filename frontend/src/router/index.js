import Vue from 'vue'
import Router from 'vue-router'
import TopList from '@/components/TopList'
import WebsiteList from '@/components/WebsiteList'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import MyPage from '@/components/MyPage'
import WebsiteDetail from '@/components/WebsiteDetail'
// import Edit from '@/components/Edit'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'TopList',
      component: TopList
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/mypage',
      name: 'MyPage',
      component: MyPage,
      meta: { requiresAuth: true }
    },
    {
      path: '/website_detail/:id',
      name: 'WebsiteDetail',
      component: WebsiteDetail
    },
    {
      path: '/website_list',
      name: 'WebsiteList',
      component: WebsiteList
    }
    // {
    //   path: '/edit',
    //   name: 'Edit',
    //   component: Edit,
    //   meta: { requiresAuth: true }
    // }
  ]
})

export default router
