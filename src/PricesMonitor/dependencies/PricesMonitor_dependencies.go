package dependencies

import (
	"database/sql"

	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/application"
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/infraestructure/controllers"
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/infraestructure/persistence"
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/infraestructure/routers"
	"github.com/gin-gonic/gin"
)

type PriceMonitorDependencies struct {
	DB *sql.DB
}

func NewPriceMonitorDependencies(db *sql.DB) *PriceMonitorDependencies {
	return &PriceMonitorDependencies{DB: db}
}

func (d *PriceMonitorDependencies) Execute(r *gin.Engine) {
	priceRepository := persistence.NewPriceMonitorRepository(d.DB)
	priceMonitorUseCase := application.NewPriceMonitorUseCase(*priceRepository)
	priceMonitorController := controllers.NewPriceMonitorController(priceMonitorUseCase)

	routers.RegisterPriceMonitorRoutes(r, priceMonitorController)

}
