<template lang="pug">
  li.flex.justify-center.m-2.px-6.py-2.items-center.bg-green-400.text-white.text-sm.font-semibold.text-center.border.rounded-md.transition.duration-100
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
      await axios.get(process.env.VUE_APP_API_BASE_URL + '/api/categories_website/' + website_id)
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