<template>
      <div ref="chartDom1" class="linechart"></div>
</template>
    
<script setup>
import * as echarts from "echarts"
import { onMounted, onUnmounted, ref} from 'vue'

const props = defineProps({
    interval: Number,
    title: String,
    line1name: String,
    line2name: String,
    line1data: Object,
    line2data: Object
})
const chartDom1 = ref()
let myChart = null


function randomData() {
  now = new Date(+now+oneDay);
  value = value + Math.random() * 21 - 10;
  return {
    name: now.toString(),
    value: [
      [now.getFullYear(), now.getMonth(), now.getDate()].join('-') + " " + [now.getHours(), now.getMinutes(), now.getSeconds()].join(':'),
      Math.round(value)
    ]
  };
}
let temChip = [];
let temBoard = [];
let now = new Date(1997, 9, 3);
let oneDay = 36 * 1000;
let value = Math.random() * 1000;
for (var i = 0; i < 1000; i++) {
  temBoard.push(randomData());
  temChip.push(randomData());
}

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
        }
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
        data: temBoard
        },
        {
        name: props.line2name,
        type: 'line',
        showSymbol: false,
        data: temChip
        }
    ]
    };

    option && myChart.setOption(option);
})



setInterval(function () {
  for (var i = 0; i < 5; i++) {
    temBoard.shift();
    temChip.shift();
    temBoard.push(randomData());
    temChip.push(randomData());
    // console.log(randomData())
  }
  myChart.setOption({
    series: [
      {
        data: temBoard
      },
      {
        data: temChip
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