<template lang="pug">
  .card-box
    .card-box-inner(v-for="item in this.itemList")
      .card-box-txt
        p {{item.name}}
      sub-category(:ancestorID = "item.id")

</template>

<script>
import SubCategory from '/src/atoms/SubCategory.vue'

export default {
  name: 'TagSlider',
  components: {
    SubCategory,
  },
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
