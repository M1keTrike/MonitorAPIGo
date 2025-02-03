package entities


type PriceMonitor struct {
	ProductID int     `json:"product_id"`
	OldPrice  float64 `json:"old_price"`
	NewPrice  float64 `json:"new_price"`
}