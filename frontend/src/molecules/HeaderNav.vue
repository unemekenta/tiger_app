<template lang="pug">
.header-nav
  .header-nav-top(v-if="signedIn")
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
        router-link(to="/mypage")
          p
            fa-icon(icon="user")
            | マイページ
      .header-nav-item
        p.header-nav-item-title(@click="$_loginMixin_signOut")
          fa-icon(icon="sign-out-alt")
          | ログアウト
    .header-nav-button.is-only-sp(@click="onClick()")
      fa-icon(v-if="iconContent", icon="times")
      fa-icon(v-else, icon="bars")

  .header-nav-top(v-else)
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
        router-link(to="/signin")
          p
            fa-icon(icon="sign-in-alt")
            | ログイン
    .header-nav-button.is-only-sp(@click="onClick()")
      fa-icon(v-if="iconContent", icon="times")
      fa-icon(v-else, icon="bars")

  #headerNavMenu.header-nav.is-hidden(v-if="signedIn")
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
        router-link(to="/mypage")
          p
            fa-icon(icon="user")
            | マイページ
          fa-icon(icon="chevron-right")
      .header-nav-item
        p.header-nav-item-title(@click="$_loginMixin_signOut")
          fa-icon(icon="sign-out-alt")
          | ログアウト
      .header-nav-item
        p カテゴリから探す
        .header-nav-item-categories(
          v-for="(allCatrgory, key) in this.allCategories",
          :key="key"
        )
          li(@click="reset(allCatrgory.id)")
            | {{ allCatrgory.name }}
          sub-category(:ancestorID="allCatrgory.id")
  #headerNavMenu.header-nav.is-hidden(v-else)
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
        router-link(to="/signin")
          p
            fa-icon(icon="sign-in-alt")
            | ログイン
          fa-icon(icon="chevron-right")
      .header-nav-item
        p カテゴリから探す
        .header-nav-item-categories(
          v-for="(allCatrgory, key) in this.allCategories",
          :key="key"
        )
          li
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
    onClick() {
      var manu = document.getElementById("headerNavMenu");
      manu.classList.toggle("is-hidden");
      var background = document.getElementById("container");
      background.classList.toggle("is-bgBlack");
      this.iconContent = !this.iconContent;
    },
  },
};
</script>

<style scoped lang="scss" src="../assets/css/headerNav.scss">
</style>
