package repository

import (
	"context"

	"github.com/wildandhya/mnc-rest-api/internal/entity"
	"github.com/wildandhya/mnc-rest-api/internal/model"
)

type Repository interface {
	Save(ctx context.Context, user entity.User) <-chan model.Result
	GetByFields(ctx context.Context, filters map[string]interface{}) <-chan model.Result
	TopUp(data entity.TopUp) <-chan model.Result
	Payment(data entity.Payment) <-chan model.Result
	Transfer(data entity.Transfer) <-chan model.Result
	GetTransaction(ctx context.Context) <-chan model.Result
	UpdateUser(ctx context.Context, params map[string]interface{}, body map[string]interface{}) <-chan model.Result
}
