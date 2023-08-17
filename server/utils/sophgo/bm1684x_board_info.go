package sophgo

import (
	"math/rand"
	"time"

	"sophgo.com/mixmode/model"
	"sophgo.com/mixmode/service/isql"
)

func GetBoardInfo() {
	isql.Bm1684xInfo.AddTemItem(&model.TemLog{
		Time:             time.Now().Format("2006.01.02 15:04:05"),
		BoardTemperature: rand.Intn(100),
		ChipTemperature:  rand.Intn(100),
	})
	isql.Bm1684xInfo.AddPowerItem(&model.PowerLog{
		Time:       time.Now().Format("2006.01.02 15:04:05"),
		BoardPower: rand.Float32() * 50,
		TPUPower:   rand.Float32() * 50,
	})
}
