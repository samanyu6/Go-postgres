package main

import (
	"factly/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// All routers available
	r := routers.Router()
	fmt.Println("Starting server on port 8000..")

	log.Fatal(http.ListenAndServe(":8000", r))
}
