export const priceMixin = {
  methods: {
    $_commify(price) {
        return this.getPriceAsNumber(price).toLocaleString()
    },
    getPriceAsNumber(price) { // 文字列ではなく、数字として金額を取得
        price = price.toString().replace(/[^0-9]+/g, '') // 数字以外を除去
        return parseInt(price);
    }
  }
};