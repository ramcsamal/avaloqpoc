package main

import (
	"avaloqpoc/internal/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve OpenAPI Document
	r.HandleFunc("/api/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/openapi.yaml")
	})

	// Serve Swagger UI
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./third_party/swaggerui/"))))
	r.HandleFunc("/api/v1/execute", api.ExecuteCommandHandler).Methods("GET")

	// Redirect root to Swagger UI
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swaggerui/index.html?url=/api/openapi.yaml", http.StatusSeeOther)
	})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
