<script>
import { Pie, mixins } from 'vue-chartjs'
import 'chartjs-plugin-colorschemes'

export default {
  extends: Pie,
  mixins: [mixins.reactiveProp],
  name: 'ChartPie',
  props: {
    chartData: Object,
    options: Object,
  },
  data () {
    return {
    }
  },
  created () {
  },
  mounted () {
    if (this.chartData) {
      this.renderChart(this.chartData, this.options);
    }
  },
  watch: {
    chartData: function onChangeData() {
      if (this.chartData) {
        this.renderChart(this.chartData, this.options);
      }
    }
  },
  methods: {
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped lang="scss" src="../assets/css/organisms/chartPie.scss">
</style>
