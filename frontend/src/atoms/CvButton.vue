<template lang="pug">
button.block.w-full.bg-green-600.my-3.text-gray-100.font-bold.border.border-gray-200.rounded-lg.py-3.px-3.leading-tight(type="button" @click.prevent.self="jumpTo(url)")
  | {{this.label}}
  fa-icon.ml-5(icon="angle-double-right")

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
      window.open(url, '_blank')
    },
  },
}
</script>
