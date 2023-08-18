package sophon

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"time"

	"sophgo.com/mixmode/model"
	"sophgo.com/mixmode/service/isql"
	"sophgo.com/mixmode/utils/common"
)

func GetBoardInfo() {
	boardTem, chipTem := GetTemperature()
	isql.Bm1684xInfo.AddTemItem(&model.TemLog{
		Time:             time.Now().Format("2006.01.02 15:04:05"),
		BoardTemperature: boardTem,
		ChipTemperature:  chipTem,
	})
	// isql.Bm1684xInfo.AddTemItem(&model.TemLog{
	// 	Time:             time.Now().Format("2006.01.02 15:04:05"),
	// 	BoardTemperature: rand.Intn(100),
	// 	ChipTemperature:  rand.Intn(100),
	// })
	isql.Bm1684xInfo.AddPowerItem(&model.PowerLog{
		Time:       time.Now().Format("2006.01.02 15:04:05"),
		BoardPower: rand.Float32() * 50,
		TPUPower:   rand.Float32() * 50,
	})
}

func GetTemperature() (int, int) {
	boardTem := 0
	chiptem := 0
	out, err := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()
	fmt.Println(out)
	if err != nil {
		common.Log.Fatal(err)
	} else {
		if val, err := strconv.Atoi(string(out[0 : len(out)-1])); err == nil {
			chiptem = val / 1000
		}
	}

	out, err = exec.Command("cat", "/sys/class/thermal/thermal_zone1/temp").Output()
	if err != nil {
		common.Log.Fatal(err)
	} else {
		if val, err := strconv.Atoi(string(out[0 : len(out)-1])); err == nil {
			boardTem = val / 1000
		}
	}
	fmt.Println(boardTem, chiptem)
	return boardTem, chiptem
	// fmt.Printf("The date is %s\n", out)
	// tpuVol := 0.0
	// cpuVol := 0.0
	// tpuCur := 0.0
	// cpuCur := 0.0
	// str := `I2C port 0, addr 0x50, type 0x3, reg 0x0, value 0x0
	// ISL68127 revision 0x33
	// ISL68127 switch to output 0, ret=0
	// ISL68127 output voltage: 764mV
	// ISL68127 output current: 6300mA
	// ISL68127 temperature 1: 54C
	// ISL68127 output power: 4W
	// ISL68127 switch to output 1, ret=0
	// ISL68127 output voltage: 928mV
	// ISL68127 output current: 4200mA
	// ISL68127 temperature 1: 54C
	// ISL68127 output power: 3W
	// firmware load succeed`
	// numRe := regexp.MustCompile("[0-9]+")
	// volRe := regexp.MustCompile(`voltage:.*`)
	// curRe := regexp.MustCompile(`current:.*`)
	// if vols := volRe.FindAllString(str, -1); vols != nil && len(vols) == 2 {
	// 	if tmp := numRe.FindAllString(vols[0], -1); tmp != nil {
	// 		tpuVol, _ = strconv.ParseFloat(tmp[0], 64)
	// 	}
	// 	if tmp := numRe.FindAllString(vols[1], -1); tmp != nil {
	// 		cpuVol, _ = strconv.ParseFloat(tmp[0], 64)
	// 	}
	// }
	// if curs := curRe.FindAllString(str, -1); curs != nil && len(curs) == 2 {
	// 	if tmp := numRe.FindAllString(curs[0], -1); tmp != nil {
	// 		tpuCur, _ = strconv.ParseFloat(tmp[0], 64)
	// 	}
	// 	if tmp := numRe.FindAllString(curs[1], -1); tmp != nil {
	// 		cpuCur, _ = strconv.ParseFloat(tmp[0], 64)
	// 	}
	// }
	// fmt.Println(tpuVol/1000, cpuVol/1000, tpuCur/1000, cpuCur/1000)
}
