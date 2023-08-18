<template>
    <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span >测试命令</span>
      </div>
    </template>
    <el-button type="primary" plane @click="subTestGDMA" >test_gdma</el-button>
    <el-button type="primary" plane @click="subTestCDMA" >test_cdma</el-button>
  </el-card>
    <div      
      v-loading="loading"
      element-loading-text="Loading..."
      :element-loading-spinner="svg"
      element-loading-background="rgba(122, 122, 122, 0.8)"
      :data="tableData"
      style="width: 100%">
      <p v-text="str" class="res"></p>
    </div>

    
  </template>
  
  <script  setup>
  import { ref } from 'vue'
  import {testCDMA, testGDMA} from '@/api/sophonTest'
  let str = ref("")
  const loading = ref(false)


  const subTestCDMA = async() => {
    str.value=""
    loading.value = true
    const res = await testCDMA()
    str.value = res.data.res
    loading.value =false
  }

  const subTestGDMA = async() => {
    str.value=""
    loading.value = true
    const res = await testGDMA()
    str.value = res.data.res
    loading.value =false
  }


  </script>
  <style>
    .example-showcase .el-loading-mask {
    z-index: 9;
    }
    .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    }
  .res{
    white-space:pre-wrap;
  }
  </style>