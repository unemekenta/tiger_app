<template lang="pug">
.menu02
  .menu02-content
    .menu02-content-head
      | メニュー
    .menu02-content-items(v-if="signedIn")
      search-box
      .menu02-content-items-item
        router-link( to="/" )
          p
            fa-icon(icon='home')
            | HOME
      .menu02-content-items-item
        router-link( to="/website_list" )
          p
            fa-icon(icon='list')
            | メディア一覧
      .menu02-content-items-item
        router-link( to="/mypage" )
          p
            fa-icon(icon='user')
            | マイページ
      .menu02-content-items-item(@click="$_loginMixin_signOut")
        p
          fa-icon(icon='sign-out-alt')
          | ログアウト

    .menu02-content-items(v-else)
      search-box
      .menu02-content-items-item
        router-link( to="/" )
          p
            fa-icon(icon='home')
            | HOME
      .menu02-content-items-item
        router-link( to="/website_list" )
          p
            fa-icon(icon='list')
            | メディア一覧
      .menu02-content-items-item
        router-link( to="/signin" )
          p
            fa-icon(icon='sign-in-alt')
            | ログイン

</template>

<script>
import { loginMixin } from '../mixin/loginable'
import SearchBox from '/src/atoms/SearchBox.vue'

export default {
  name: 'Menu02',
  mixins: [loginMixin],
  components: {
    SearchBox,
  },
  data () {
    return {
    }
  },
  created () {
  },
  methods: {
  },
}
</script>

<style scoped lang="scss" src="../assets/css/menu02.scss">
</style>

