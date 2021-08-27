<template lang="pug">
.top
    .header
      p マイページ
      button.header-logoutBtn(@click="signOut") ログアウト
    .user
      //- img.user-icon(:src="userData.ImageURL")
      .user-name {{userName}} さんのページ
    .main
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
      .main-menu02
        .menu-menu02-content
          .menu-menu02-content-head
            | メニュー
          .menu-menu02-content-items
            .menu-menu02-content-items-item
              router-link( to="/" ) HOME
            .menu-menu02-content-items-item
              router-link(to="/edit").header-editBtn 個人情報編集
            .menu-menu02-content-items-item(@click="signOut") ログアウト

</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'
import axios from 'axios'

export default {
  name: 'MyPage',
  data () {
    return {
      userName: '',
    }
  },
  created () {
    this.getUser();
  },
  methods: {
    signOut () {
      window.$cookies.remove('jwt');
      this.$router.push('/signin');
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
<style scoped>
  @import "../assets/css/myPage.scss";
</style>
