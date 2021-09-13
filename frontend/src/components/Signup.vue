<template lang="pug">
  .signUpPage
    h2 Sign up
    form.signUpPage-form
      input.signUpPage-form-name(type="text" placeholder="Name" v-model="formName")
      input.signUpPage-form-email(type="email" placeholder="Email" v-model="formEmail")
      input.signUpPage-form-password(type="password" placeholder="Password" v-model="formPassword")
      button.signUpPage-form-button(@click="signUp") 登録
      p アカウントをすでにお持ちの方はこちら
      router-link(to="/signin") ログインする
</template>

<script>
import axios from 'axios'

export default {
  name: 'Signup',
  data () {
    return {
      formName: '',
      formEmail: '',
      formPassword: ''
    }
  },
  methods: {
    async signUp () {
      var params = new URLSearchParams()
      params.append('name', this.formName)
      params.append('email', this.formEmail)
      params.append('password', this.formPassword)
      await axios.post('http://localhost:8000/api/signup', params)
      .then((response) => {
        alert('アカウントを作成しました', response)
        this.$router.push('/')
      })
      .catch(error => {
        alert('アカウントを作成できませんでした', error)
      });
    }
  }
}
</script>

<style scoped lang="scss" src="../assets/css/signUp.scss">
</style>
