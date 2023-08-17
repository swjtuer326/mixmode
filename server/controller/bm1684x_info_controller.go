package controller

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/logic"
	"sophgo.com/mixmode/model/request"
)

type Bm1684xInfoController struct{}

func (m *Bm1684xInfoController) GetTemData(c *gin.Context) {
	// common.Log.Info(c.Request)
	req := new(request.Bm1684xTemReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Bm1684xInfo.GetTemData(c, req)
	})
}

func (m *Bm1684xInfoController) GetPowerData(c *gin.Context) {
	req := new(request.Bm1684xPowerReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Bm1684xInfo.GetPowerData(c, req)
	})
}
