<template lang="pug">
  router-link(:to="{name:'MyAssetItemDetail', params: {id: id}}")
    li.subscription-label
      .subscription-label-contents
        p.title
          fa-icon(icon="check-circle")
          | {{this.label}}
        p.amount ¥ {{$_commify(this.amount)}}
      .subscription-label-period
        p.period(v-if="this.startYear && this.startMonth && this.endYear && this.endMonth")
          | {{this.startYear}}/{{this.startMonth}} ~ {{this.endYear}}/{{this.endMonth}}
        p.period(v-else-if="this.startYear && this.startMonth")
          | {{this.startYear}}/{{this.startMonth}} ~
        p.period(v-else-if="this.endYear && this.endMonth")
          | ~ {{this.endYear}}/{{this.endMonth}}

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
    startYear: Number,
    startMonth: Number,
    endYear: Number,
    endMonth: Number,
  },
  methods: {
  },
}
</script>

<style scoped lang="scss" src="../assets/css/atoms/subscriptionLabel.scss">
</style>
