<template lang="pug">
  .menu01
    .menu01-content
      .menu01-content-head
        | カテゴリ
      .menu01-content-item(v-for="(allCatrgory, key) in this.allCategories" :key="key")
        li(@click="reset(allCatrgory.id)")
          | {{allCatrgory.name}}

</template>

<script>

export default {
  name: 'Menu01',
  props: {
    allCategories: Array,
  },
  methods: {
    reset (id) {
      this.$router.push({name: 'WebsiteListByCategory', params: {category_id: id}})
    }
  }
}
</script>

<style scoped lang="scss" src="../assets/css/menu01.scss">
</style>
