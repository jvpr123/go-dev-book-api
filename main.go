package main

import (
	"dev-book-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API running...")

	router := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":9000", router))
}
