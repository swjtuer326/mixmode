package logic

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/utils/common"
	"sophgo.com/mixmode/utils/tools"
)

type SystemInfoLogic struct{}

func (l SystemInfoLogic) GetSysInfo(c *gin.Context) (data interface{}, rspError interface{}) {
	var s tools.Server
	var err error
	s.Os = tools.InitOS()
	if s.Cpu, err = tools.InitCPU(); err != nil {
		common.Log.Errorf("InitCPU ERROR: %s\n", err)
		return s, err
	}
	if s.Disk, err = tools.InitDisk(); err != nil {
		common.Log.Errorf("InitDisk ERROR: %s\n", err)
		return s, err
	}
	if s.Ram, err = tools.InitRAM(); err != nil {
		common.Log.Errorf("InitRAM ERROR: %s\n", err)
		return s, err
	}
	return s, nil
}
