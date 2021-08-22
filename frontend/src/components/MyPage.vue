<template lang="pug">
  .top
    .header
      p マイページ
    .main
      router-link( to="/" ) HOME
      button(@click="signOut") ログアウト
</template>

<script>

export default {
  name: 'MyPage',
  data () {
    return {
    }
  },
  created () {
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt')
      this.$router.push('/signin')
    },
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/myPage.scss";
</style>
