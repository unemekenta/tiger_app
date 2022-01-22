<template lang="pug">
.weblist
  h2.weblist-title
    | メディア一覧
  .weblist-contents
    .weblist-contents-inner(v-for="(item, key) in this.slicedList" :key="key")
      router-link( :to="{name:'WebsiteDetail', params: {id:`${item.id}`}}").wrapper
        .weblist-contents-inner-img(v-if="item.image")
          img(:src="item.image")
        .weblist-contents-inner-img(v-else)
          img(src="../assets/images/no_image.jpg")
        .weblist-contents-inner-quote
          q(cite="item.url")
            | 引用：{{item.name}} 公式HP
        .weblist-contents-inner-icon
          fa-icon(icon="building")
          p.weblist-contents-inner-icon-name
            | {{item.companyName}}
        p.weblist-contents-inner-txt
          | {{item.name}}
        .weblist-contents-inner-btm
          .weblist-contents-inner-btm-left
            .weblist-contents-inner-btm-left-category
              category-tag(:websiteID="item.id")
            p.weblist-contents-inner-btm-left-time
              | {{createDate(item.updatedAt)}}
          .weblist-contents-inner-btm-right
            cv-button(label="公式サイトへ" :url="item.url")
  pager(:maxPage="maxPage" @changePage="this.changePageNum")

</template>

<script>
import CvButton from '/src/atoms/CvButton.vue'
import Pager from '/src/atoms/Pager.vue'
import CategoryTag from '/src/atoms/CategoryTag.vue'

export default {
  name: 'WebsiteListItem',
  components: {
    CvButton,
    Pager,
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

