package persistence

import (
	"database/sql"
	"log"

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

	query := `
        SELECT product_id, old_price, new_price, change_time 
        FROM price_changes_log 
        ORDER BY change_time DESC 
        LIMIT 50
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error fetching price changes:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var change entities.PriceMonitor
		err := rows.Scan(&change.ProductID, &change.OldPrice, &change.NewPrice, &change.ChangeTime)
		if err != nil {
			log.Println("Error scanning price changes:", err)
			continue
		}
		changes = append(changes, change)
	}

	return changes, nil
}
