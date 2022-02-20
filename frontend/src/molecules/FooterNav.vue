<template lang="pug">
.footer-nav
  .footer-nav-wrapper
    .footer-nav-top
      li
        router-link( to="/terms_of_service" )
          p 利用規約
      li
        router-link( to="/privacy_policy" )
          p プライバシーポリシー
    .footer-nav-btm
      p {{copyright}}

</template>

<script>

export default {
  name: 'FooterNav',
  data () {
    return {
      copyright: process.env.VUE_APP_COPYRIGHT,
    }
  }
}
</script>

<style scoped lang="scss" src="../assets/css/footerNav.scss">
</style>
