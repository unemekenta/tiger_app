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
    }
  },
  created () {
    this.getUser();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
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
      let user_id = VueJwtDecode.decode(window.$cookies.get('jwt')).id;
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/auth/api/user/'+ user_id,
      {
        headers: {'Authorization': 'Bearer ' + window.$cookies.get('jwt')}
      })
        .then((res) => {
          this.userName = res.data.name;
        })
        .catch(error => {
          alert(error)
        });
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/myPage.scss">
</style>
