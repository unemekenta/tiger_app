<template lang="pug">
  .sub-category
    ul.sub-category-item(v-for="(allSubCatrgory, key) in this.allSubCatrgories" :key="key")
      li.flex.items-center.p-2.text-base.font-normal.text-gray-900.rounded-lg.cursor-pointer(@click="reset(allSubCatrgory.id)" class="hover:bg-gray-100")
        | {{allSubCatrgory.name}}

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
