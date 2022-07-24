<template lang="pug">
.menu02
  .menu02-content
    .menu02-content-head
      | メニュー
    .menu02-content-items
      search-box
      .menu02-content-items-item
        router-link( to="/" )
          p
            fa-icon(icon='home')
            | HOME
      .menu02-content-items-item
        router-link( to="/website_list" )
          p
            fa-icon(icon='list')
            | メディア一覧
      .menu02-content-items-item
        router-link( to="/my_asset" )
          p
            fa-icon(icon='coins')
            | あといくら？
      .menu02-content-items-item
        router-link( to="/website/new" )
          p
            fa-icon(icon='book')
            | Book Log

</template>

<script>
import SearchBox from '/src/atoms/SearchBox.vue'

export default {
  name: 'Menu02Item',
  components: {
    SearchBox,
  },
  data () {
    return {
    }
  },
  created () {
  },
  methods: {
  },
}
</script>

<style scoped lang="scss" src="../assets/css/menu02.scss">
</style>

