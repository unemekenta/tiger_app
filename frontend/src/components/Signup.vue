<template lang="pug">
  .signup-page
    header-nav(:allCategories="allCategories")
    h2 Sign up
    form.signup-page-form
      input.signup-page-form-name(type="text" placeholder="Name" v-model="formName")
      input.signup-page-form-email(type="email" placeholder="Email" v-model="formEmail")
      input.signup-page-form-password(type="password" placeholder="Password" v-model="formPassword")
      .signup-page-form-agree
        p 登録には
          router-link(to="/terms_of_service") 利用規約
          | 、
          router-link(to="/privacy_policy") プライバシーポリシー
          | の同意が必要になります。
        .signup-page-form-agree-btn
          input#checkbox1(type="checkbox" v-model="agree")
          label(for="checkbox1") 利用規約に同意する
      button.signup-page-form-button(@click="signUp") 登録

    .signup-page-to-signin
      p アカウントをすでにお持ちの方はこちら
      router-link(to="/signin") ログインする

    footer-nav

</template>

<script>
import axios from 'axios'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'Signup',
  components: {
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      formName: '',
      formEmail: '',
      formPassword: '',
      agree: false,
      allCategories: [],
    }
  },
  created () {
    this.getAllCategories();
  },
  methods: {
    async signUp () {
      var params = new URLSearchParams()
      params.append('name', this.formName)
      params.append('email', this.formEmail)
      params.append('password', this.formPassword)
      if (this.agree && this.formName && this.formEmail && this.formPassword) {
        await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/signup', params)
          .then((response) => {
            alert('アカウントを作成しました.', response)
            this.$router.push('/')
          })
          .catch(error => {
            alert('アカウントを作成できませんでした。emailがすでに登録されている可能性があります。', error)
          });
      } else {
        alert('未記入の項目があります')
      }
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

<style scoped lang="scss" src="../assets/css/signUp.scss">
</style>
