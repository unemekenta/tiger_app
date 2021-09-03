<template lang="pug">
.weblist
  h2.weblist-title
    | website
  .weblist-contents
    router-link( :to="{name:'WebsiteDetail', params: {id:`${detailObject.id}`}}").wrapper
      .weblist-contents-icon
        img(src="../../src/assets/images/logo.png")
        p.weblist-contents-icon-name
          | {{detailObject.company_name}}
      p.weblist-contents-txt
        | {{detailObject.name}}
      .weblist-contents-btm
        .weblist-contents-btm-left
          .weblist-contents-btm-left-category
            li カテゴリ
            li カテゴリ
            li カテゴリ
          p.weblist-contents-btm-left-time
            | {{createDate(detailObject.CreatedAt)}}
        .weblist-contents-btm-right
          cv-button(label="公式サイトへ" :url="detailObject.url")

</template>

<script>
import CvButton from '/src/atoms/CvButton.vue'

export default {
  name: 'WebsiteDetailItem',
  components: {
    CvButton,
  },
  props: {
    detailObject: Object,
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

<style scoped lang="scss" src="../assets/css/websiteDetailItem.scss">
</style>

