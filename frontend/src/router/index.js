import Vue from 'vue'
import Router from 'vue-router'
import TopList from '@/components/TopList'
import PrivacyPolicy from '@/components/PrivacyPolicy'
import TermsOfService from '@/components/TermsOfService'
import WebsiteList from '@/components/WebsiteList'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import MyPage from '@/components/MyPage'
import WebsiteDetail from '@/components/WebsiteDetail'
import WebsiteListByCategory from '@/components/WebsiteListByCategory'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'TopList',
      component: TopList
    },
    {
      path: '/privacy_policy',
      name: 'PrivacyPolicy',
      component: PrivacyPolicy
    },
    {
      path: '/terms_of_service',
      name: 'TermsOfService',
      component: TermsOfService
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
    },
    {
      path: '/website_list/:category_id',
      name: 'WebsiteListByCategory',
      component: WebsiteListByCategory
    }
  ],
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
       return savedPosition
    } else {
       return { x: 0, y: 0 }
    }
  }
})

export default router
