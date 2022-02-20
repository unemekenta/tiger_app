<template lang="pug">
.header-nav
  .header-nav-top
    h1
      a(href="/")
        img(src="../../src/assets/images/logo.png")
    .header-nav-pc.is-only-pc
      .header-nav-item
        router-link(to="/")
          p
            fa-icon(icon="home")
            | HOME
      .header-nav-item
        router-link(to="/website_list")
          p
            fa-icon(icon="list")
            | メディア一覧
      .header-nav-item
        router-link(to="/my_asset")
          p
            fa-icon(icon="coins")
            | あといくら？
      .header-nav-item(v-if="signedIn")
        router-link(to="/mypage")
          p
            fa-icon(icon="user")
            | マイページ
      .header-nav-item(v-if="signedIn")
        p.header-nav-item-title(@click="$_loginMixin_signOut")
          fa-icon(icon="sign-out-alt")
          | ログアウト
      .header-nav-item(v-if="!signedIn")
        router-link(to="/signin")
          p
            fa-icon(icon="sign-in-alt")
            | ログイン

    .header-nav-button.is-only-sp(@click="onClick()")
      fa-icon(v-if="iconContent", icon="times")
      fa-icon(v-else, icon="bars")

  #headerNavMenu.header-nav.is-hidden
    .header-nav-inner
      search-box
      .header-nav-item
        router-link(to="/")
          p
            fa-icon(icon="home")
            | HOME
          fa-icon(icon="chevron-right")
      .header-nav-item
        router-link(to="/website_list")
          p
            fa-icon(icon="list")
            | メディア一覧
          fa-icon(icon="chevron-right")
      .header-nav-item
        router-link(to="/my_asset")
          p
            fa-icon(icon="coins")
            | あといくら？
          fa-icon(icon="chevron-right")
      .header-nav-item(v-if="signedIn")
        router-link(to="/mypage")
          p
            fa-icon(icon="user")
            | マイページ
          fa-icon(icon="chevron-right")
      .header-nav-item(v-if="signedIn")
        p.header-nav-item-title(@click="$_loginMixin_signOut")
          fa-icon(icon="sign-out-alt")
          | ログアウト
      .header-nav-item(v-if="!signedIn")
        router-link(to="/signin")
          p
            fa-icon(icon="sign-in-alt")
            | ログイン
      .header-nav-item
        p カテゴリから探す
        .header-nav-item-categories(
          v-for="(allCatrgory, key) in this.allCategories",
          :key="key"
        )
          li(@click="reset(allCatrgory.id)")
            | {{ allCatrgory.name }}
          sub-category(:ancestorID="allCatrgory.id")

</template>

<script>
import { loginMixin } from '../mixin/loginable'
import SubCategory from "/src/atoms/SubCategory.vue";
import SearchBox from "/src/atoms/SearchBox.vue";

export default {
  name: "HeaderNav",
  mixins: [loginMixin],
  components: {
    SubCategory,
    SearchBox,
  },
  data() {
    return {
      iconContent: false,
    };
  },
  props: {
    allCategories: Array,
  },
  created() {
  },
  methods: {
    toggleView(iconContent, menu, background) {
      if(iconContent) {
        menu.classList.remove("is-hidden");
        background.classList.add("is-bgBlack");
      } else {
        menu.classList.add("is-hidden");
        background.classList.remove("is-bgBlack");
      }
    },
    onClick() {
      const menu = document.getElementById("headerNavMenu");
      const background = document.getElementById("container");
      if (!this.iconContent) {
        this.iconContent = true;
        this.toggleView(this.iconContent, menu, background);
      }
      else {
        this.iconContent = false;
        this.toggleView(this.iconContent, menu, background);
      }
    },
  },
};
</script>

<style scoped lang="scss" src="../assets/css/headerNav.scss">
</style>
