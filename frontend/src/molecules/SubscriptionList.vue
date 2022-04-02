<template lang="pug">
  div
    ul(v-for="i in arr")
      subscription-label(:id="i.id" :label="i.title" :amount="i.amount")

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
