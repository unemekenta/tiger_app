<template lang="pug">
  .signInPage
    h2 Sign in
    .signInPage-form
      input.signInPage-form-email(type="email" placeholder="Email" v-model="formEmail")
      input.signInPage-form-password(type="password" placeholder="Password" v-model="formPassword")
      button.signInPage-form-button(@click="signIn") ログイン
      p アカウントをまだお持ちでない方はこちら
      router-link(to="/signup") アカウントを作成する
</template>

<script>
import axios from 'axios'

export default {
  name: 'Signin',
  data: function () {
    return {
      formEmail: '',
      formPassword: ''
    }
  },
  methods: {
    async signIn () {
      var params = new URLSearchParams()
      params.append('email', this.formEmail)
      params.append('password', this.formPassword)
      await axios.post('http://localhost:8000/api/login', params)
      .then((response) => {
        this.$cookies.config('1d', '', '', true);
        this.$cookies.set('jwt', response.data.token);
        console.log(this.$cookies.get('jwt'))
        alert('ログインしました。')
        this.$router.push('/')
      })
      .catch(error => {
        alert(error)
      });
    }
  }
}
</script>

<style scoped>
  @import "../assets/css/signIn.scss";
</style>
