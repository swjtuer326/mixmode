package test

import (
	"fmt"
	"testing"
	"time"

	"sophgo.com/mixmode/config"
	"sophgo.com/mixmode/model"
	"sophgo.com/mixmode/model/request"
	"sophgo.com/mixmode/service/isql"
	"sophgo.com/mixmode/utils/common"
)

func Test(t *testing.T) {
	config.InitConfig()
	common.InitDB()

	for i := 1; i < 3000; i++ {
		m, _ := time.ParseDuration(fmt.Sprintf("%dm", -i))
		// fmt.Println(time.Now())
		Time := time.Now().Add(m).Format("2006.01.02 15:04:05")
		i := model.TemLog{Time: Time, BoardTemperature: 1, ChipTemperature: 1}
		// t.Logf("%v\n", i)
		isql.Bm1684xInfo.AddTemItem(&i)
	}
	if data, err := isql.Bm1684xInfo.ListTem(&request.Bm1684xTemReq{LastMinute: 20}); err == nil {
		for _, value := range data {
			t.Log(value)
		}
	}
	data1, _ := isql.Bm1684xInfo.GetLatestTem()
	t.Log(data1)
}
