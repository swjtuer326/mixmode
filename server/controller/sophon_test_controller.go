package controller

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/logic"
)

type SophonTestController struct{}

func (m *SophonTestController) TestCDMA(c *gin.Context) {
	RunWoReq(c, func() (interface{}, interface{}) {
		return logic.SophonTest.TestCDMA(c)
	})
}

func (m *SophonTestController) TestGDMA(c *gin.Context) {
	RunWoReq(c, func() (interface{}, interface{}) {
		return logic.SophonTest.TestGDMA(c)
	})
}
