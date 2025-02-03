package dependencies

import (
	"database/sql"

	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/application"
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/infraestructure/controllers"
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/infraestructure/persistence"
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/infraestructure/routers"
	"github.com/gin-gonic/gin"
)

type TableMonitorDependencies struct {
	DB *sql.DB
}

func NewTableMonitorDependencies(db *sql.DB) *TableMonitorDependencies {
	return &TableMonitorDependencies{DB: db}
}

func (d *TableMonitorDependencies) Execute(r *gin.Engine) {
	tableRepository := persistence.NewTableMonitorRepository(d.DB)
	tableMonitorUseCase := application.NewTableMonitorUseCase(*tableRepository)
	tableMonitorController := controllers.NewTableMonitorController(tableMonitorUseCase)

	routers.RegisterTableMonitorRoutes(r, tableMonitorController)

}
