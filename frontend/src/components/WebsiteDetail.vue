<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .header
        p
          fa-icon(icon="info-circle")
          | DETAIL
      .main
        menu01-item(:allCategories="allCategories").is-only-pc
        website-detail-item(:detailObject="websiteDetail" :detailObjectContent="websiteDetailContent")
        menu02-item.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import Menu01Item from '/src/organisms/Menu01.vue'
import Menu02Item from '/src/organisms/Menu02.vue'
import WebsiteDetailItem from '/src/organisms/WebsiteDetailItem.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'WebsiteDetail',
  components: {
    Menu01Item,
    Menu02Item,
    WebsiteDetailItem,
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      websiteDetail: {},
      websiteDetailContent: {},
      allCategories: [],
    }
  },
  created () {
    this.getWebsiteDetail();
    this.getWebsiteDetailContent();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async getWebsiteDetail () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/website/' + this.$route.params.id)
      .then(res => {
        this.websiteDetail = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    async getWebsiteDetailContent () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/website_content/' + this.$route.params.id)
      .then(res => {
        this.websiteDetailContent = res.data;
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

<style scoped lang="scss" src="../assets/css/websiteDetail.scss">
</style>
