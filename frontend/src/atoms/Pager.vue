<template lang="pug">
  paginate(
    v-model="currentPage"
    :page-count="maxPage"
    :page-range="2"
    :margin-pages="2"
    :click-handler="clickCallback"
    :prev-class="'m-0'"
    :prev-link-class="'block py-2 px-3 ml-0 leading-tight text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700'"
    :next-class="'m-0'"
    :next-link-class="'block py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700'"
    :container-class="'inline-flex items-center -space-x-px my-4'"
    :page-link-class="'block py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700'"
    :active-class="'text-green-700'"
    :disabled-class="'text-gray-700'"
  )
</template>

<script>
import Paginate from 'vuejs-paginate'

export default {
  name: 'PagerItem',
  components: {
    Paginate,
  },
  methods: {
    clickCallback (pageNum) {
      this.currentPage = Number(pageNum);
      this.$emit('changePage', pageNum);
    }
  },
  data () {
    return {
      currentPage: 1,
    }
  },
  props: {
    maxPage: Number,
  },
}
</script> 

<style lang="scss" src="../assets/css/pager.scss">
// pagerはscopedが効かないため記述しない
</style>
