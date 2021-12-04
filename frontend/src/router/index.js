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
      component: TopList,
      meta: { title: 'TOPページ' }
    },
    {
      path: '/privacy_policy',
      name: 'PrivacyPolicy',
      component: PrivacyPolicy,
      meta: { title: 'プライバシーポリシー' }
    },
    {
      path: '/terms_of_service',
      name: 'TermsOfService',
      component: TermsOfService,
      meta: { title: 'TermsOfService' }
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup,
      meta: { title: 'サインアップ' }
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin,
      meta: { title: 'サインイン' }
    },
    {
      path: '/mypage',
      name: 'MyPage',
      component: MyPage,
      meta: {
        title: 'マイページ',
        requiresAuth: true
      },
    },
    {
      path: '/website_detail/:id',
      name: 'WebsiteDetail',
      component: WebsiteDetail,
      meta: { title: 'ウェブサイト詳細' }
    },
    {
      path: '/website_list',
      name: 'WebsiteList',
      component: WebsiteList,
      meta: { title: 'ウェブ一覧' }
    },
    {
      path: '/website_list/:category_id',
      name: 'WebsiteListByCategory',
      component: WebsiteListByCategory
      ,meta: { title: 'カテゴリー別一覧' }
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
