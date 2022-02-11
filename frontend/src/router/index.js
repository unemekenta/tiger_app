import Vue from 'vue'
import Router from 'vue-router'
import TopList from '@/components/TopList'
import PrivacyPolicy from '@/components/PrivacyPolicy'
import TermsOfService from '@/components/TermsOfService'
import WebsiteList from '@/components/WebsiteList'
import SignupItem from '@/components/Signup'
import SigninItem from '@/components/Signin'
import MyPage from '@/components/MyPage'
import WebsiteDetail from '@/components/WebsiteDetail'
import WebsiteListByCategory from '@/components/WebsiteListByCategory'
import MyAsset from '@/components/MyAsset'

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
      component: SignupItem,
      meta: { title: 'サインアップ' }
    },
    {
      path: '/signin',
      name: 'Signin',
      component: SigninItem,
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
    },
    {
      path: '/my-asset',
      name: 'MyAsset',
      component: MyAsset,
      meta: {
        title: '今月いくら？TOP',
        requiresAuth: true
      },
    },
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
