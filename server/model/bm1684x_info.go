package model

type TemLog struct {
	// gorm.Model
	Time             string `gorm:"type:DATETIME;comment:'记录时间'" json:"time"`
	BoardTemperature int    `gorm:"type:TINYINT;comment:'主板温度'" json:"boardTemperature"`
	ChipTemperature  int    `gorm:"type:TINYINT;comment:'芯片温度'" json:"chipTemperature"`
}

type PowerLog struct {
	// gorm.Model
	Time       string  `gorm:"comment:'记录时间'" json:"time"`
	BoardPower float32 `gorm:"type:real;comment:'整机功率'" json:"boardPower"`
	TPUPower   float32 `gorm:"type:real;comment:'TPU功率'" json:"TPUPower"`
}
