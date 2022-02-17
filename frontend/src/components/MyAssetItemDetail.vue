<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .main
        menu01-item(:allCategories="allCategories").is-only-pc
        .main-myasset
          h2.main-myasset-title {{userName}} さんのページ
          .main-myasset-detail-contents
            .main-myasset-contents-form
              form
                .main-myasset-contents-form-component
                  label(for="year")
                    p 年
                  select(name="year" v-model="itemData.year")
                    option(v-for="option in formYearOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="month")
                    p 月
                  select(name="month" v-model="itemData.month")
                    option(v-for="option in formMonthOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="title")
                    p 項目名
                  input(name="title" type="text" v-model="itemData.title")
                .main-myasset-contents-form-component
                  label(for="label")
                    p 増減
                  select(name="label" v-model="itemData.moneyAccountLabelId")
                    option(v-for="option in formMoneyAccountLabelOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="amount")
                    p 金額(円)
                  input(name="amount" type="number" min="0" v-model="itemData.amount")
                .main-myasset-contents-form-component
                  label(for="contents")
                    p メモ
                  input(name="contents" type="text" v-model="itemData.contents")
                //- button(@click.prevent="putMoneyAccount()")
                //-   p 更新
                button(@click.prevent="deleteMoneyAccount()")
                  p 削除

        menu02-item.is-only-pc
  footer-nav

</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import Menu01Item from '../organisms/Menu01.vue'
import Menu02Item from '../organisms/Menu02.vue'
import HeaderNav from '../molecules/HeaderNav.vue'
import FooterNav from '../molecules/FooterNav.vue'
import MoneyList from '../molecules/MoneyList.vue'

export default {
  name: 'MyAssetItemDetail',
  components: {
    Menu01Item,
    Menu02Item,
    HeaderNav,
    FooterNav,
    MoneyList,
  },
  data () {
    return {
      userName: '',
      userID: 0,
      itemData: '',
      year: 2022,
      month: 2,
      formMonthOptions: [
        { text: '1月', value: '1' },
        { text: '2月', value: '2' },
        { text: '3月', value: '3' },
        { text: '4月', value: '4' },
        { text: '5月', value: '5' },
        { text: '6月', value: '6' },
        { text: '7月', value: '7' },
        { text: '8月', value: '8' },
        { text: '9月', value: '9' },
        { text: '10月', value: '10' },
        { text: '11月', value: '11' },
        { text: '12月', value: '12' },
      ],
      formYearOptions: [
        { text: '2018年', value: '2018' },
        { text: '2019年', value: '2019' },
        { text: '2020年', value: '2020' },
        { text: '2021年', value: '2021' },
        { text: '2022年', value: '2022' },
      ],
      allCategories: [],
      jwtUserData: '',
      formMoneyAccountLabelOptions: [
        { text: '収入', value: '1' },
        { text: '支出', value: '2' },
      ],
    }
  },
  async created() {
    this.getAllCategories();
    await this.getUser();
    this.getItemData();
  },
  methods: {
    async getItemData() {
       await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user/detail/' + this.$route.params.id,
      {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(res => {
         this.itemData = res.data;
         return;
      })
      .catch(error => {
        console.log(error);
        return;
      });
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
    async getUser () {
      if (window.$cookies.isKey('uuid')) {
        console.log(window.$cookies.get('uuid'));
        await axios
          .post(
            process.env.VUE_APP_API_BASE_URL + "/api/login_check",
            {
              uuid: window.$cookies.get('uuid')
            },
          )
          .then((res) => {
            this.jwtUserData = res.data;
          })
          .catch((error) => {
            alert('ログインしてください');
            return error;
          });

        let user_id = VueJwtDecode.decode(this.jwtUserData).id;
        await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/user/'+ user_id,
        {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
          .then((res) => {
            this.userName = res.data.name;
            this.userID = res.data.id;
            return;
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
    async deleteMoneyAccount() {
      await axios.delete(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user/' + this.itemData.id, {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(() => {
        alert( this.itemData.title + 'を削除しました。');
        this.$router.push({name: 'MyAsset'})
      })
      .catch(error => {
        alert(error)
      });
    },
    // async postMoneyAccount() {
    //   let params = new URLSearchParams()
    //   params.append('userId', this.userID)
    //   params.append('moneyAccountLabelId', this.formMoneyAccountLabelId)
    //   params.append('amount', this.formAmount)
    //   params.append('title', this.formTitle)
    //   params.append('contents', this.formContents)
    //   params.append('year', this.formYear)
    //   params.append('month', this.formMonth)
    //   await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user', params, {
    //     headers: {'Authorization': 'Bearer ' + this.jwtUserData}
    //   })
    //   .then(res => {
    //     alert( this.formTitle + 'を登録しました。');
    //     this.getMoneyAccount(this.userID);
    //     this.formVisibleFlg = false;
    //   })
    //   .catch(error => {
    //     alert(error)
    //   });
    // },
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/components/myAsset.scss">
</style>
