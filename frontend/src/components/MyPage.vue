<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .user
        //- img.user-icon(src="../../src/assets/images/logo.png")
        .user-name
          b {{userName}}
          |さんのページ
      .main
        menu01(:allCategories="allCategories").is-only-pc
        .main-mypage
          h2.main-mypage-title マイクリップ
          .main-mypage-contents
           p 準備中..
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
        menu02.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'MyPage',
  components: {
    Menu01,
    Menu02,
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      userName: '',
      allCategories: [],
      jwtUserData: '',
    }
  },
  created () {
    this.getUser();
    this.getAllCategories();
  },
  methods: {
    async getAllCategories () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/categories')
      .then(res => {
        this.allCategories = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    async getUser () {
      if (window.$cookies.isKey('uuid')) {
        await axios
          .get(process.env.VUE_APP_API_BASE_URL + "/api/login_check", {
            params: {
              key: window.$cookies.get('uuid')
            },
          })
          .then((res) => {
            this.jwtUserData = res.data;
          })
          .catch(error => {
            alert('ログインしてください', error);
            return;
          });

        let user_id = VueJwtDecode.decode(this.jwtUserData).id;
        await axios.get(process.env.VUE_APP_API_BASE_URL + '/auth/api/user/'+ user_id,
        {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
          .then((res) => {
            this.userName = res.data.name;
          })
          .catch(error => {
            alert('権限がありません', error);
            return;
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
