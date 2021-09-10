<template lang="pug">
.webdetail
  h2.webdetail-title
    | website
  .webdetail-contents
    .wrapper
      .webdetail-contents-icon
        fa-icon(icon="building")
        p.webdetail-contents-icon-name
          | {{detailObject.company_name}}
      p.webdetail-contents-txt
        | {{detailObject.name}}
      .webdetail-contents-btm
        .webdetail-contents-btm-left
          .webdetail-contents-btm-left-category
            category-tag(:websiteID="detailObject.id")
          p.webdetail-contents-btm-left-time
            | {{createDate(detailObject.CreatedAt)}}
        .webdetail-contents-btm-right
          cv-button(label="公式サイトへ" :url="detailObject.url")

</template>

<script>
import CvButton from '/src/atoms/CvButton.vue'
import CategoryTag from '/src/atoms/CategoryTag.vue'

export default {
  name: 'WebsiteDetailItem',
  components: {
    CvButton,
    CategoryTag,
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

