package main

import (
	"log"

	pmdependencies "github.com/M1keTrike/MonitorAPIGo/src/PricesMonitor/dependencies"
	tbdependencies "github.com/M1keTrike/MonitorAPIGo/src/TablesMonitor/dependencies"
	database "github.com/M1keTrike/MonitorAPIGo/src/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not establish database connection:", err)
	}
	defer db.Close()

	r := gin.Default()

	pmdeps := pmdependencies.NewPriceMonitorDependencies(db)
	pmdeps.Execute(r)

	tmdeps := tbdependencies.NewTableMonitorDependencies(db)
	tmdeps.Execute(r)
	
	r.Run(":8080")
}
