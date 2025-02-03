package domain

import "github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/domain/entities"

type IPriceMonitorRepository interface {
	GetPriceChanges() ([]entities.PriceMonitor, error)
}