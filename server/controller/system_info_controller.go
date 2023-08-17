package controller

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/logic"
)

type SystemInfoController struct{}

func (m *SystemInfoController) GetSysInfo(c *gin.Context) {
	RunWoReq(c, func() (interface{}, interface{}) {
		return logic.SystemInfo.GetSysInfo(c)
	})
}
