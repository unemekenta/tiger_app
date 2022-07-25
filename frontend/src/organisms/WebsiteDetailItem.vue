<template lang="pug">
.webdetail
  h2.webdetail-title
    | 詳細情報
  .webdetail-contents
    .text-gray-600.body-font.overflow-hidden.bg-white.w-full
      .container.px-5.py-23.mx-auto.w-full
        .-my-8.divide-y-2.divide-gray-100
          .py-8.flex
            .mr-6.w-56
              img.rounded-lg.shadow-xl(src="../assets/images/no_image.jpg")
            .mb-6.text-left.w-full
              h2.text-3xl.font-medium.text-gray-900.title-font.mb-2 {{detailObject.name}}
              p.leading-relaxed {{detailObject.companyName}}
              p.leading-relaxed
                q(:cite="detailObject.url")
                  | 引用：{{detailObject.name}} 公式HP
              .flex.flex-wrap.gap-3
                category-tag(:websiteID="detailObject.id")
              p.leading-relaxed.text-right.text-xs {{createDate(detailObject.updatedAt)}}

  .webdetail-info.text-left.w-full(v-if="detailObjectContent")
    cv-button(label="公式サイトへ" :url="detailObject.url")
    h3.text-3xl.font-medium.text-gray-900.title-font.mb-2 {{detailObjectContent.title}}
    p.leading-relaxed.my-10 {{detailObjectContent.contents}}
    p.leading-relaxed.my-10.text-xs
      | 参考・引用：<br>公式サイト {{detailObject.url}}


</template>

<script>
import CategoryTag from '/src/atoms/CategoryTag.vue'
import CvButton from '/src/atoms/CvButton.vue'

export default {
  name: 'WebsiteDetailItem',
  components: {
    CategoryTag,
    CvButton,
  },
  props: {
    detailObject: Object,
    detailObjectContent: Object,
  },
  methods: {
    createDate (updatedAt) {
      let formattedDate = updatedAt.split(/T|:|-/);
      formattedDate.pop()
      let dateString = formattedDate[0] + '年' + formattedDate[1]+ '月' + formattedDate[2] + '日';
      return dateString;
    },
  }
}
</script>

<style scoped lang="scss" src="../assets/css/websiteDetailItem.scss">
</style>

