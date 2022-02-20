<template lang="pug">
  div
    ul(v-for="i in arr")
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
