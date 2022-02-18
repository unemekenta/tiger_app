<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .main
        .main-myasset
          h2.main-myasset-title {{userName}} さん
          .main-myasset-contents

            chart-pie(
              :chartData="chartData"
              :options="chartOptions"
              :style='"width: 70%; margin: 2rem auto;"'
            )

            .main-myasset-contents-top
              h2
                button.main-myasset-contents-top-icon(@click="onClickBefore()")
                  fa-icon(icon='angle-left')
                | {{targetYear}} 年 {{targetMonth}} 月
                button.main-myasset-contents-top-icon(@click="onClickAfter()")
                  fa-icon(icon='angle-right')
              .main-myasset-contents-top-income
                .main-myasset-contents-top-income-label
                  p 収入
                .main-myasset-contents-top-income-amount
                  p {{ $_commify(sumIncome) }}
                    span.main-myasset-contents-top-income-yen
                      |  円
              .main-myasset-contents-top-expenses
                .main-myasset-contents-top-expenses-label
                  p 支出
                .main-myasset-contents-top-expenses-amount
                  p {{ $_commify(sumExpenses) }}
                    span.main-myasset-contents-top-expenses-yen
                      |  円
              .main-myasset-contents-top-sum
                .main-myasset-contents-top-sum-label
                  p 残金
                .main-myasset-contents-top-sum-amount
                  p {{ $_commify(sumIncome - sumExpenses) }}
                    span.main-myasset-contents-top-sum-yen
                      |  円

            .main-myasset-contents-details
              h3 収支内訳
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
                    option(v-for="y in 50" :value="y + 2020")
                      | {{ y + 2020 }}年
                .main-myasset-contents-form-component
                  label(for="month")
                    p 月
                  select(name="month" v-model="formMonth")
                    option(v-for="m in 12" :value="m")
                      | {{ m }}月
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
                  textarea(name="contents", rows="5", v-model="formContents")
                button(@click.prevent="postMoneyAccount()")
                  p 保存
              button(v-if="formVisibleFlg" @click="changeFormVisibleFlg()")
                p -
              button(v-else @click="changeFormVisibleFlg()")
                p +
  footer-nav

</template>

<script>
import { priceMixin } from '../mixin/price'
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import HeaderNav from '../molecules/HeaderNav.vue'
import FooterNav from '../molecules/FooterNav.vue'
import ChartPie from '../organisms/ChartPie.vue';
import MoneyList from '../molecules/MoneyList.vue'

export default {
  name: 'MyAsset',
  mixins: [priceMixin],
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
      targetYear: 2020,
      targetMonth: 1,
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
    this.getUser();
    this.getDate();
    await this.sumAmount(this.income, this.sumIncome);
  },
  methods: {
    onClickBefore() {
      if (this.targetMonth !== 1) {
        this.targetMonth = this.targetMonth -1;
        this.getMoneyAccount(this.userID);
      }
    },
    onClickAfter() {
      if (this.targetMonth !== 12) {
        this.targetMonth = this.targetMonth +1;
        this.getMoneyAccount(this.userID);
      }
    },
    getDate() {
      const date = new Date();
      this.targetYear = date.getFullYear();
      this.targetMonth = date.getMonth()+1;
    },
    async getMoneyAccount (userID) {
      this.resetMoneyAccount();
      this.resetChartData();
      this.resetSumData();
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/user/' + userID+ '/' + this.targetYear+ '/' + this.targetMonth,
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
      if (!this.formContents) {
        this.formContents = "";
      }
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
