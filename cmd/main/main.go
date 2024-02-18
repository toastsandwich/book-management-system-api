// entry point
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/toastsandwich/book-management-system-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9101", nil))
}
