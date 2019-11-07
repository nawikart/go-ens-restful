package main

import (
	"fmt"
	"log"
	"net/http"

	"./cors"
	"./routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.Ens(r)

	fmt.Println("Api-dev running at HTTP port :8081")
	log.Fatal(http.ListenAndServe(":8081", cors.Handler()(r)))
}
