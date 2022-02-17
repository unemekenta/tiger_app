<template lang="pug">
.allWrapper
  .l-container#container
    .top-fv
      img(src="../assets/images/top-fv.png")
      .top-fv-txt
        h1 Web Media Database
        .top-fv-button
          router-link( to="/website_list" ontouchstart="")
            p Web Media Database
        .top-fv-button
          router-link( to="/my_asset" ontouchstart="")
            p あといくら？
    .top-list
      //- .top-list-features
      //-   h3 新着
      //-   card-slider(:itemList="this.cardItems")
      //- .top-list-categories
      //-   h3 カテゴリから探す
      //-   tag-slider(:itemList="this.allCategories")
  footer-nav


</template>

<script>
import CardSlider from '/src/organisms/CardSlider.vue'
import TagSlider from '/src/organisms/TagSlider.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'TopList',
  components: {
    CardSlider,
    TagSlider,
    FooterNav,
  },
  data () {
    return {
      // items準備
      cardItems:['aa', 'bb', 'cc'],
      // allCategories: [],
    }
  },
  created () {
    // this.getAllCategories();
  },
  methods: {
    // async getAllCategories () {
    //   await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/categories')
    //   .then(res => {
    //     this.allCategories = res.data;
    //     return;
    //   })
    //   .catch(error => {
    //     console.log(error);
    //     return;
    //   });
    // },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/topList.scss">
</style>
