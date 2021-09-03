<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p HOME
      .main
        menu01(:allCategories="allCategories")

        .main-allMessage
            h2.main-allMessage-title 一覧
            .main-allMessage-contents(v-for="(term, key) in this.allTerms" :key="key")
              router-link( :to="{name:'Detail', params: {id:`${term.id}`}}").wrapper
                .main-allMessage-contents-userIcon
                  img(src="../../src/assets/images/logo.png")
                  p.main-allMessage-contents-userIcon-name
                    | 運営会社
                p.main-allMessage-contents-txt 
                  | {{term.name}}
                .main-allMessage-contents-category
                  li カテゴリ
                  li カテゴリ
                  li カテゴリ
                p.main-allMessage-contents-time 
                  | {{createDate(term.CreatedAt)}}

        menu02

</template>

<script>
import axios from 'axios'
import Menu01 from '/src/organisms/Menu01.vue'
import Menu02 from '/src/organisms/Menu02.vue'

export default {
  name: 'TopList',
  components: {
    Menu01,
    Menu02,
  },
  data () {
    return {
      allTerms: [],
      allCategories: [],
    }
  },
  created () {
    this.getAllTerms();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async getAllTerms () {
      let resAllTerms = await axios.get('http://localhost:8000/api/terms');
      this.allTerms = resAllTerms.data;
    },
    async getAllCategories () {
      let resAllCategories = await axios.get('http://localhost:8000/api/categories');
      this.allCategories = resAllCategories.data;
    },
    createDate (createdAt) {
      let formattedDate = createdAt.split(/T|:|-/);
      formattedDate.pop()
      let dateString = formattedDate[0] + '年' + formattedDate[1]+ '月' + formattedDate[2] + '日' + formattedDate[3] + '時' + formattedDate[4] + '分';
      return dateString;
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/topList.scss">
</style>
