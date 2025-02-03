package entities

type TableMonitor struct {
	Table     string `json:"table"`
	Action    string `json:"action"`
	Details   string `json:"details"`
	EventTime string `json:"event-time"`
}