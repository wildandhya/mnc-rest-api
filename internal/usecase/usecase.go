package usecase

import (
	"context"

	"github.com/wildandhya/mnc-rest-api/internal/model"
)

type Usecase interface {
	Register(ctx context.Context, body model.RegisterRequest) model.Result
	Login(ctx context.Context, body model.LoginRequest) model.Result
	TopUp(ctx context.Context, body model.TopUpRequest) model.Result
	Payment(ctx context.Context, body model.PaymentRequest) model.Result
	Transfer(ctx context.Context, body model.TransferRequest) model.Result
	TransactionReport(ctx context.Context) model.Result
	UpdateProfile(ctx context.Context, body model.UpdateProfileRequest ) model.Result
}
