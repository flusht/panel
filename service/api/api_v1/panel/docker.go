package panel

import (
	"sun-panel/api/api_v1/common/apiReturn"
	"sun-panel/lib/docker"

	"github.com/gin-gonic/gin"
)

type Docker struct{}

// List 获取容器列表
func (d *Docker) List(c *gin.Context) {
	list, err := docker.ListContainers()
	if err != nil {
		apiReturn.Error(c, "获取容器列表失败: "+err.Error())
		return
	}
	apiReturn.SuccessData(c, list)
}

// Start 启动容器
func (d *Docker) Start(c *gin.Context) {
	type Req struct {
		ID string `json:"id" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		apiReturn.ErrorParamFomat(c, err.Error())
		return
	}

	if err := docker.StartContainer(req.ID); err != nil {
		apiReturn.Error(c, "启动容器失败: "+err.Error())
		return
	}
	apiReturn.Success(c)
}

// Stop 停止容器
func (d *Docker) Stop(c *gin.Context) {
	type Req struct {
		ID string `json:"id" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		apiReturn.ErrorParamFomat(c, err.Error())
		return
	}

	if err := docker.StopContainer(req.ID); err != nil {
		apiReturn.Error(c, "停止容器失败: "+err.Error())
		return
	}
	apiReturn.Success(c)
}

// Restart 重启容器
func (d *Docker) Restart(c *gin.Context) {
	type Req struct {
		ID string `json:"id" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		apiReturn.ErrorParamFomat(c, err.Error())
		return
	}

	if err := docker.RestartContainer(req.ID); err != nil {
		apiReturn.Error(c, "重启容器失败: "+err.Error())
		return
	}
	apiReturn.Success(c)
}

// Logs 获取日志
func (d *Docker) Logs(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		apiReturn.ErrorParamFomat(c, "id is required")
		return
	}

	logs, err := docker.GetContainerLogs(id)
	if err != nil {
		apiReturn.Error(c, "获取日志失败: "+err.Error())
		return
	}
	apiReturn.SuccessData(c, logs)
}
