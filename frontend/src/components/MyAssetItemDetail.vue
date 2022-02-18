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
              form
                .main-myasset-detail-contents-form-component
                  label(for="year")
                    p 年
                  select(name="year" v-model="formYear")
                    option(v-for="y in 50" :value="y + 2020")
                      | {{ y + 2020 }}年
                .main-myasset-detail-contents-form-component
                  label(for="month")
                    p 月
                  select(name="month", v-model="formMonth")
                    option(v-for="m in 12" :value="m")
                      | {{ m }}月
                .main-myasset-detail-contents-form-component
                  label(for="title")
                    p 項目名
                  input(name="title", type="text", v-model="formTitle")
                .main-myasset-detail-contents-form-component
                  label(for="label")
                    p 増減
                  select(name="label", v-model="formMoneyAccountLabelId")
                    option(
                      v-for="option in formMoneyAccountLabelOptions",
                      :value="option.value"
                    )
                      | {{ option.text }}
                .main-myasset-detail-contents-form-component
                  label(for="amount")
                    p 金額(円)
                  input(
                    name="amount",
                    type="number",
                    min="0",
                    v-model="formAmount"
                  )
                .main-myasset-detail-contents-form-component
                  label(for="contents")
                    p メモ
                  textarea(name="contents", rows="5", v-model="formContents")
                button.basic-btn(@click.prevent="putMoneyAccount()")
                  p 更新
                button.warn-btn(@click.prevent="deleteMoneyAccount()")
                  p 削除

  footer-nav
</template>

<script>
import axios from "axios";
import VueJwtDecode from "vue-jwt-decode";
import HeaderNav from "../molecules/HeaderNav.vue";
import FooterNav from "../molecules/FooterNav.vue";
import MoneyList from "../molecules/MoneyList.vue";

export default {
  name: "MyAssetItemDetail",
  components: {
    HeaderNav,
    FooterNav,
    MoneyList,
  },
  data() {
    return {
      userName: "",
      userID: 0,
      allCategories: [],
      jwtUserData: "",
      formMoneyAccountLabelOptions: [
        { text: "収入", value: "1" },
        { text: "支出", value: "2" },
      ],
      itemId: 0,
      year: 0,
      month: 0,
      moneyAccountLabelId: 0,
      amount: 0,
      title: "",
      contents: "",
      newYear: 0,
      newMonth: 0,
      newMoneyAccountLabelId: 0,
      newAmount: 0,
      newTitle: "",
      newContents: "",
    };
  },
  async created() {
    await this.getUser();
    this.getItemData();
  },
  mounted() {
    this.formTitle = this.title
  },
  computed: {
    formYear: {
      get() {
        if (!this.newYear) {
          return this.year;
        }
        return this.newYear;
      },
      set(val) {
        this.newYear = val;
      },
    },
    formMonth: {
      get() {
        if (!this.newMonth) {
          return this.month;
        }
        return this.newMonth;
      },
      set(val) {
        this.newMonth = val;
      },
    },
    formTitle: {
      get() {
        if (!this.newTitle) {
          return this.title;
        }
        this.clearTitle();
        return this.newTitle;
      },
      set(val) {
        this.newTitle = val;
      },
    },
    formMoneyAccountLabelId: {
      get() {
        if (!this.newMoneyAccountLabelId) {
          return this.moneyAccountLabelId;
        }
        return this.newMoneyAccountLabelId;
      },
      set(val) {
        this.newMoneyAccountLabelId = val;
      },
    },
    formAmount: {
      get() {
        if (!this.newAmount) {
          return this.amount;
        }
        this.clearAmount();
        return this.newAmount;
      },
      set(val) {
        this.newAmount = val;
      },
    },
    formContents: {
      get() {
        if (!this.newContents) {
          return this.contents;
        }
        this.clearContents();
        return this.newContents;
      },
      set(val) {
        this.newContents = val;
      },
    },
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
        console.log(window.$cookies.get("uuid"));
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

        let user_id = VueJwtDecode.decode(this.jwtUserData).id;
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
    async deleteMoneyAccount() {
      await axios
        .delete(
          process.env.VUE_APP_API_BASE_URL +
            "/api/auth/money_account/user/" +
            this.itemId,
          {
            headers: { Authorization: "Bearer " + this.jwtUserData },
          }
        )
        .then(() => {
          alert("削除しました。");
          this.$router.push({ name: "MyAsset" });
        })
        .catch((error) => {
          alert(error);
        });
    },
    async putMoneyAccount() {
      let params = new URLSearchParams();
      params.append("userId", this.userID);
      params.append("moneyAccountLabelId", this.formMoneyAccountLabelId);
      params.append("amount", this.formAmount);
      params.append("title", this.formTitle);
      params.append("contents", this.formContents);
      params.append("year", this.formYear);
      params.append("month", this.formMonth);
      await axios
        .put(
          process.env.VUE_APP_API_BASE_URL +
            "/api/auth/money_account/user/" +
            this.itemId,
          params,
          {
            headers: { Authorization: "Bearer " + this.jwtUserData },
          }
        )
        .then((res) => {
          alert(res.data.title + "を更新しました。");
          this.$router.push({ name: "MyAsset" });
        })
        .catch((error) => {
          alert(error);
        });
    },
    clearTitle() {
      this.title = "";
    },
    clearAmount() {
      this.amount = 0;
    },
    clearContents() {
      this.contents = "";
    },
  },
};
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/components/myAssetItemDetail.scss">
</style>
