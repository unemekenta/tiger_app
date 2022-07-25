<template lang="pug">
  div
    ul.divide-y.divide-gray-200
      li.text-left.block(v-for="i in arr")
        money-list-item-shadow(:id="i.id" :label="i.title" :amount="i.amount")

</template>

<script>
import MoneyListItemShadow from '../atoms/MoneyListShadow.vue'

export default {
  name: 'MoneyList',
  components: {
    MoneyListItemShadow,
  },
  props: {
    arr: Array,
  },
}
</script>
