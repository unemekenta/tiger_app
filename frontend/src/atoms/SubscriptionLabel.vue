<template lang="pug">
  router-link(:to="{name:'MyAssetItemDetail', params: {id: id}}")
    li.subscription-label
      p
        fa-icon(icon="check-circle")
        | {{this.label}}
      p ¥ {{$_commify(this.amount)}}

</template>

<script>
import { priceMixin } from '../mixin/price'

export default {
  name: 'SubscriptionLabel',
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

<style scoped lang="scss" src="../assets/css/atoms/subscriptionLabel.scss">
</style>
