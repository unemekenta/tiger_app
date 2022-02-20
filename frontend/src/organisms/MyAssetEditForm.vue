<template lang="pug">
  form.myasset-form
    .myasset-form-component
      label(for="year")
        p 年
      select(name="year" v-model="formYear")
        option(v-for="y in 50" :value="y + 2020")
          | {{ y + 2020 }}年
    .myasset-form-component
      label(for="month")
        p 月
      select(name="month", v-model="formMonth")
        option(v-for="m in 12" :value="m")
          | {{ m }}月
    .myasset-form-component
      label(for="title")
        p 項目名
      input(name="title", type="text", v-model="formTitle")
      p.myasset-form-component-v-error(v-if="isInValidTitle") 項目名は1文字以上で入力してください
    .myasset-form-component
      label(for="label")
        p 増減
      select(name="label", v-model="formMoneyAccountLabelId")
        option(
          v-for="option in formMoneyAccountLabelOptions",
          :value="option.value"
        )
          | {{ option.text }}
    .myasset-form-component
      label(for="amount")
        p 金額(円)
      input(
        name="amount",
        type="number",
        min="0",
        v-model="formAmount"
      )
      p.myasset-form-component-v-error(v-if="isInValidAmount") 金額は1以上で入力してください
    .myasset-form-component
      label(for="contents")
        p メモ
      textarea(name="contents", rows="5", v-model="formContents")
      p.myasset-form-component-v-error(v-if="isInValidContents") メモは3000文字以下で入力してください
    button.basic-btn(@click.prevent="putMoneyAccount()" :disabled="isInValidSubmit")
      p 更新
    button.warn-btn(@click.prevent="deleteMoneyAccount()")
      p 削除

</template>

<script>
import axios from "axios";

export default {
  name: "MyAssetEditForm",
  components: {
  },
  props: {
    userID: Number,
    jwtUserData: String,
    itemId: Number,
    year: Number,
    month: Number,
    moneyAccountLabelId: Number,
    amount: Number,
    title: String,
    contents: String,
  },
  data() {
    return {
      formMoneyAccountLabelOptions: [
        { text: "収入", value: "1" },
        { text: "支出", value: "2" },
      ],
      newYear: 0,
      newMonth: 0,
      newMoneyAccountLabelId: 0,
      newAmount: null,
      newTitle: "",
      newContents: "",
    };
  },
  created() {
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
    isInValidYear(){
      //数値が1文字以上かチェックする
      return ((this.newYear < 1 || this.newYear == "") && (this.year < 1 || this.year == ""));
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
    isInValidMonth(){
      //数値が1文字以上かチェックする
      return ((this.newMonth < 1 || this.newMonth == "") && (this.month < 1 || this.month == ""));
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
    isInValidTitle(){
      //文字列が1文字以上かチェックする
      return (this.newTitle.length < 1 && this.title.length < 1);
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
    isInValidMoneyAccountLabelId(){
      //数値が1文字以上かチェックする
      return ((this.newMoneyAccountLabelId < 1 || this.newMoneyAccountLabelId == "") && (this.moneyAccountLabelId < 1 || this.moneyAccountLabelId == ""));
    },
    formAmount: {
      get() {
        if (this.newAmount == null) {
          return this.amount;
        }
        this.clearAmount();
        return this.newAmount;
      },
      set(val) {
        this.newAmount = val;
      },
    },
    isInValidAmount(){
      //数値が1文字以上かチェックする
      return ((this.newAmount < 1 || this.newAmount == "") && (this.amount < 1 || this.amount == ""));
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
    isInValidContents() {
      //文字列が3000文字以下かチェックする
      return (this.newContents.length > 3000);
    },
    isInValidSubmit() {
      return (this.isInValidYear || this.isInValidMonth || this.isInValidTitle || this.isInValidMoneyAccountLabelId || this.isInValidAmount || this.isInValidContents);
    },
  },
  methods: {
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
      this.amount = null;
    },
    clearContents() {
      this.contents = "";
    },
  },
};
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/organisms/myAssetForm.scss">
</style>
