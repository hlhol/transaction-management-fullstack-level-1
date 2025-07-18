package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"backend/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// valdition for eader
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Specified content type not allowed", http.StatusUnsupportedMediaType)
		return
	}

	type Input struct {
		AccountID string `json:"account_id"`
		Amount    int64  `json:"amount"`
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Mandatory body parameters missing or have incorrect type", http.StatusBadRequest)
		return
	}

	// validation
	if input.AccountID == "" {
		http.Error(w, "account_id is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// auto create account if it doesn't exist
	if _, ok := accountBalances[input.AccountID]; !ok {
		accountBalances[input.AccountID] = 0
	}

	// create transaction
	t := model.Transaction{
		TransactionID: uuid.New().String(),
		AccountID:     input.AccountID,
		Amount:        input.Amount,
		CreatedAt:     time.Now().UTC().Format(time.RFC3339),
	}

	// update account balance
	accountBalances[t.AccountID] += t.Amount
	t.Balance = accountBalances[t.AccountID]

	transactions = append([]model.Transaction{t}, transactions...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]

	if transactionID == "" {
		http.Error(w, "transaction_id is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range transactions {
		if t.TransactionID == transactionID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}