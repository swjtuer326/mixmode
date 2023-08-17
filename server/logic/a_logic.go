package logic

import (
	"fmt"

	"sophgo.com/mixmode/utils/tools"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))
	OperationLog = &OperationLogLogic{}
	SystemInfo   = &SystemInfoLogic{}
	Bm1684xInfo  = &Bm1684xInfoLogic{}
)
