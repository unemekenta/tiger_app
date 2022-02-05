<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .header
        p HOME
      .main(v-if="allWebsites")
        menu01-item(:allCategories="allCategories").is-only-pc
        website-list-item(:allLists="allWebsites")
        menu02-item.is-only-pc
      .main(v-else)
        menu01-item(:allCategories="allCategories").is-only-pc
        website-list-item(:allLists="[]")
        menu02-item.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import Menu01Item from '/src/organisms/Menu01.vue'
import Menu02Item from '/src/organisms/Menu02.vue'
import WebsiteListItem from '/src/organisms/WebsiteListItem.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'WebsiteList',
  components: {
    Menu01Item,
    Menu02Item,
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
  watch: {
    '$route' () {
      this.getAllWebsites();
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
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/websites_category/' + this.$route.params.category_id)
      .then(res => {
        this.allWebsites = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
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
  },
}
</script>

<style scoped lang="scss" src="../assets/css/websiteList.scss">
</style>
