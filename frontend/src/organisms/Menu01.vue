<template lang="pug">
  .menu01
    .menu01-content
      .menu01-content-head
        | カテゴリ
      .menu01-content-item(v-for="(allCatrgory, key) in this.allCategories" :key="key")
        li
          | {{allCatrgory.name}}
        sub-category(:ancestorID = "allCatrgory.id")

</template>

<script>
import SubCategory from '/src/atoms/SubCategory.vue'

export default {
  name: 'Menu01',
  components: {
    SubCategory,
  },
  props: {
    allCategories: Array,
  },
  methods: {
  }
}
</script>

<style scoped lang="scss" src="../assets/css/menu01.scss">
</style>
