package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subhrapaladhi/mongoapi/router"
)

func main() {
	fmt.Println("mongo api")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000")
}
