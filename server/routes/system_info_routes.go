package routes

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/controller"
)

// 注册基础路由
func InitSystemInfoRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/system")
	{
		base.POST("SystemInfo", controller.SystemInfo.GetSysInfo)
	}
	return r
}
