<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p DETAIL
      .main
        menu01(:allCategories="allCategories")
        website-detail-item(:detailObject="websiteDetail")
        menu02

</template>

<script>
import axios from 'axios'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import WebsiteDetailItem from '/src/organisms/WebsiteDetailItem.vue'

export default {
  name: 'WebsiteDetail',
  components: {
    Menu01,
    Menu02,
    WebsiteDetailItem,
  },
  data () {
    return {
      websiteDetail: {},
      allCategories: [],
    }
  },
  created () {
    this.getWebsiteDetail();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async getWebsiteDetail () {
      let resWebsiteDetail = await axios.get('http://localhost:8000/api/websites/' + this.$route.params.id);
      this.websiteDetail = resWebsiteDetail.data;
    },
    async getAllCategories () {
      let resAllCategories = await axios.get('http://localhost:8000/api/categories');
      this.allCategories = resAllCategories.data;
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/websiteDetail.scss">
</style>
