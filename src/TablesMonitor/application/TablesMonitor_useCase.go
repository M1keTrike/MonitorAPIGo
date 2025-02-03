package application

import (
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/domain/entities"
	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/infraestructure/persistence"
)

type TableMonitorUseCase struct {
	repo persistence.TableMonitorRepository
}

func NewTableMonitorUseCase(repo persistence.TableMonitorRepository) *TableMonitorUseCase {
	return &TableMonitorUseCase{repo: repo}
}

func (uc *TableMonitorUseCase) Execute() ([]entities.TableMonitor, error) {
	return uc.repo.GetTableChanges()
}