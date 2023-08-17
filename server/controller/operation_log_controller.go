package controller

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/logic"
	"sophgo.com/mixmode/model/request"
)

type OperationLogController struct{}

// List 记录列表
func (m *OperationLogController) List(c *gin.Context) {
	req := new(request.OperationLogListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.OperationLog.List(c, req)
	})
}

// Delete 删除记录
func (m *OperationLogController) Delete(c *gin.Context) {
	req := new(request.OperationLogDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.OperationLog.Delete(c, req)
	})
}
