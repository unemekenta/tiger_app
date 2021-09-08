<template lang="pug">
.top
  .user
    img.user-icon(src="../../src/assets/images/logo.png")
    .user-name 
      b {{userName}} 
      |さんのページ
  .main
    menu01(:allCategories="allCategories")
    .main-mypage
      h2.main-mypage-title マイクリップ
      .main-mypage-contents
        .main-mypage-contents-item
        .main-mypage-contents-item
        .main-mypage-contents-item
        .main-mypage-contents-item
        .main-mypage-contents-item
        .main-mypage-contents-item
    menu02

</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import axios from 'axios'

export default {
  name: 'MyPage',
  components: {
    Menu01,
    Menu02,
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
      let resAllCategories = await axios.get('http://localhost:8000/api/categories');
      this.allCategories = resAllCategories.data;
    },
    async getUser () {
      let user_id = VueJwtDecode.decode(window.$cookies.get('jwt')).id;
      console.log(window.$cookies.get('jwt'));
      await axios.get('http://localhost:8000/auth/api/user/'+ user_id,
      {
        headers: {'Authorization': 'Bearer ' + window.$cookies.get('jwt')}
      })
        .then((res) => {
          this.userName = res.data.name;
          console.log(res);
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
