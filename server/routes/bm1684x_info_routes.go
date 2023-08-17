package routes

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/controller"
)

// 注册基础路由
func InitBm1684xInfoRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/bm1684x")
	{
		base.POST("TemData", controller.Bm1684xInfo.GetTemData)
		base.POST("PowerData", controller.Bm1684xInfo.GetPowerData)
	}
	return r
}
