package model

import "time"

type Result struct {
	Data  interface{}
	Error error
}

type Response struct {
	Status  string      `json:"status"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

type LoginResponse struct {
	AccessToken   string `json:"access_token"`
	ResponseToken string `json:"response_token"`
}

type TransactionResponse struct {
	TransferId      string    `json:"transfer_id"`
	UserId          string    `json:"user_id"`
	Status          string    `json:"status"`
	TransactionType string    `json:"transaction_type"`
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedDate     time.Time `json:"created_date" `
	UpdatedDate     time.Time `json:"updated_date" `
}
