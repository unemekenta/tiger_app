<template lang="pug">
  paginate(
    v-model="currentPage"
    :page-count="maxPage"
    :page-range="2"
    :margin-pages="2"
    :click-handler="clickCallback"
    :prev-text="'<'"
    :next-text="'>'"
    :container-class="'pagination'"
    :page-class="'pagination-item'"
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
