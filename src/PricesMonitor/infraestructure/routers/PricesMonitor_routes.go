package routers

import (
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPriceMonitorRoutes(router *gin.Engine, controller *controllers.PriceMonitorController) {
	monitorGroup := router.Group("/price-monitor")
	{
		monitorGroup.GET("/changes", controller.Execute)
	}
}