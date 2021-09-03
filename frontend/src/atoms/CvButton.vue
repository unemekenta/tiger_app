<template lang="pug">
button.cvbutton(type="button" @click.prevent.self="jumpTo(url)") 
  | {{this.label}}
  fa-icon(icon="angle-double-right")

</template>

<script>
export default {
  name: 'CvButton',
  props: {
    label: String,
    url: String,
  },
  methods: {
    jumpTo (url) {
      location.href= url;
    },
  },
}
</script>

<style scoped lang="scss" src="../assets/css/cvButton.scss">
</style>
