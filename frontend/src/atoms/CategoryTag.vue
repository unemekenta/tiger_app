<template lang="pug">
  li.category-tag
    | {{category}}

</template>

<script>
import axios from 'axios'

export default {
  name: 'CategoryTag',
  props: {
    websiteID: Number,
  },
  data () {
    return {
      category: ''
    }
  },
  created () {
    this.getCategory();
  },
  watch: {
    websiteID() {
      this.category = this.getCategory();
    }
  },
  methods: {
    async getCategory () {
      const website_id = this.websiteID;
      let resCategory = await axios.get('http://localhost:8000/api/categories_website/' + website_id);
      console.log('http://localhost:8000/api/categories_website/' + website_id);
      const category = resCategory.data.name;
      console.log(category);
      this.category = category;
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/categoryTag.scss">
</style>
