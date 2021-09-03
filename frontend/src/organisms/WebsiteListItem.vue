<template lang="pug">
.weblist
  h2.weblist-title 
    | website
  .weblist-contents(v-for="(item, key) in this.allLists" :key="key")
    router-link( :to="{name:'WebsiteDetail', params: {id:`${item.id}`}}").wrapper
      .weblist-contents-icon
        img(src="../../src/assets/images/logo.png")
        p.weblist-contents-icon-name
          | {{item.company_name}}
      p.weblist-contents-txt
        | {{item.name}}
      .weblist-contents-btm
        .weblist-contents-btm-left
          .weblist-contents-btm-left-category
            li カテゴリ
            li カテゴリ
            li カテゴリ
          p.weblist-contents-btm-left-time 
            | {{createDate(item.CreatedAt)}}
        .weblist-contents-btm-right
          cv-button(label="公式サイトへ" :url="item.url") 

</template>

<script>
import CvButton from '/src/atoms/CvButton.vue'

export default {
  name: 'WebsiteListItem',
  components: {
    CvButton,
  },
  props: {
    allLists: Array,
  },
  methods: {
    createDate (createdAt) {
      let formattedDate = createdAt.split(/T|:|-/);
      formattedDate.pop()
      let dateString = formattedDate[0] + '年' + formattedDate[1]+ '月' + formattedDate[2] + '日' + formattedDate[3] + '時' + formattedDate[4] + '分';
      return dateString;
    },
  }
}
</script>

<style scoped lang="scss" src="../assets/css/websiteListItem.scss">
</style>

