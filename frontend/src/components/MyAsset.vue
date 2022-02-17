<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .main
        .main-myasset
          h2.main-myasset-title {{userName}} さんのページ
          .main-myasset-contents

            chart-pie(
              :chartData="chartData"
              :options="chartOptions"
              :style='"width: 70%; margin: 2rem auto;"'
            )

            .main-myasset-contents-top
              h2 {{year}} 年 {{month}} 月の収支
              .main-myasset-contents-top-income
                .main-myasset-contents-top-income-label
                  p 収入
                .main-myasset-contents-top-income-amount
                  p {{ sumIncome }} 
                    span.main-myasset-contents-top-income-yen
                      | 円
              .main-myasset-contents-top-expenses
                .main-myasset-contents-top-expenses-label
                  p 支出
                .main-myasset-contents-top-expenses-amount
                  p {{ sumExpenses }} 
                    span.main-myasset-contents-top-expenses-yen
                      | 円
              .main-myasset-contents-top-sum
                .main-myasset-contents-top-sum-label
                  p 残金
                .main-myasset-contents-top-sum-amount
                  p {{ sumIncome - sumExpenses }}
                    span.main-myasset-contents-top-sum-yen
                      | 円

            .main-myasset-contents-details
              h3 {{year}} 年 {{month}} 月の収支内訳
              .main-myasset-contents-details-item
                .main-myasset-contents-details-item-inner
                  h4 収入
                  money-list(:arr="income")
                .main-myasset-contents-details-item-inner
                  h4 支出
                  money-list(:arr="expenses")
            .main-myasset-contents-form
              form(v-if="formVisibleFlg")
                .main-myasset-contents-form-component
                  label(for="year")
                    p 年
                  select(name="year" v-model="formYear")
                    option(v-for="option in formYearOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="month")
                    p 月
                  select(name="month" v-model="formMonth")
                    option(v-for="option in formMonthOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="title")
                    p 項目名
                  input(name="title" type="text" v-model="formTitle" placeholder="食費")
                .main-myasset-contents-form-component
                  label(for="label")
                    p 増減
                  select(name="label" v-model="formMoneyAccountLabelId")
                    option(v-for="option in formMoneyAccountLabelOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="amount")
                    p 金額(円)
                  input(name="amount" type="number" min="0" v-model="formAmount" placeholder="30000")
                .main-myasset-contents-form-component
                  label(for="contents")
                    p メモ
                  input(name="contents" type="text" v-model="formContents")
                button(@click.prevent="postMoneyAccount()")
                  p 保存
              button(v-if="formVisibleFlg" @click="changeFormVisibleFlg()")
                p -
              button(v-else @click="changeFormVisibleFlg()")
                p +
  footer-nav

</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import HeaderNav from '../molecules/HeaderNav.vue'
import FooterNav from '../molecules/FooterNav.vue'
import ChartPie from '../organisms/ChartPie.vue';
import MoneyList from '../molecules/MoneyList.vue'

export default {
  name: 'MyAsset',
  components: {
    HeaderNav,
    FooterNav,
    ChartPie,
    MoneyList,
  },
  data () {
    return {
      userName: '',
      userID: 0,
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
      jwtUserData: '',
      income: [],
      expenses: [],
      sumIncome: 0,
      sumExpenses: 0,
      chartData: null,
      chartOptions: {
        responsive: true
      },
      formVisibleFlg: false,
      formMoneyAccountLabelOptions: [
        { text: '収入', value: '1' },
        { text: '支出', value: '2' },
      ],
    }
  },
  async created() {
    await this.getUser();
    this.sumAmount(this.income, this.sumIncome);
  },
  methods: {
    async getMoneyAccount (userID) {
      this.resetMoneyAccount();
      this.resetChartData();
      this.resetSumData();
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user/' + userID+ '/' + this.year+ '/' + this.month,
      {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(res => {
        const dataObj = res.data

        this.chartData = {
          labels: [],
          datasets: [
            {
              data: []
            },
          ],
          options: {
            responsive: true,
            maintainAspectRatio: true,
            plugins: {
              colorschemes: {
                scheme: 'tableau.Tableau20',
              },
            },
          },
        }

        dataObj.forEach (element => {
          this.insertChartData(element);
          this.classifyMoneyAccount(element);
        })
        return;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    insertChartData(data) {
      this.chartData.labels.push(data.title);
      this.chartData.datasets[0].data.push(data.amount);
    },
    resetMoneyAccount() {
      this.income = [];
      this.expenses = [];
    },
    resetChartData() {
      this.chartData = {};
    },
    resetSumData() {
      this.sumIncome = 0;
      this.sumExpenses = 0;
    },
    classifyMoneyAccount(data) {
      switch (data.moneyAccountLabelId) {
        case 1:
          this.income.push(data);
          this.sumIncome = this.sumIncome + data.amount
          break;
        case 2:
          this.expenses.push(data);
          this.sumExpenses = this.sumExpenses + data.amount
          break;
        default:
          break;
      }
    },

    async getAllCategories () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/categories')
      .then(res => {
        this.allCategories = res.data;
        return;
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
            return;
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
            this.getMoneyAccount(this.userID);
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
    changeFormVisibleFlg() {
      this.formVisibleFlg =! this.formVisibleFlg;
    },
    async postMoneyAccount() {
      let params = new URLSearchParams()
      params.append('userId', this.userID)
      params.append('moneyAccountLabelId', this.formMoneyAccountLabelId)
      params.append('amount', this.formAmount)
      params.append('title', this.formTitle)
      params.append('contents', this.formContents)
      params.append('year', this.formYear)
      params.append('month', this.formMonth)
      await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user', params, {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(() => {
        alert( this.formTitle + 'を登録しました。');
        this.getMoneyAccount(this.userID);
        this.formVisibleFlg = false;
        return;
      })
      .catch(error => {
        alert(error);
        return;
      });
    },
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/components/myAsset.scss">
</style>
