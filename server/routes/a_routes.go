package routes

import (
	"sophgo.com/mixmode/config"
	"sophgo.com/mixmode/utils/common"

	"github.com/gin-gonic/gin"
)

// 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	// r := gin.New()
	// r.Use(gin.Recovery())

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 注册路由
	InitBaseRoutes(apiGroup) // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	InitSystemInfoRoutes(apiGroup)
	InitBm1684xInfoRoutes(apiGroup)
	InitSophonTestRoutes(apiGroup)

	common.Log.Info("初始化路由完成！")
	return r
}
