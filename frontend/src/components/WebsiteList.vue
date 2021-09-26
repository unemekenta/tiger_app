<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container
    .top
      .image
        img(src="../assets/images/list-top-fv.png")
      .header
        p Top List
      .main
        menu01(:allCategories="allCategories").is-only-pc
        website-list-item(:allLists="allWebsites")
        menu02.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import WebsiteListItem from '/src/organisms/WebsiteListItem.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'WebsiteList',
  components: {
    Menu01,
    Menu02,
    WebsiteListItem,
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      allWebsites: [],
      allCategories: [],
    }
  },
  created () {
    this.getAllWebsites();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async getAllWebsites () {
      await axios.get('http://localhost:8000/api/websites')
      .then(res => {
        this.allWebsites = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    async getAllCategories () {
      await axios.get('http://localhost:8000/api/categories')
      .then(res => {
        this.allCategories = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/websiteList.scss">
</style>
