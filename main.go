package main

import (
	"GoAssignment2/config"
	"GoAssignment2/db"
	"GoAssignment2/routes"
	"log"
	"net/http"
)

func main() {
	psqlInfo := config.GetPsqlInfo()
	dbConn, err := db.Connect(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	dbGORMConn, err := db.ConnectAdvanced(psqlInfo)
	if err != nil {
		panic(err)
	}

	dbConn.CreateTable()
	dbGORMConn.CreateTables()
	router := routes.SetupRoutes(dbConn, dbGORMConn)

	log.Fatal(http.ListenAndServe(":8080", router))
}
