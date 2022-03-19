<script>
import { HorizontalBar, mixins } from 'vue-chartjs'
import 'chartjs-plugin-colorschemes'

export default {
  extends: HorizontalBar,
  mixins: [mixins.reactiveProp],
  name: 'ChartHorizontalBar',
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
