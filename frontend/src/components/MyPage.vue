<template lang="pug">
.top
  .user
    img.user-icon(src="../../src/assets/images/logo.png")
    .user-name {{userName}} さんのページ
  .main
    menu01(:allCategories="allCategories")
    .main-myMessage
      h2.main-myMessage-title マイコメント
      //- .main-myMessage-contents(v-for="(myMessage, key) in this.myMessages" :key="key")
      //-   .main-myMessage-contents-top
      //-     fa-icon.main-myMessage-contents-top-deleteBtn(@click="deleteMessage(myMessage.ID)")(icon="times-circle")
      //-   p.main-myMessage-contents-txt {{myMessage.Content}}
      //-   .main-myMessage-contents-category
      //-     li カテゴリ
      //-     li カテゴリ
      //-   p.main-myMessage-contents-time {{createDate(myMessage.CreatedAt)}}
    .menu02
      .menu02-content
        .menu02-content-head
          | メニュー
        .menu02-content-items
          .menu02-content-items-item
            router-link( to="/" ) HOME
          .menu02-content-items-item
            router-link(to="/edit").header-editBtn 個人情報編集
          .menu02-content-items-item(@click="signOut") ログアウト

</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
import Menu01 from '/src/organisms/Menu01.vue'
import axios from 'axios'

export default {
  name: 'MyPage',
  components: {
    Menu01,
  },
  data () {
    return {
      userName: '',
      allCategories: [],
    }
  },
  created () {
    this.getUser();
    this.getAllCategories();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
    },
    async getAllCategories () {
      let resAllCategories = await axios.get('http://localhost:8000/api/categories');
      this.allCategories = resAllCategories.data;
    },
    async getUser () {
      let user_id = VueJwtDecode.decode(window.$cookies.get('jwt')).id;
      console.log(window.$cookies.get('jwt'));
      await axios.get('http://localhost:8000/auth/api/user/'+ user_id,
      {
        headers: {'Authorization': 'Bearer ' + window.$cookies.get('jwt')}
      })
        .then((res) => {
          this.userName = res.data.name;
          console.log(res);
        })
        .catch(error => {
          alert(error)
        });
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/myPage.scss">
</style>
