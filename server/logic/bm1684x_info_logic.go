package logic

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/model/request"
	"sophgo.com/mixmode/model/response"
	"sophgo.com/mixmode/service/isql"
	"sophgo.com/mixmode/utils/tools"
)

type Bm1684xInfoLogic struct{}

func (l Bm1684xInfoLogic) GetTemData(c *gin.Context, req interface{}) (interface{}, interface{}) {
	r, ok := req.(*request.Bm1684xTemReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	if r.LastMinute == -1 {
		data, err := isql.Bm1684xInfo.GetLatestTem()
		if err != nil || data == nil {
			return data, tools.NewSqliteError(fmt.Errorf("获取数据失败: %s", err.Error()))
		}

		rspData := response.Bm1684xTemRsp{
			Time:             []string{strings.ReplaceAll(data.Time, ".", "-")},
			BoardTemperature: []int{data.BoardTemperature},
			ChipTemperature:  []int{data.ChipTemperature},
		}
		return rspData, nil
	} else {
		data, err := isql.Bm1684xInfo.ListTem(r)
		if err != nil || len(data) == 0 {
			return data, tools.NewSqliteError(fmt.Errorf("获取数据失败: %s", err.Error()))
		}
		rspData := response.Bm1684xTemRsp{
			Time:             make([]string, len(data)),
			BoardTemperature: make([]int, len(data)),
			ChipTemperature:  make([]int, len(data)),
		}

		for index, v := range data {
			rspData.Time[index] = strings.ReplaceAll(v.Time, ".", "-")
			rspData.BoardTemperature[index] = v.BoardTemperature
			rspData.ChipTemperature[index] = v.ChipTemperature
		}
		// fmt.Println(rspData)
		return rspData, nil
	}
}

func (l Bm1684xInfoLogic) GetPowerData(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.Bm1684xPowerReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if r.LastMinute == -1 {
		data, err := isql.Bm1684xInfo.GetLatestPower()
		if err != nil || data == nil {
			return data, tools.NewSqliteError(fmt.Errorf("获取数据失败: %s", err.Error()))
		}
		rspData := response.Bm1684xPowerRsp{
			Time:       []string{strings.ReplaceAll(data.Time, ".", "-")},
			BoardPower: []float32{data.BoardPower},
			TPUPower:   []float32{data.TPUPower},
		}
		return rspData, nil
	} else {
		data, err := isql.Bm1684xInfo.ListPower(r)
		if err != nil || len(data) == 0 {
			return data, tools.NewSqliteError(fmt.Errorf("获取数据失败: %s", err.Error()))
		}
		rspData := response.Bm1684xPowerRsp{
			Time:       make([]string, len(data)),
			BoardPower: make([]float32, len(data)),
			TPUPower:   make([]float32, len(data)),
		}
		for index, v := range data {
			rspData.Time[index] = strings.ReplaceAll(v.Time, ".", "-")
			rspData.BoardPower[index] = v.BoardPower
			rspData.TPUPower[index] = v.TPUPower
		}
		return rspData, nil
	}
}
