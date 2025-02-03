package controllers

import (
	"net/http"

	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/application"
	"github.com/gin-gonic/gin"
)

type TableMonitorController struct {
	useCase *application.TableMonitorUseCase
}

func NewTableMonitorController(useCase *application.TableMonitorUseCase) *TableMonitorController {
	return &TableMonitorController{useCase: useCase}
}

func (ctrl *TableMonitorController) Execute(c *gin.Context) {
	changes, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"table_changes": changes})
}