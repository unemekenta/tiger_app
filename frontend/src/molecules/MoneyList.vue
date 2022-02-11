<template lang="pug">
  div
    ul(v-for="i in arr")
      money-list-item-shadow(:label="i.title" :amount="i.amount" :url="i.title")

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

<style scoped lang="scss" src="../assets/css/atoms/MoneyListShadow.scss">
</style>
