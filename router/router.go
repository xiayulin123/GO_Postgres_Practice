package router

import (
	"GO_POSTGRESQL/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStockById).Methonds("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetStock).Methonds("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.Newstock).Methonds("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methonds("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methonds("DELETE", "OPTIONS")
}