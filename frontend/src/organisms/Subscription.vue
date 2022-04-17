<template lang="pug">
  .subscriptions
    .subscriptions-content
      .subscriptions-content-head
        subscription-list(:arr="subscriptions")
</template>

<script>
import SubscriptionList from '../molecules/SubscriptionList.vue'

export default {
  name: 'Subscription',
  components: {
    SubscriptionList,
  },
  props: {
    subscriptions: Array,
  },
  methods: {
  }
}
</script>

<style scoped lang="scss" src="../assets/css/organisms/subscription.scss">
</style>
