<template lang="pug">
  div
    ul(v-for="i in arr")
      subscription-label(:id="i.moneyAccount.id" :label="i.moneyAccount.title" :amount="i.moneyAccount.amount" :startYear="i.startYear" :startMonth="i.startMonth" :endYear="i.endYear" :endMonth="i.endMonth")

</template>

<script>
import SubscriptionLabel from '../atoms/SubscriptionLabel.vue'

export default {
  name: 'SubscriptionList',
  components: {
    SubscriptionLabel,
  },
  props: {
    arr: Array,
  },
}
</script>
