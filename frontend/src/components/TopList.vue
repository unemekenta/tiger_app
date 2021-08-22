<template lang="pug">
.allWrapper
  .main(v-for="(term, key) in this.allTerms" :key="key")
    li {{term.name}}
    li {{term.content}}

</template>

<script>
import axios from 'axios'

export default {
  name: 'TopList',
  data () {
    return {
      allTerms: [],
    }
  },
  async created () {
    let resAllTerms = await axios.get('http://localhost:8000/api/terms');
    console.log(resAllTerms.data);
    this.allTerms = resAllTerms.data;
  },
  methods: {
  },
  computed: {
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/topList.scss";
</style>
