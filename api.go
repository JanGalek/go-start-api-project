package main

import (
	"github.com/joho/godotenv"
	"jan-galek/api/src/libs/database"
	routes "jan-galek/api/src/router"
	"log"
	"net/http"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, con := database.GetConnection()
	database.Setup(con)

	router := routes.HandleRequest()
	log.Fatal(http.ListenAndServe(":8000", router))
}
