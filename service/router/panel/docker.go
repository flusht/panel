package panel

import (
	"sun-panel/api/api_v1/middleware"
	"sun-panel/api/api_v1/panel"

	"github.com/gin-gonic/gin"
)

func InitDocker(router *gin.RouterGroup) {
	dockerApi := panel.Docker{}
	// Requires Login
	r := router.Group("", middleware.LoginInterceptor)
	{
		r.GET("panel/docker/list", dockerApi.List)
		r.POST("panel/docker/start", dockerApi.Start)
		r.POST("panel/docker/stop", dockerApi.Stop)
		r.POST("panel/docker/restart", dockerApi.Restart)
		r.GET("panel/docker/logs", dockerApi.Logs)
	}
}
