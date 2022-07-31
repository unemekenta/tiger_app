<template lang="pug">
.allWrapper
  header-nav(:allCategories="allCategories")
  #container.l-container
    .top
      .main
        .main-myasset-detail
          h2.main-myasset-detail-title 登録項目詳細
          .main-myasset-detail-contents
            .main-myasset-detail-contents-form
              my-asset-edit-form(
                :userID="this.userID",
                :jwtUserData="this.jwtUserData",
                :itemId="this.itemId",
                :year="this.year",
                :month="this.month",
                :moneyAccountLabelId="this.moneyAccountLabelId",
                :amount="this.amount",
                :title="this.title",
                :contents="this.contents",
              )
          .main-myasset-detail-link-toback
            router-link(:to="{name:'MyAsset'}")
              fa-icon(icon='angle-left')
              p 一覧へ戻る

  footer-nav
</template>

<script>
import axios from "axios";
import jwtDecode from 'jwt-decode';
import HeaderNav from "../molecules/HeaderNav.vue";
import FooterNav from "../molecules/FooterNav.vue";
import MoneyList from "../molecules/MoneyList.vue";
import MyAssetEditForm from "../organisms/MyAssetEditForm.vue";

export default {
  name: "MyAssetItemDetail",
  components: {
    HeaderNav,
    FooterNav,
    MoneyList,
    MyAssetEditForm,
  },
  data() {
    return {
      userName: "",
      userID: 0,
      jwtUserData: "",
      itemId: 0,
      year: 0,
      month: 0,
      moneyAccountLabelId: 0,
      amount: 0,
      title: "",
      contents: "",
    };
  },
  async created() {
    await this.getUser();
    this.getItemData();
  },
  mounted() {
    this.formTitle = this.title
  },
  methods: {
    async getItemData() {
      await axios
        .get(
          process.env.VUE_APP_API_BASE_URL +
            "/api/auth/money_account/user/detail/" +
            this.$route.params.id,
          {
            headers: { Authorization: "Bearer " + this.jwtUserData },
          }
        )
        .then((res) => {
          const d = res.data;
          this.itemId = d.id;
          this.title = d.title;
          this.year= d.year
          this.month = d.month;
          this.moneyAccountLabelId = d.moneyAccountLabelId;
          this.amount = d.amount;
          this.contents = d.contents;
          return;
        })
        .catch((error) => {
          console.log(error);
          return;
        });
    },
    async getUser() {
      if (window.$cookies.isKey("uuid")) {
        await axios
          .post(process.env.VUE_APP_API_BASE_URL + "/api/login_check", {
            uuid: window.$cookies.get("uuid"),
          })
          .then((res) => {
            this.jwtUserData = res.data;
          })
          .catch((error) => {
            alert("ログインしてください");
            return error;
          });
        let decodedHeader = jwtDecode(this.jwtUserData);
        let user_id = decodedHeader.id;
        await axios
          .get(process.env.VUE_APP_API_BASE_URL + "/api/auth/user/" + user_id, {
            headers: { Authorization: "Bearer " + this.jwtUserData },
          })
          .then((res) => {
            this.userName = res.data.name;
            this.userID = res.data.id;
            return;
          })
          .catch((error) => {
            alert("権限がありません。");
            return error;
          });
      } else {
        alert("ログインしてください");
        return;
      }
    },
  },
};
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/components/myAssetItemDetail.scss">
</style>
