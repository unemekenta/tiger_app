<template lang="pug">
  ul.card-box
    li(v-for="item in this.itemList")
      img(src="../assets/images/top-bg.png")
      .card-box-txt
        p title
        p あああ

</template>

<script>

export default {
  name: 'CardSlider',
  props: {
    itemList: Array,
  }
}
</script>

<style scoped lang="scss" src="../assets/css/carLlist.scss">
</style>
