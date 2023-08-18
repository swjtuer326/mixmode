<template>
      <div ref="chartDom1" class="linechart"></div>
</template>
    
<script setup>
import * as echarts from "echarts"
import { onMounted, onUnmounted, ref} from 'vue'

const props = defineProps({
    interval: Number,
    max: Number,
    title: String,
    line1name: String,
    line2name: String,
    line1data: Array,
    line2data: Array
})
const chartDom1 = ref()
let myChart = null

onMounted( () => {
    myChart = echarts.init(chartDom1.value);
    var option;
    option = {
    title: {
        text: props.title
    },
    legend: {},
    tooltip: {
        trigger: 'axis',
        axisPointer: {
        type: 'cross',
        label: {
            backgroundColor: '#283b56'
        }
        }
    },
    xAxis: [
    {
        type: 'time',
        splitLine: {
            show: false
        }
    },
    {
        type: 'time',
        splitLine: {
            show: false
        }
    }],
    yAxis: [
    {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
        show: false
        },
        min: 0,
        max: props.max
    },
    {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
        show: false
        }
    }],
    series: [
        {
        name: props.line1name,
        type: 'line',
        showSymbol: false,
        data: props.line1data
        },
        {
        name: props.line2name,
        type: 'line',
        showSymbol: false,
        data: props.line2data
        }
    ]
    };

    option && myChart.setOption(option);
})



setInterval(function () {

  myChart.setOption({
    series: [
      {
        data: props.line1data
      },
      {
        data: props.line2data
      }
    ]
  });
}, props.interval);


onUnmounted(() => {
    myChart.dispose()
})


</script>

<style scoped>

    .linechart {
    height: calc(50vh);
    /* min-height: 150px; */
    }
</style>