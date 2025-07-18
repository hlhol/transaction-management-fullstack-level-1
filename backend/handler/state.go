package handler

import (
	"sync"
	"backend/model"
)


var (
	transactions    []model.Transaction
	accountBalances = make(map[string]int64)
	mu              sync.Mutex
)
