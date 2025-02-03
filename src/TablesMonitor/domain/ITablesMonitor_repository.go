package domain

import "github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/domain/entities"

type ITableMonitorRepository interface {
	GetTableChanges() ([]entities.TableMonitor, error)
}