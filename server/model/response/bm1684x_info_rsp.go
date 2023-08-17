package response

type Bm1684xTemRsp struct {
	Time             []string `json:"time"`
	BoardTemperature []int    `json:"boardTemperature"`
	ChipTemperature  []int    `json:"chipTemperature"`
}

type Bm1684xPowerRsp struct {
	Time       []string  `json:"time"`
	BoardPower []float32 `json:"boardPower"`
	TPUPower   []float32 `json:"tpuPower"`
}
