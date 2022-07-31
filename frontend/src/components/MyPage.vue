<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .user
        //- img.user-icon(src="@/assets/images/logo.png")
        .user-name
          b {{userName}}
          |さん
      .main
        menu01-item(:allCategories="allCategories").is-only-pc
        .main-mypage
          h2.main-mypage-title マイページ
          .main-mypage-contents
            p 準備中...
          h3 登録中のサブスク
          subscription-item(:subscriptions="subscriptions")
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
        menu02-item.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios';
import jwtDecode from 'jwt-decode';
import Menu01Item from '/src/organisms/Menu01.vue';
import Menu02Item from '/src/organisms/Menu02.vue';
import SubscriptionItem from '/src/organisms/Subscription.vue';
import HeaderNav from '/src/molecules/HeaderNav.vue';
import FooterNav from '/src/molecules/FooterNav.vue';

export default {
  name: 'MyPage',
  components: {
    Menu01Item,
    Menu02Item,
    SubscriptionItem,
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      userName: '',
      subscriptions: [],
      jwtUserData: '',
      userID: 0,
    }
  },
  async created () {
    await this.getUser();
    this.getSubscriptions();
  },
  methods: {
    async getSubscriptions () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/subscription/user/' + this.userID,
      {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(res => {
        this.subscriptions = res.data;
        return;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    async getUser () {
      if (window.$cookies.isKey('uuid')) {
        await axios
          .post(
            process.env.VUE_APP_API_BASE_URL + "/api/login_check",
            {
              uuid: window.$cookies.get('uuid')
            },
          )
          .then((res) => {
            this.jwtUserData = res.data;
          })
          .catch(error => {
            alert('ログインしてください', error);
            return;
          });

        let decodedHeader = jwtDecode(this.jwtUserData);
        this.userID = decodedHeader.id;
        await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/user/'+ this.userID,
        {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
          .then((res) => {
            this.userName = res.data.name;
            return;
          })
          .catch(error => {
            alert('権限がありません。');
            return error;
          });
        return;
      } else {
        alert('ログインしてください');
        return;
      }
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/myPage.scss">
</style>
