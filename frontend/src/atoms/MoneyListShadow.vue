<template lang="pug">
  router-link(:to="{name:'MyAssetItemDetail', params: {id: id}}").block.w-full
    .flex.space-x-4.justify-between.my-3
      .flex-5.min-w-0
        p.text-sm.font-medium.text-gray-900.truncate
          | {{this.label}}
      .flex-5.min-w-0
        p.text-sm.font-medium.text-gray-900.truncate
          | ¥ {{$_commify(this.amount)}}

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
