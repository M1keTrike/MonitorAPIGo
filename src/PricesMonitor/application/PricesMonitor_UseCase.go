package application

import (
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/domain/entities"
	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/infraestructure/persistence"
)

type PriceMonitorUseCase struct {
	repo persistence.PriceMonitorRepository
}

func NewPriceMonitorUseCase(repo persistence.PriceMonitorRepository) *PriceMonitorUseCase {
	return &PriceMonitorUseCase{repo: repo}
}

func (uc *PriceMonitorUseCase) Execute() ([]entities.PriceMonitor, error) {
	return uc.repo.GetPriceChanges()
}