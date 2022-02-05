<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  .l-container#container
    .top
      .user
        //- img.user-icon(src="../../src/assets/images/logo.png")
        .user-name
          b {{userName}}
          |さんのページ
      .main
        menu01-item(:allCategories="allCategories").is-only-pc
        .main-mypage
          h2.main-mypage-title マイクリップ
          .main-mypage-contents
          //-  p 準備中..
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
            //- .main-mypage-contents-item
          chart-pie(
            :chartData="chartIncomeData"
            :options="chartOptions"
          )
          chart-pie(
            :chartData="chartExpenseData"
            :options="chartOptions"
          )
        menu02-item.is-only-pc
  footer-nav

</template>

<script lang="ts">
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import Menu01Item from '../organisms/Menu01.vue'
import Menu02Item from '../organisms/Menu02.vue'
import HeaderNav from '../molecules/HeaderNav.vue'
import FooterNav from '../molecules/FooterNav.vue'
import ChartPie from "../organisms/ChartPie.vue";

export default {
  name: 'MyAsset',
  components: {
    Menu01Item,
    Menu02Item,
    HeaderNav,
    FooterNav,
    ChartPie
  },
  data () {
    return {
      userName: '',
      allCategories: [],
      jwtUserData: '',

      chartIncomeData: {
        labels: [],
        // 表示するデータ
        datasets: [
          {
            data: [],
            backgroundColor: ['#f87979', '#aa4c8f', '#38b48b', '#006e54', '#c1e4e9', '#89c3eb', '#c3d825']
          }
        ]
      },
      chartExpenseData: {
        labels: [],
        // 表示するデータ
        datasets: [
          {
            data: [],
            backgroundColor: ['#f87979', '#aa4c8f', '#38b48b', '#006e54', '#c1e4e9', '#89c3eb', '#c3d825']
          }
        ]
      },
      chartOptions: {
        responsive: true
      },
    }
  },
  created () {
    this.getUser();
    this.getAllCategories();
  },
  methods: {
    async getMoneyAccount (userID: number) {
      console.log(process.env.VUE_APP_API_BASE_URL + '/api/money_account/user/' + userID)
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/money_account/user/' + userID)
      .then(res => {
        console.log("asset--------------------", res.data)
        const dataObj = res.data
        for (let i in dataObj) {
          let incomeCheck = dataObj[i].moneyAccountLabelId;
          switch (incomeCheck) {
            case 1:
              console.log("incomeCheck", incomeCheck, dataObj[i])
              this.insertChartData(dataObj, i, "chartIncomeData");
              break;
            case 2:
              // this.insertChartData(dataObj, i, this.chartExpenseData);
              break;
            default:
              break;
          }
          // this.chartData.labels[i] = dataObj[i].title;
          // console.log(this.chartData.labels, typeof(this.chartData.labels))
          // if (this.chartData.datasets) {
          //   console.log(this.chartData.datasets[0].data, typeof(this.chartData.datasets.data));
          //   this.chartData.datasets[0].data[i] = dataObj[i].amount;
          // }
        }
        // this.allCategories = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
    insertChartData(dataObj: any, i: number, chartData: string) {
      switch (chartData) {
        case "chartIncomeData":
          if (this.chartIncomeData) {
            console.log("--ここからinsertChartData-----------", typeof(this.chartIncomeData.labels))
            this.chartIncomeData.labels.data.push(dataObj[i].title);
            if (this.chartIncomeData.datasets) {
              this.chartIncomeData.datasets[0].data.push(dataObj[i].amount);
            }
          }
          break;

        default:
          break;
      }
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
        await axios
          .post(
            process.env.VUE_APP_API_BASE_URL + "/api/login_check",
            {
              uuid: window.$cookies.get('uuid')
            },
          )
          .then((res: any) => {
            this.jwtUserData = res.data;
          })
          .catch((error: any) => {
            alert('ログインしてください');
            return error;
          });

        let user_id = VueJwtDecode.decode(this.jwtUserData).id;
        await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/auth/user/'+ user_id,
        {
          headers: {'Authorization': 'Bearer ' + this.jwtUserData}
        })
          .then((res: any) => {
            this.userName = res.data.name;
            const userID = res.data.id;
            this.getMoneyAccount(userID);
          })
          .catch((error: any) => {
            alert('権限がありません');
            return error;
          });
        return;
      } else {
        alert('ログインしてください');
        return;
      }
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/myPage.scss">
</style>
