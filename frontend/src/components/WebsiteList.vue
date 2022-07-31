<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .image
        img(src="@/assets/images/list-top-fv.png")
      .header
        p
          fa-icon(icon="desktop")
          | Top List
      .main
        menu01-item(:allCategories="allCategories").is-only-pc
        website-list-item(:allLists="allWebsites")
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
      serchQuery: this.$route.query.q,
    }
  },
  created () {
    this.getWebsites();
    this.getAllCategories();
  },
  async beforeRouteUpdate (to, from, next) {
    // URL の クエリが変わったときに 関数を実行してデータを更新する
    next();
    this.serchQuery = this.$route.query.q;
    this.getWebsites()
  },
  methods: {
    signOut () {
      window.$cookies.remove('uuid');
      this.$router.push('/signin');
    },
    async getWebsites () {
      if (this.serchQuery == null) {
        this.getAllWebsites();
      } else {
        this.getSearchedWebsites();
      }
    },
    async getAllWebsites() {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/websites')
        .then(res => {
          this.allWebsites = res.data;
          return;
        })
        .catch(error => {
          console.log(error);
          return;
        });
    },
    async getSearchedWebsites() {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/websites/search?q=' + this.serchQuery)
        .then(res => {
          this.allWebsites = res.data;
          return;
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
        return;
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
