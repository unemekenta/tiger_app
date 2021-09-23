<template lang="pug">
.menu02
  .menu02-content
    .menu02-content-head
      | メニュー
    .menu02-content-items(v-if="signedIn")
      .menu02-content-items-item
        router-link( to="/" ) HOME
      .menu02-content-items-item
        router-link( to="/website_list" ) メディア一覧
      .menu02-content-items-item
        router-link( to="/mypage" ) マイページ
      .menu02-content-items-item(@click="signOut") ログアウト
    
    .menu02-content-items(v-else)
      .menu02-content-items-item
        router-link( to="/" ) HOME
      .menu02-content-items-item
        router-link( to="/website_list" ) メディア一覧
      .menu02-content-items-item
        router-link( to="/signin" ) ログイン

</template>

<script>
import axios from 'axios'

export default {
  name: 'Menu02',
  created () {
    this.loginCheck();
  },
  data () {
    return {
      signedIn: false,
    }
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async loginCheck () {
      var data =''
      if (window.$cookies.isKey('jwt')) {
        await axios.post('http://localhost:8000/auth/api/authCheck', data,
        {
          headers: {'Authorization': 'Bearer ' + window.$cookies.get('jwt')}
        })
          .then(() => {
            this.signedIn = true;
          })
          .catch(() => {
            this.signedIn = false;
          });
      }
    }
  },
}
</script>

<style scoped lang="scss" src="../assets/css/menu02.scss">
</style>

