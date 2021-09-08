<template lang="pug">
.allWrapper
  .l-container
    .top-fv
      h1 Find your Media
      .top-fv-button
        router-link( to="/website_list" ) 一覧ページへ
    .top-list
      .top-list-features
        h3 新着
        card-slider(:itemList="this.cardItems")
      .top-list-categories
        h3 カテゴリから探す
        card-slider(:itemList="this.cardItems")


</template>

<script>
import CardSlider from '/src/organisms/CardSlider.vue'

export default {
  name: 'TopList',
  components: {
    CardSlider,
  },
  data () {
    return {
      // items準備
      cardItems:['aa', 'bb', 'cc'],
    }
  },
  created () {
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/topList.scss">
</style>
