<template>
      <div ref="chartDom" class="gauge"></div>
</template>
    
<script setup>
import * as echarts from "echarts"
import { onMounted, onUnmounted, ref} from 'vue'

// 使用 <script setup>
const props = defineProps({
  title: String,
  valueFormat: String,
  value: Number,
  max: Number
})

let myChart = null
const chartDom = ref()
onMounted(() => {
    // var chartDom = document.getElementById('zoom');
    myChart = echarts.init(chartDom.value);
    var option;

    option = {
    series: [
    {
        type: 'gauge',
        radius: '115%',
        center: ['50%','60%'],
        min: 0,
        max: props.max,
        axisLine: {
        lineStyle: {
            width: 12,
            color: [
            [0.3, '#67e0e3'],
            [0.7, '#37a2da'],
            [1, '#fd666d']
            ]
        }
        },
        pointer: {
        itemStyle: {
            color: 'auto'
        }
        },
        axisTick: {
        distance: -10,
        length: 8,
        lineStyle: {
            color: '#fff',
            width: 2
        }
        },
        splitLine: {
        distance: -20,
        length: 20,
        lineStyle: {
            color: '#fff',
            width: 3
        }
        },
        axisLabel: {
        color: 'inherit',
        distance: 18,
        fontSize: 14
        },
        detail: {
        valueAnimation: true,
        formatter: props.valueFormat,
        color: 'inherit',
        fontSize: 16
        },
        title: {
            show: true,
            offsetCenter: [0, '-20%']
        },
        data: [
        {
            value: props.value,
            name: props.title
        }
        ]
    }
    ]
};

option && myChart.setOption(option);
})
setInterval(function () {
    myChart.setOption({
    series: [
        {
        data: [
            {
                name: props.title,
                value: props.value
            }
        ]
        }
    ]
    });
}, 2000);
onUnmounted(() => {
    myChart.dispose()
})
</script>

<style scoped>

    .gauge {
    height: calc(25vh);
    /* min-height: 150px; */
    }
</style>