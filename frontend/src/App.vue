<template lang="pug">
  #app
    router-view
</template>
<script>
export default {
  mounted() {
      var routeInstance = this.$route;
      this.createTitleDesc(routeInstance);
  },
  methods: {
    createTitleDesc : function(routeInstance){
       // タイトルを設定
      if(routeInstance.meta.title){
        var setTitle = routeInstance.meta.title + ' | WEB MEDIA DB';
        document.title = setTitle;
      } else {
        document.title = 'WEB MEDIA DB'
      }
    },
  }
}
</script>

<style lang="scss">
@import "/assets/css/common/variable.scss";
@import "/assets/css/common/reset.scss";
@import '/assets/css/common/tailwind.scss'
</style>