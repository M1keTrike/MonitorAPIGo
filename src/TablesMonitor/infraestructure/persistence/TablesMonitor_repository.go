package persistence

import (
	"database/sql"

	"github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/domain/entities"
)

type TableMonitorRepository struct {
	DB *sql.DB
}

func NewTableMonitorRepository(db *sql.DB) *TableMonitorRepository {
	return &TableMonitorRepository{DB: db}
}

func (r *TableMonitorRepository) GetTableChanges() ([]entities.TableMonitor, error) {
	var changes []entities.TableMonitor
	return changes, nil
}