<template lang="pug">
.menu02.bg-white.flex.flex-col.h-screen.px-4.py-8.border-r.rounded-lg
  .menu02-content
    .menu02-content-head
      | メニュー
    .menu02-content-items
      search-box
      .bg-green-400.text-grey-darkest.font-bold.py-2.px-4.rounded.inline-flex.items-center.my-4(class="hover:bg-gray-300")
        router-link( to="/" )
          p
            fa-icon.mr-4(icon='home')
            | HOME
      .bg-green-400.text-grey-darkest.font-bold.py-2.px-4.rounded.inline-flex.items-center.my-4(class="hover:bg-gray-300")
        router-link( to="/website_list" )
          p
            fa-icon.mr-4(icon='list')
            | メディア一覧
      .bg-green-400.text-grey-darkest.font-bold.py-2.px-4.rounded.inline-flex.items-center.my-4(class="hover:bg-gray-300")
        router-link( to="/my_asset" )
          p
            fa-icon.mr-4(icon='coins')
            | あといくら？
      .bg-green-400.text-grey-darkest.font-bold.py-2.px-4.rounded.inline-flex.items-center.my-4(class="hover:bg-gray-300")
        router-link( to="/website/new" )
          p
            fa-icon.mr-4(icon='book')
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

