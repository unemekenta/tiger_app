<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p HOME
      .main(v-if="allWebsites")
        menu01(:allCategories="allCategories")
        website-list-item(:allLists="allWebsites")
        menu02
      .main(v-else)
        menu01(:allCategories="allCategories")
        website-list-item(:allLists="[]")
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
      await axios.get('http://localhost:8000/api/websites_category/' + this.$route.params.category_id)
      .then(res => {
        console.log(res.data, typeof(res.data))
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
