<template lang="pug">
.weblist
  h2.weblist-title
    | メディア一覧
  .weblist-contents
    section.text-gray-600.body-font.overflow-hidden.bg-white.w-full
      .container.px-5.py-23.mx-auto.w-full(v-for="(item, key) in this.slicedList" :key="key")
        router-link( :to="{name:'WebsiteDetail', params: {id:`${item.id}`}}").-my-8.divide-y-2.divide-gray-100
          .py-8.flex
            .mr-6.w-56
              img.rounded-lg.shadow-xl(src="@/assets/images/no_image.jpg")
            .mb-6.text-left.w-full
              h2.text-3xl.font-medium.text-gray-900.title-font.mb-2 {{item.name}}
              p.leading-relaxed {{item.companyName}}
              p.leading-relaxed
                q.text-xs(:cite="item.url")
                  | 引用：{{item.name}} 公式HP
              .flex.flex-wrap.gap-3
                category-tag(:websiteID="item.id")
              p.leading-relaxed.text-right.text-xs {{createDate(item.updatedAt)}}

  pager-item(:maxPage="maxPage" @changePage="this.changePageNum")

</template>

<script>
import PagerItem from '/src/atoms/Pager.vue'
import CategoryTag from '/src/atoms/CategoryTag.vue'

export default {
  name: 'WebsiteListItem',
  components: {
    PagerItem,
    CategoryTag,
  },
  props: {
    allLists: Array,
  },
  data () {
    return {
      itemNumPerPage: 20,
      maxPage: 1,
      currentPage: 1,
      slicedList: [],
    }
  },
  created () {
    this.maxPage = this.createMaxPage;
    this.slicedList = this.createListsPerPage;
  },
  watch: {
    allLists() {
      this.maxPage = this.createMaxPage;
      this.slicedList = this.createListsPerPage;
    }
  },
  computed: {
    createMaxPage: {
      get () {
        let maxPageNum = Math.ceil(this.allLists.length / this.itemNumPerPage);
        return maxPageNum;
      }
    },
    createListsPerPage: {
      get () {
        let current = this.currentPage * this.itemNumPerPage;
        let start = current - this.itemNumPerPage;
        let slicedList = this.allLists.slice(start, current);
        return slicedList;
      }
    }
  },
  methods: {
    createDate (updatedAt) {
      let formattedDate = updatedAt.split(/T|:|-/);
      formattedDate.pop()
      let dateString = formattedDate[0] + '年' + formattedDate[1]+ '月' + formattedDate[2] + '日';
      return dateString;
    },
    changePageNum (pageNum) {
      this.currentPage = pageNum;
      this.slicedList = this.createListsPerPage;
    },
  }
}
</script>

<style scoped lang="scss" src="../assets/css/websiteListItem.scss">
</style>

