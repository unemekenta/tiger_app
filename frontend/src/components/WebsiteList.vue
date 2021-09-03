<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p HOME
      .main
        menu01(:allCategories="allCategories")
        website-list-item(:allLists="allWebsites")
        menu02

</template>

<script>
import axios from 'axios'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import WebsiteListItem from '/src/organisms/WebsiteListItem.vue'

export default {
  name: 'WebsiteList',
  components: {
    Menu01,
    Menu02,
    WebsiteListItem,
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
      let resAllWebsites = await axios.get('http://localhost:8000/api/websites');
      this.allWebsites = resAllWebsites.data;
    },
    async getAllCategories () {
      let resAllCategories = await axios.get('http://localhost:8000/api/categories');
      this.allCategories = resAllCategories.data;
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/websiteList.scss">
</style>
