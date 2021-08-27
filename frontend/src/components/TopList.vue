<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p HOME
      .main
        .main-menu01
          .menu-menu01-content
            .menu-menu01-content-head
              | カテゴリ
            .menu-menu01-content-item(v-for="(allCatrgory, key) in this.allCategories" :key="key")
              li {{allCatrgory.name}}

        .main-allMessage
            h2.main-allMessage-title コメント一覧
            .main-allMessage-contents(v-for="(term, key) in this.allTerms" :key="key")
              .wrapper
                .main-allMessage-contents-userIcon
                  img(src="../../src/assets/images/logo.png")
                  p.main-allMessage-contents-userIcon-name
                    | username
                p.main-allMessage-contents-txt 
                  | {{term.name}}
                .main-allMessage-contents-category
                  li カテゴリ
                p.main-allMessage-contents-time 
                  | {{createDate(term.CreatedAt)}}
        .main-menu02
          .menu-menu02-content
            .menu-menu02-content-head
              | メニュー
            .menu-menu02-content-items
              .menu-menu02-content-items-item
                router-link( to="/" ) HOME
              .menu-menu02-content-items-item
                router-link( to="/mypage" ) マイページ
              .menu-menu02-content-items-item(@click="signOut") ログアウト

</template>

<script>
import axios from 'axios'

export default {
  name: 'TopList',
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

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/topList.scss";
</style>
