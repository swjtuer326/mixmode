package routes

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/controller"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("ping", controller.Demo)
	}
	return r
}
