package entity

import "time"

type Transfer struct {
	TransferId      string    `json:"transfer_id"`
	UserId          string    `json:"user_id"`
	TargetUserId    string    `json:"target_user_id"`
	Status          string    `json:"status"`
	TransactionType string    `json:"transaction_type"`
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedDate     time.Time `json:"created_date" `
	UpdatedDate     time.Time `json:"updated_date" `
}
