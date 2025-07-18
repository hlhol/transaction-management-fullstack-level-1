package model

type Transaction struct {
	TransactionID string `json:"transaction_id"` // UUID
	AccountID     string `json:"account_id"`     // UUID
	Amount        int64  `json:"amount"`         // integer
	CreatedAt     string `json:"created_at"`     // RFC3339 datetime string
	Balance       int64  `json:"balance"`        // Not in API spec, but for your logic
}

type TransactionRequest struct {
	AccountID string `json:"account_id"` // UUID
	Amount    int64  `json:"amount"`     // integer
}


type Account struct {
	AccountID string `json:"account_id"`
	Balance   int64  `json:"balance"`
}

