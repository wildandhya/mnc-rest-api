package model

type RegisterRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pin         string `json:"pin" validate:"required,len=6"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pin         string `json:"pin" validate:"required,len=6"`
}

type TopUpRequest struct {
	Amount int64 `json:"amount" validate:"required"`
}

type PaymentRequest struct {
	Amount  int64  `json:"amount" validate:"required"`
	Remarks string `json:"remarks" validate:"required"`
}

type TransferRequest struct {
	Amount     int64  `json:"amount" validate:"required"`
	TargetUser string `json:"target_user" validate:"required"`
	Remarks    string `json:"remarks" validate:"required"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Address   string `json:"address" `
}
