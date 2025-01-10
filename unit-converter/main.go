package main

import (
	"net/http"

	config "github.com/alielmi98/Unit-Converter-Web-APP-Go-Implementation/configs"
	"github.com/alielmi98/Unit-Converter-Web-APP-Go-Implementation/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/length", handlers.LengthHandler).Methods("GET", "POST")
	r.HandleFunc("/weight", handlers.WeightHandler).Methods("GET", "POST")
	r.HandleFunc("/temperature", handlers.TemperatureHandler).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	http.ListenAndServe(config.GetPort(), r)
}
