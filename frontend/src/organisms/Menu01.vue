<template lang="pug">
  .menu01.bg-white.flex.flex-col.h-screen.px-4.py-8.border-r.rounded-lg
    .menu01-content.text-left
      .menu01-content-head
        | カテゴリ
      aside.menu01-content-item.w-full(aria-label="Sidebar")
        .px-3.py-4.overflow-y-auto.rounded.bg-white(v-for="(allCatrgory, key) in this.allCategories" :key="key")
          ul.space-y-2
            li.font-bold.my-2
              | {{allCatrgory.name}}
            sub-category(:ancestorID = "allCatrgory.id")

</template>

<script>
import SubCategory from '/src/atoms/SubCategory.vue'

export default {
  name: 'Menu01Item',
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
