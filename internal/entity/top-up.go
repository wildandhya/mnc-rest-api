package entity

import (
	"time"
)

type TopUp struct {
	TopUpId       string    `json:"top_up_id"`
	UserId        string    `json:"user_id"`
	Amount        int64    `json:"amount"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date" `
	UpdatedDate   time.Time `json:"updated_date" `
}

// func (b *TopUp) BeforeCreate(tx *gorm.DB) (err error) {
// 	b.TopUpId = uuid.New().String()
// 	return
// }
