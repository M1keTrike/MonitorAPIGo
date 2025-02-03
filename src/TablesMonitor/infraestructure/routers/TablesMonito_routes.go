package routers

import (
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTableMonitorRoutes(router *gin.Engine, controller *controllers.TableMonitorController) {
	monitorGroup := router.Group("/table-monitor")
	{
		monitorGroup.GET("/changes", controller.Execute)
	}
}
