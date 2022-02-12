<template lang="pug">
  li.moneylist-shadow(@click.prevent.self="jumpTo(id)")
    p {{this.label}}
    p {{this.amount}} 円

</template>

<script>
export default {
  name: 'MoneyListItemShadow',
  props: {
    label: String,
    amount: Number,
    id: Number,
  },
  methods: {
    jumpTo (id) {
      this.$router.push({name: 'MyAssetItemDetail', params: {id: id}})
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/atoms/MoneyListShadow.scss">
</style>
