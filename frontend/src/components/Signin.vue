<template lang="pug">
  .signin-page
    header-nav(:allCategories="allCategories")
    h2 Sign in
    .signin-page-form
      input.signin-page-form-email(type="email" placeholder="Email" v-model="formEmail")
      input.signin-page-form-password(type="password" placeholder="Password" v-model="formPassword")
      button.signin-page-form-button(@click="signIn") ログイン

    .signin-page-to-signin
      p アカウントをまだお持ちでない方はこちら
      router-link(to="/signup") アカウントを作成する

    footer-nav
</template>

<script>
import axios from 'axios'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'Signin',
  components: {
    HeaderNav,
    FooterNav,
  },
  data: function () {
    return {
      formEmail: '',
      formPassword: ''
    }
  },
  created () {
    this.getAllCategories();
  },
  methods: {
    async signIn () {
      var params = new URLSearchParams()
      params.append('email', this.formEmail)
      params.append('password', this.formPassword)
      await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/login', params)
      .then(res => {
        this.$cookies.config('1d', '', '', true);
        this.$cookies.set('uuid', res.data.uuid);
        alert('ログインしました。')
        this.$router.push({name: 'WebsiteList'})
      })
      .catch(error => {
        alert(error)
      });
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
  }
}
</script>

<style scoped lang="scss" src="../assets/css/signIn.scss">
</style>
