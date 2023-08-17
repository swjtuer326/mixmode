<template>
  <el-card class="card_item">
    <el-row :gutter="2">
    <el-col :span="6">
      <div class="grid-content" >
        <Gauge title="芯片温度" :value="panelData.temChip" max="100" value-format="{value} °C"/>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="grid-content" >
        <Gauge title="主板温度" :value="panelData.temBoard" max="100" value-format="{value} °C"/> 
      </div>
    </el-col>
    <el-col :span="6">
      <div class="grid-content" >
        <Gauge title="整机功率" :value="panelData.powerBoard" max="100" value-format="{value} W"/>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="grid-content" >
        <Gauge title="TPU功率" :value="panelData.powerTPU" max="100" value-format="{value} W"/>
      </div>
    </el-col>
  </el-row>
  </el-card>
  <br>
  <el-card>
    <DynamicLineChart title="温度曲线" line1name="芯片温度" line2name="主板温度" interval="10000" :line1data="temData.chip" :line2data="temData.board"/>
  </el-card>
  <br>
  <el-card>
    <DynamicLineChart title="功耗曲线" line1name="TPU功率" line2name="整机功率" interval="10000" :line1data="powerData.tpu" :line2data="powerData.board"/>
  </el-card>
  
  <!-- <TemperaturePanel :title="title" :value="tem"/> -->
</template>

<script setup>
import * as echarts from "echarts"
import { onMounted, onUnmounted, ref, reactive} from 'vue'
import Gauge from "../components/Gauge.vue";
import DynamicLineChart from "../components/DynamicLineChart.vue";
import { getTemData , getPowerData} from "@/api/bmInfo"

let title = ref("1")
const panelData  = reactive({
  temChip: 0,
  temBoard: 0,
  powerTPU: 0,
  powerBoard: 0
})

const temData = ref({
  board: [],
  chip: []
})
const powerData = ref({
  board: [],
  tpu: []
})

const reload = async(lastMinute,func) => {
    let tem = await getTemData({lastMinute: lastMinute})
    let power = await getPowerData({lastMinute: lastMinute})

    // console.log(tem.data)
    // console.log(power.data)
    // console.log(tem.data.chipTemperature[tem.data.chipTemperature.length-1])
    func(tem.data)
    func(power.data)

    panelData.powerBoard = power.data.boardPower[power.data.boardPower.length-1]
    panelData.powerTPU = power.data.tpuPower[power.data.tpuPower.length-1]
    panelData.temChip = tem.data.chipTemperature[tem.data.chipTemperature.length-1]
    panelData.temBoard = tem.data.boardTemperature[tem.data.boardTemperature.length-1]

}


function refactorData(data){
  if ("boardTemperature" in data) {
    for (var i = 0; i < data.boardTemperature.length; i++) {
      temData.value.board.push([data.time[i],data.boardTemperature[i]])
      temData.value.chip.push([data.time[i],data.chipTemperature[i]])
    }
  }
  else{
    for (var i = 0; i < data.boardPower.length; i++) {
      powerData.value.board.push([data.time[i],data.boardPower[i]])
      powerData.value.tpu.push([data.time[i],data.tpuPower[i]])
    }
  }
}

function updateData(data){
  if ("boardTemperature" in data) {
    temData.value.board.shift()
    temData.value.chip.shift()
    temData.value.board.push([data.time[0],data.boardTemperature[0]])
    temData.value.chip.push([data.time[0],data.chipTemperature[0]])
  }
  else{
    powerData.value.board.shift()
    powerData.value.tpu.shift()
    powerData.value.board.push([data.time[0],data.boardPower[0]])
    powerData.value.tpu.push([data.time[0],data.tpuPower[0]])
  }
}

reload(10, refactorData)
// console.log(tem)
// refactorData(tem.data)
// refactorData(power.data)
// console.log(temData.value)

setInterval(function () {
  panelData.value = (Math.random() * 100).toFixed(0)
  reload(-1, updateData)
  // reload(10, refactorData)
  // console.log(temData.value.chip)
}, 10000);

</script>

<style scoped>

  /* #zoom {
    height: calc(25vh);
  } */
  .card_item {
    /* width: calc(80vw); */
  }

.el-row {
  margin-top: 0px;
  margin-bottom: 0px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
  width: calc(20vw);
}

</style>