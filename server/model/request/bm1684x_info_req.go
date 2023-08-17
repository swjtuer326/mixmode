package request

type Bm1684xTemReq struct {
	LastMinute int `json:"lastMinute" validate:"required"`
}

type Bm1684xPowerReq struct {
	LastMinute int `json:"lastMinute" validate:"required"`
}
