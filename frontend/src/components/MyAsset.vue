<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .main
        .main-myasset
          h2.main-myasset-title あといくら？
          .main-myasset-contents
            .main-myasset-contents-data
              .main-myasset-contents-data-info(v-if="targetYear === thisYear && targetMonth === thisMonth")
                p {{userName}} さん
                p
                  | 今月の残り日数 
                  span.main-myasset-contents-data-info-bold {{ remainingDateOfThisMonth }}
                  | 日
                p
                  | 1日 
                  span.main-myasset-contents-data-info-bold {{ $_commify(remainingMoneyPerDay) }} 
                  | 円で生きろ！
              .main-myasset-contents-data-info(v-else)
                p {{userName}} さん
                p よくがんばりました。

              .main-myasset-contents-data-chart
                chart-pie(
                  :chartData="chartData"
                  :options="chartOptions"
                  :style='"width: 100%;"'
                )

                chart-horizontal-bar(
                  :chartData="chartBarData"
                  :options="chartBarOptions"
                  :style='"width: 100%;"'
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
                  p
                    span.main-myasset-contents-top-income-yen
                      | ¥
                    | {{ $_commify(sumIncome) }}
              .main-myasset-contents-top-expenses
                .main-myasset-contents-top-expenses-label
                  p 支出
                .main-myasset-contents-top-expenses-amount
                  p
                    span.main-myasset-contents-top-expenses-yen
                      | ¥
                    | {{ $_commify(sumExpenses) }}
              .main-myasset-contents-top-sum
                .main-myasset-contents-top-sum-label
                  p 残金
                .main-myasset-contents-top-sum-amount
                  p
                    span.main-myasset-contents-top-sum-yen
                      | ¥
                    | {{ $_commify(remainingMoney) }}

            .main-myasset-contents-details
              h3 収支内訳
              .main-myasset-contents-details-item
                .main-myasset-contents-details-item-inner.my-4.w-full
                  h4.my-2 収入
                  money-list(:arr="income")
                .main-myasset-contents-details-item-inner.my-4.w-full
                  h4.my-2 支出
                  money-list(:arr="expenses")

            .main-myasset-contents-form
              form(v-if="formVisibleFlg")
                .main-myasset-contents-form-component-flex
                  label(for="subscriptionsFlg")
                    input(type="checkbox" id="subscriptionsFlg" name="subscriptionsFlg" v-model="formSubscriptionsFlg")
                    | サブスク/固定費で登録する
                .main-myasset-contents-form-component
                  label(for="formDate")
                    | 登録月
                  date-input-form(@change-date="this.reflectDate" :name="formDateName")
                .main-myasset-contents-form-component(v-if="formSubscriptionsFlg")
                  label(for="formStartDate")
                    | 開始月
                  date-input-form(@change-date="this.reflectStartDate" :name="formStartDateName")
                  label(for="formEndDate")
                    | 終了月
                  date-input-form(@change-date="this.reflectEndDate" :name="formEndDateName")
                .main-myasset-contents-form-component
                  label(for="title")
                    | 項目名
                  input(name="title" type="text" v-model="formTitle" placeholder="食費")
                .main-myasset-contents-form-component
                  label(for="label")
                    | 増減
                  select(name="label" v-model="formMoneyAccountLabelId")
                    option(v-for="option in formMoneyAccountLabelOptions" :value="option.value")
                      | {{ option.text }}
                .main-myasset-contents-form-component
                  label(for="amount")
                    | 金額(円)
                  input(name="amount" type="number" min="0" v-model="formAmount" placeholder="30000")
                .main-myasset-contents-form-component
                  label(for="contents")
                    | メモ
                  textarea(name="contents", rows="5", v-model="formContents")
                button(@click.prevent="postMoneyAccount()")
                  | 保存
              button(v-if="formVisibleFlg" @click="changeFormVisibleFlg()")
                p -
              button(v-else @click="changeFormVisibleFlg()")
                p +

            .main-myasset-contents-details
              h3 サブスク / 固定費
              subscription-item(:subscriptions="subscriptions")
  footer-nav

</template>

<script>
import { priceMixin } from '../mixin/price';
import axios from 'axios';
import jwtDecode from 'jwt-decode';
import HeaderNav from '../molecules/HeaderNav.vue';
import FooterNav from '../molecules/FooterNav.vue';
import ChartPie from '../organisms/ChartPie.vue';
import ChartHorizontalBar from '../organisms/ChartBar.vue';
import MoneyList from '../molecules/MoneyList.vue'
import SubscriptionItem from '../organisms/Subscription.vue';
import DateInputForm from '../molecules/DateInputForm.vue';

export default {
  name: 'MyAsset',
  mixins: [priceMixin],
  components: {
    HeaderNav,
    FooterNav,
    ChartPie,
    ChartHorizontalBar,
    MoneyList,
    SubscriptionItem,
    DateInputForm,
  },
  data () {
    return {
      subscriptions: [],
      userName: '',
      userID: 0,
      thisYear: 2022,
      thisMonth: 3,
      targetYear: 2020,
      targetMonth: 1,
      jwtUserData: '',
      remainingMoney: 0,
      remainingMoneyPerDay: 0,
      income: [],
      expenses: [],
      sumIncome: 0,
      sumExpenses: 0,
      remainingDateOfThisMonth: 0,
      chartData: null,
      chartBarData: null,
      chartOptions: {
        responsive: true,
        title: {
          display: true,
          text: '支出割合',
        },
        legend: {                          //凡例設定
          display: true,                 //表示設定
          position: "bottom",     // タイトルでの position と同じ
          labels: {
            fontSize: 8,
          },
        },
      },
      chartBarOptions: {
        responsive: true,
        title: {
          display: true,
          text: '項目別支出',
        },
        tooltips: {
          mode: 'index',
          intersect: false,
        },
        legend: {                          //凡例設定
          display: false                 //表示設定
        },
        scales: {
          xAxes: [
            {
              display: true,
              ticks: {
                  fontSize: 8             //フォントサイズ
              },
              scaleLabel: {                 //軸ラベル設定
                display: false,             //表示設定
              },
              borderWidth: 0.4,
            },
          ],
          yAxes: [
            {
              ticks: {                      //最大値最小値設定
                fontSize: 8,             //フォントサイズ
                stepSize: 3               //軸間隔
              },
            },
          ],
        },
      },
      formVisibleFlg: false,
      formMoneyAccountLabelOptions: [
        { text: '収入', value: '1' },
        { text: '支出', value: '2' },
      ],
      formDateName: 'formDate',
      formDate: null,
      // formYear: null,
      // formMonth: null,
      formMoneyAccountLabelId: 0,
      formAmount: null,
      formTitle: "",
      formContents: "",
      formSubscriptionsFlg: false,
      formStartDateName: 'formStartDate',
      formEndDateName: 'formEndDate',
      formStartDate: null,
      formEndDate: null,
      // formStartYear: 0,
      // formStartMonth: 0,
      // formEndYear: 0,
      // formEndMonth: 0,
    }
  },
  async created() {
    await this.getDate();
    await this.getUser();
    this.getSubscriptions();
  },
  methods: {
    reflectDate (value) {
      let date = new Date(value);
      this.formDate = date
    },
    reflectStartDate (value) {
      let date = new Date(value);
      this.formStartDate = date
    },
    reflectEndDate (value) {
      let date = new Date(value);
      this.formEndDate = date
    },
    async getSubscriptions () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/subscription/user/' + this.userID,
      {
        headers: {'Authorization': 'Bearer ' + this.jwtUserData}
      })
      .then(res => {
        this.subscriptions = res.data;
        return;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
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
      this.thisYear = date.getFullYear();
      this.thisMonth = date.getMonth()+1;
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
        const dataObjIncome = res.data.moneyAccountListIncome;
        const dataObjExpenses = res.data.moneyAccountListExpenses;
        const incomeSum = res.data.moneyAccountListIncomeSum;
        const expensesSum = res.data.moneyAccountListExpensesSum;
        const remainingMoney = res.data.remainingMoney;
        const remainingDateOfThisMonth = res.data.remainingDateOfThisMonth;
        const remainingMoneyPerDay = res.data.remainingMoneyPerDay;
        this.income = dataObjIncome;
        this.expenses = dataObjExpenses;
        this.sumIncome = incomeSum;
        this.sumExpenses = expensesSum;
        this.remainingMoney = remainingMoney;
        this.remainingDateOfThisMonth = remainingDateOfThisMonth;
        this.remainingMoneyPerDay = remainingMoneyPerDay

        this.chartData = {
          labels: [],
          datasets: [
            {
              data: []
            },
          ],
          options: {
            maintainAspectRatio: true,
            plugins: {
              colorschemes: {
                scheme: 'tableau.Tableau20',
              },
            },
          },
        }

        this.chartBarData = {
          labels: [],
          datasets: [
            {
              data: [],
              barPercentage: 0.5,
              categoryPercentage: 0.4,      //棒グラフ幅
            },
          ],
          options: {
            plugins: {
              colorschemes: {
                scheme: 'tableau.Tableau20',
              },
            },
          },
        }

        dataObjExpenses.forEach (element => {
          this.insertChartData(element);
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
      this.chartBarData.labels.push(data.title);
      this.chartBarData.datasets[0].data.push(data.amount);
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

        let decodedHeader = jwtDecode(this.jwtUserData);
        let user_id = decodedHeader.id;
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
    // 登録
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
      params.append('year', this.formDate.getFullYear())
      params.append('month', this.formDate.getMonth()+1)

      if (this.formSubscriptionsFlg) {
        params.append('subscriptionsFlg', this.formSubscriptionsFlg)
        params.append('startYear', this.formStartDate.getFullYear())
        params.append('startMonth', this.formStartDate.getMonth()+1)
        params.append('endYear', this.formEndDate.getFullYear())
        params.append('endMonth', this.formEndDate.getMonth()+1)

        await axios.post(process.env.VUE_APP_API_BASE_URL + '/api/auth/money_account/subscription/user/new', params, {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
        .then(() => {
          alert( this.formTitle + 'をサブスク/固定費に登録しました。');
          this.getMoneyAccount(this.userID);
          this.getSubscriptions();
          this.formVisibleFlg = false;
          return;
        })
        .catch(error => {
          alert(error);
          return;
        });
      } else {
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
      }
    },
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/components/myAsset.scss">
</style>
