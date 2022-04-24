<template lang="pug">
  .date-form
    datepicker(
      @selected="this.updateDate"
      v-model="date"
      :name="this.name"
      :format="this.datePickerFormat"
      :maximum-view="this.maxViewMode"
      :minimum-view="this.minViewMode"
      :initial-view="this.iniViewMode"
      :calendar-button="this.calendarButton"
      :clear-button="this.clearButton"
      :language="this.language"
      :wrapperClass="this.wrapperClass"
      :inputClass ="this.inputClass"
    )
</template>

<script>
import Datepicker from 'vuejs-datepicker'
import {ja} from 'vuejs-datepicker/dist/locale'

export default {
  name: "DateInputForm",
  components: {
    Datepicker
  },
  props: {
    name: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      date: null,
      datePickerFormat: 'yyyy/MM',
      language: ja,
      // type: "month",
      autoclose: true,
      maxViewMode: 'year',
      minViewMode: 'month',
      iniViewMode: 'month',
      // calendarButton: true,
      // calendarButtonIcon: 'fa-calendar',
      clearButton: true,
      calendarButton: false,
      wrapperClass: 'date-calendar-wrapper',
      inputClass: 'date-calendar-input',
    }
  },
  methods: {
    updateDate() {
      this.$nextTick(() => {
        this.$emit("change-date", this.date);
      });
    }
  }
}
</script>
<style scoped lang="scss" src="../assets/css/molecules/dateInputForm.scss">
</style>