package controllers

import (
	"net/http"

	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/application"
	"github.com/gin-gonic/gin"
)

type PriceMonitorController struct {
	useCase *application.PriceMonitorUseCase
}

func NewPriceMonitorController(useCase *application.PriceMonitorUseCase) *PriceMonitorController {
	return &PriceMonitorController{useCase: useCase}
}

func (ctrl *PriceMonitorController) Execute(c *gin.Context) {
	changes, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"price_changes": changes})
}