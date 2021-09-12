<template lang="pug">
  ul.card-box
    li(v-for="item in this.itemList" @click="reset(item.id)")
      .card-box-txt
        p {{item.name}}

</template>

<script>

export default {
  name: 'TagSlider',
  props: {
    itemList: Array,
  },
  methods: {
    reset (id) {
      this.$router.push({name: 'WebsiteListByCategory', params: {category_id: id}})
    }
  }
}
</script>

<style scoped lang="scss" src="../assets/css/tagSlider.scss">
</style>
