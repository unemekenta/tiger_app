<template lang="pug">
.menu02
  .menu02-content
    .menu02-content-head
      | メニュー
    .menu02-content-items
      .menu02-content-items-item
        router-link( to="/" ) HOME
      .menu02-content-items-item
        router-link( to="/mypage" ) マイページ
      .menu02-content-items-item(@click="signOut") ログアウト

</template>

<script>
export default {
  name: 'Menu02',
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/menu02.scss">
</style>

