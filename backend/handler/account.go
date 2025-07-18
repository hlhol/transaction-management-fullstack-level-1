package handler

import (
	"encoding/json"
	"net/http"

	"backend/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountID := uuid.New().String()

	mu.Lock()
	accountBalances[accountID] = 0
	mu.Unlock()

	account := model.Account{
		AccountID: accountID,
		Balance:   0,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["account_id"]

	mu.Lock()
	defer mu.Unlock()

	balance, ok := accountBalances[accountID]
	if !ok {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	account := model.Account{
		AccountID: accountID,
		Balance:   balance,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}