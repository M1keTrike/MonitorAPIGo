package persistence

import (
	"database/sql"
	"log"

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

	query := `
        SELECT event_time, table_name, operation, details 
        FROM table_changes_log 
        ORDER BY event_time DESC 
        LIMIT 50
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error fetching table changes:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var change entities.TableMonitor
		err := rows.Scan(&change.EventTime, &change.Table, &change.Action, &change.Details)
		if err != nil {
			log.Println("Error scanning table changes:", err)
			continue
		}
		changes = append(changes, change)
	}

	return changes, nil
}
