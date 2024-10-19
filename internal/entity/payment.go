package entity

import (
	"time"
)

type Payment struct {
	PaymentId     string    `json:"payment_id"`
	UserId        string    `json:"user_id"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date" `
	UpdatedDate   time.Time `json:"updated_date" `
}

// func (b *Payment) BeforeCreate(tx *gorm.DB) (err error) {
// 	b.PaymentId = uuid.New().String()
// 	return
// }
