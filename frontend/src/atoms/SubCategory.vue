<template lang="pug">
  .sub-category
    .sub-category-item(v-for="(allSubCatrgory, key) in this.allSubCatrgories" :key="key")
      li(@click="reset(allSubCatrgory.id)")
        | {{allSubCatrgory.name}}
      fa-icon(icon="chevron-right")
</template>

<script>
import axios from 'axios'

export default {
  name: 'SubCategory',
  props: {
    ancestorID: Number,
  },
  data () {
    return {
      allSubCatrgories: [],
    }
  },
  created () {
    this.getAllSubCatrgories();
  },
  methods: {
    reset (id) {
      this.$router.push({name: 'WebsiteListByCategory', params: {category_id: id}})
      location.reload()
    },
    async getAllSubCatrgories () {
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/categories_by_ancestor/' + this.ancestorID)
      .then(res => {
        this.allSubCatrgories = res.data;
      })
      .catch(error => {
        console.log(error);
        return;
      });
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/subCategory.scss">
</style>
