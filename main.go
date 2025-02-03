package main

import (
	"log"

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

	
	
	r.Run(":8080")
}
