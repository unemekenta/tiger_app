<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .header
        p Book Logを作成する
      .main
        .w-full.flex.justify-center.my-2.mx-4
          form.w-full.max-w-xl.bg-white.rounded-lg.shadow-md.p-6
           .flex.flex-wrap.-mx-3.mb-6
            .w-full.px-3.mb-6
              label.block.uppercase.tracking-wide.text-gray-700.text-xs.font-bold.mb-2(for="title")
                | TITLE
              input#title.block.w-full.bg-white.text-gray-900.font-medium.border.border-gray-400.rounded-lg.py-3.px-3.leading-tight(type="text", required, v-model="form.name")
            .w-full.px-3.mb-6
              label.block.uppercase.tracking-wide.text-gray-700.text-xs.font-bold.mb-2(for="url")
                | URL
              input#url.block.w-full.text-gray-900.font-medium.border.border-gray-400.rounded-lg.py-3.px-3.leading-tight(type="text", required, v-model="form.url")
            .w-full.px-3.mb-6
              label.block.uppercase.tracking-wide.text-gray-700.text-xs.font-bold.mb-2(for="companyName")
                | COMPANY NAME
              input#companyName.block.w-full.text-gray-900.font-medium.border.border-gray-400.rounded-lg.py-3.px-3.leading-tight(type="text", required, v-model="form.companyName")
            .w-full.px-3.mb-6
              label.block.uppercase.tracking-wide.text-gray-700.text-xs.font-bold.mb-2(for="image")
                | IMAGE
              input#image.block.w-full.text-gray-900.font-medium.border.border-gray-400.rounded-lg.py-3.px-3.leading-tight(type="text", v-model="form.image")
            .w-full.px-3.mb-6
              button.block.w-full.bg-blue-600.text-gray-100.font-bold.border.border-gray-200.rounded-lg.py-3.px-3.leading-tight(@click="createWebsite")
                | SUBMIT

  footer-nav

</template>

<script lang="ts">
import Vue from 'vue';
import axios from 'axios';
import jwtDecode from 'jwt-decode';
// import Menu01Item from '/src/organisms/Menu01.vue'
// import Menu02Item from '/src/organisms/Menu02.vue'
// import WebsiteListItem from '/src/organisms/WebsiteListItem.vue'
import HeaderNav from '/src/molecules/HeaderNav.vue';
import FooterNav from '/src/molecules/FooterNav.vue';

export type DataType = {
  // allWebsites: any[];
  // allCategories: any[];
  jwtUserData: string;
  userName: string;
  userID: number;
  form: {
    name: string,
    url: string,
    companyName: string,
    image: string
  };
}

export type decodedJwtData = {
  id: number;
}

export default Vue.extend({
  name: 'BookLogNew',
  components: {
    HeaderNav,
    FooterNav,
  },
  data (): DataType {
    return {
      userName: '',
      userID: 0,
      jwtUserData: '',
      form: {
        name: '',
        url: '',
        companyName: '',
        image: '',
      },
    }
  },
  async created () {
    await this.getUser();
  },
  methods: {
    async getUser () {
      if (window.$cookies.isKey('uuid')) {
        await axios
          .post(
            process.env.VUE_APP_API_BASE_URL + "/api/login_check",
            {
              uuid: window.$cookies.get('uuid')
            },
          )
          .then((res) => {
            this.jwtUserData = res.data;
            return;
          })
          .catch((error) => {
            alert('ログインしてください');
            return error;
          });

        let decodedHeader: decodedJwtData = jwtDecode(this.jwtUserData);
        let user_id = decodedHeader.id;
        await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/user/'+ user_id,
        {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
          .then((res) => {
            this.userName = res.data.name;
            this.userID = res.data.id;
          })
          .catch((error) => {
            alert('権限がありません。');
            return error;
          });
      } else {
        alert('ログインしてください');
        return;
      }
    },
    async createWebsite () {
      await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/auth/book_log/new', this.form, {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(res => {
        alert('登録しました。' + res)
        this.$router.push({name: 'BookLogNew'})
        return;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
  },
});
</script>
<style scoped lang="scss" src="../assets/css/components/base.scss">
</style>