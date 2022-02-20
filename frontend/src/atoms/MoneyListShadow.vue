<template lang="pug">
  router-link(:to="{name:'MyAssetItemDetail', params: {id: id}}")
    li.moneylist-shadow
      p {{this.label}}
      p {{$_commify(this.amount)}} 円

</template>

<script>
import { priceMixin } from '../mixin/price'

export default {
  name: 'MoneyListItemShadow',
  mixins: [priceMixin],
  props: {
    label: String,
    amount: Number,
    id: Number,
  },
  methods: {
  },
}
</script>

<style scoped lang="scss" src="../assets/css/atoms/moneyListShadow.scss">
</style>
