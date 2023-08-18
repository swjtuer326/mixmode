package routes

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/controller"
)

// 注册基础路由
func InitSophonTestRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/sophon")
	{
		base.POST("TestCDMA", controller.SophonTest.TestCDMA)
		base.POST("TestGDMA", controller.SophonTest.TestGDMA)
	}
	return r
}
