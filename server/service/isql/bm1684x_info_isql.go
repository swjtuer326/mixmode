package isql

import (
	"fmt"
	"time"

	"sophgo.com/mixmode/model"
	"sophgo.com/mixmode/model/request"
	"sophgo.com/mixmode/utils/common"
)

type Bm1684xInfoService struct{}

// 获取温度列表
func (s Bm1684xInfoService) ListTem(req *request.Bm1684xTemReq) ([]*model.TemLog, error) {
	var list []*model.TemLog
	m, _ := time.ParseDuration(fmt.Sprintf("-%dm", req.LastMinute))
	// fmt.Println(time.Now())
	Time := time.Now().Add(m).Format("2006.01.02 15:04:05")
	db := common.DB.Where("time >= ?", Time)
	err := db.Find(&list).Error
	return list, err
}

func (s Bm1684xInfoService) AddTemItem(item *model.TemLog) error {
	err := common.DB.Create(item).Error
	return err
}

// 获取最新的功耗信息
func (s Bm1684xInfoService) GetLatestTem() (data *model.TemLog, err error) {
	err = common.DB.Last(&data).Error
	return
}

// 获取功耗列表
func (s Bm1684xInfoService) ListPower(req *request.Bm1684xPowerReq) ([]*model.PowerLog, error) {
	var list []*model.PowerLog
	m, _ := time.ParseDuration(fmt.Sprintf("-%dm", req.LastMinute))
	// fmt.Println(time.Now())
	Time := time.Now().Add(m).Format("2006.01.02 15:04:05")
	db := common.DB.Where("time >= ?", Time)
	err := db.Find(&list).Error
	return list, err
}

func (s Bm1684xInfoService) AddPowerItem(item *model.PowerLog) error {
	err := common.DB.Create(item).Error
	return err
}

// 获取最新的温度信息
func (s Bm1684xInfoService) GetLatestPower() (data *model.PowerLog, err error) {
	err = common.DB.Last(&data).Error
	return
}
