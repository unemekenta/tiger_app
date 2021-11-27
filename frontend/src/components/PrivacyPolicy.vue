<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container
    .top
      .header
        p 規約
      .main
        menu01(:allCategories="allCategories").is-only-pc
        privacy-policy-contents
        menu02.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'
import PrivacyPolicyContents from '/src/organisms/PrivacyPolicyContents.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue'
import FooterNav from '/src/molecules/FooterNav.vue'

export default {
  name: 'PrivacyPolicy',
  components: {
    Menu01,
    Menu02,
    PrivacyPolicyContents,
    HeaderNav,
    FooterNav,
  },
  data () {
    return {
      allCategories: [],
    }
  },
  created () {
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
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

<style scoped lang="scss" src="../assets/css/pageNomal.scss">
</style>
