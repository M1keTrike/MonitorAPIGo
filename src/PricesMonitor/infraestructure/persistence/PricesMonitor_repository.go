package persistence

import (
	"database/sql"

	"github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/domain/entities"
)

	

type PriceMonitorRepository struct {
	DB *sql.DB
}

func NewPriceMonitorRepository(db *sql.DB) *PriceMonitorRepository {
	return &PriceMonitorRepository{DB: db}
}

func (r *PriceMonitorRepository) GetPriceChanges() ([]entities.PriceMonitor, error) {
	var changes []entities.PriceMonitor
	return changes, nil
}