package main

import (
	"log"
	"net/http"

	"backend/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")

	r.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")
	r.HandleFunc("/transactions", handler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", handler.ListTransactions).Methods("GET")
	r.HandleFunc("/transactions/{transaction_id}", handler.GetTransactionByID).Methods("GET")
	r.HandleFunc("/accounts/{account_id}", handler.GetAccount).Methods("GET")

	corsRouter := corsMiddleware(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
