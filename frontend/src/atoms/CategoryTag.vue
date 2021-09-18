<template lang="pug">
  li.category-tag(v-if="category")
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
      await axios.get('http://localhost:8000/api/categories_website/' + website_id)
      .then(res => {
        const category = res.data.name;
        this.category = category;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/categoryTag.scss">
</style>
