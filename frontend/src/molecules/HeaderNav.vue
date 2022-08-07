<template lang="pug">
.header-nav.bg-white.shadow
  nav.header-nav-top
    .container.px-6.mx-auto
      .flex.items-center.justify-between
        h1.text-xl.font-semibold.text-gray-700
          a.text-2xl.font-bold.text-gray-800.transition-colors.duration-200.transform(class="hover:text-gray-700" href="/")
            img(src="@/assets/images/logo.png")

        .flex
          button.text-gray-500(@click="onClick()" class="hover:text-gray-600" type="button", aria-label="toggle menu")
            svg.w-6.h-6.fill-current(viewBox="0 0 24 24")
              path(fill-rule="evenodd", d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z")

  #headerNavMenu.header-nav-item.is-hidden.bg-white
    .flex-1(v-if="iconContent")
      ul.flex.flex-col.-mx-4
        li.px-2.py-1.mx-2.mt-2.font-bold.transform
          router-link.px-2.py-1.w-full.block.rounded-md(class="hover:bg-gray-300" to="/")
            p.inline-flex.items-center
              fa-icon.mr-3(icon="home")
              | HOME
        li.px-2.py-1.mx-2.mt-2.font-bold.transform
          router-link.px-2.py-1.w-full.block.rounded-md(class="hover:bg-gray-300" to="/website_list")
            p.inline-flex.items-center
              fa-icon.mr-3(icon="list")
              | メディア一覧
        li.px-2.py-1.mx-2.mt-2.font-bold.transform
          router-link.px-2.py-1.w-full.block.rounded-md(class="hover:bg-gray-300" to="/my_asset")
            p.inline-flex.items-center
              fa-icon.mr-3(icon="coins")
              | あといくら？
        li.px-2.py-1.mx-2.mt-2.font-bold.transform(v-if="signedIn")
          router-link.px-2.py-1.w-full.block.rounded-md(class="hover:bg-gray-300" to="/mypage")
            p.inline-flex.items-center
              fa-icon.mr-3(icon="user")
              | マイページ
        li.px-2.py-1.mx-2.mt-2.font-bold.transform( v-if="signedIn")
          p.px-2.py-1.w-full.block.inline-flex.items-center.rounded-md.cursor-pointer(class="hover:bg-gray-300" @click="$_loginMixin_signOut")
            fa-icon.mr-3(icon="sign-out-alt")
            | ログアウト
        li.inline-flex.items-center.py-1.mx-2.mt-2.font-bold.transform.rounded-md(class="hover:bg-gray-300" v-if="!signedIn")
          router-link.px-2.py-1.w-full.block(to="/signin")
            p.inline-flex.items-center
              fa-icon.mr-3(icon="sign-in-alt")
              | ログイン
        px-2.py-1.mx-2.mt-2.font-bold.transform.rounded-md
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
