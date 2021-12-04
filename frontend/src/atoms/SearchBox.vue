<template lang="pug">
  form
    input(type="search" name="search" placeholder="メディア検索" v-model="searchQuery")
    button(@click.prevent="reset()" type="submit" name="submit")
      fa-icon(icon="search")

</template>

<script>

export default {
  name: 'SearchBox',
  props: {
    
  },
  data () {
    return {
      searchQuery: '',
    }
  },
  created () {
    
  },
  methods: {
    reset () {
      this.$router.push({name: 'WebsiteList', query: {q:this.searchQuery}, force: true})
      location.reload()
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/searchBox.scss">
</style>
