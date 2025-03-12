package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"akastra-mobile-api/src/app/bootstrap"
	"akastra-mobile-api/src/interface/routes"
)


func initLogger() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initLogger()
	deps := bootstrap.InitDependencies()

	router := routes.InitRouter(deps)
	port := os.Getenv("PORT")

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
