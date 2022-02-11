<template lang="pug">
  li.moneylist-shadow(@click.prevent.self="jumpTo(url)")
    p {{this.label}}
    p {{this.amount}} 円

</template>

<script>
export default {
  name: 'MoneyListItemShadow',
  props: {
    label: String,
    amount: Number,
    url: String,
  },
  methods: {
    jumpTo (url) {
      window.open(url, '_blank')
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/atoms/MoneyListShadow.scss">
</style>
