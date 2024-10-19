package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wildandhya/mnc-rest-api/internal/entity"
	"github.com/wildandhya/mnc-rest-api/internal/model"
	"github.com/wildandhya/mnc-rest-api/internal/repository"
	"github.com/wildandhya/mnc-rest-api/pkg/helper"
)

type UsecaseImpl struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &UsecaseImpl{
		repo: repo,
	}
}

func (uc *UsecaseImpl) Register(ctx context.Context, body model.RegisterRequest) model.Result {
	user := entity.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		PhoneNumber: body.PhoneNumber,
		Pin:         body.Pin,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	// check if phone_number exists
	filters := map[string]interface{}{
		"phone_number": body.PhoneNumber,
	}
	exists := <-uc.repo.GetByFields(ctx, filters)
	if exists.Data != nil {
		return model.Result{
			Error: errors.New("phone number already registered"),
		}
	}

	result := <-uc.repo.Save(ctx, user)
	if exists.Error == nil {
		return model.Result{
			Error: result.Error,
		}
	}

	return model.Result{
		Data: user,
	}
}

func (uc *UsecaseImpl) Login(ctx context.Context, body model.LoginRequest) model.Result {
	filters := map[string]interface{}{
		"phone_number": body.PhoneNumber,
		"pin":          body.Pin,
	}
	result := <-uc.repo.GetByFields(ctx, filters)

	if result.Error != nil {
		return model.Result{
			Error: errors.New("phone number and pin doesn`t match"),
		}
	}

	user := result.Data.(entity.User)
	accessToken, err := helper.GenerateToken(user.UserID, time.Hour*12)
	if err != nil {
		return model.Result{
			Error: errors.New("internal server error"),
		}
	}

	refreshToken, err := helper.GenerateToken(user.UserID, time.Hour*24)
	if err != nil {
		return model.Result{
			Error: errors.New("internal server error"),
		}
	}
	response := model.LoginResponse{
		AccessToken:   accessToken,
		ResponseToken: refreshToken,
	}
	return model.Result{
		Data: response,
	}
}
func (uc *UsecaseImpl) TopUp(ctx context.Context, body model.TopUpRequest) model.Result {
	userId, ok := ctx.Value("user").(string)
	if !ok {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	getUser := <-uc.repo.GetByFields(ctx, map[string]interface{}{
		"user_id": userId,
	})
	if getUser.Error != nil {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	user := getUser.Data.(entity.User)

	doc := entity.TopUp{
		TopUpId:       uuid.New().String(),
		UserId:        userId,
		Amount:        body.Amount,
		BalanceBefore: user.Balance,
		BalanceAfter:  user.Balance + body.Amount,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	result := <-uc.repo.TopUp(doc)
	if result.Error != nil {
		return model.Result{
			Error: result.Error,
		}
	}

	return model.Result{
		Data: doc,
	}
}
func (uc *UsecaseImpl) Payment(ctx context.Context, body model.PaymentRequest) model.Result {
	userId, ok := ctx.Value("user").(string)
	if !ok {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	getUser := <-uc.repo.GetByFields(ctx, map[string]interface{}{
		"user_id": userId,
	})
	if getUser.Error != nil {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	user := getUser.Data.(entity.User)

	if user.Balance < body.Amount {
		return model.Result{
			Error: errors.New("balance is not enough"),
		}
	}

	doc := entity.Payment{
		PaymentId:     uuid.New().String(),
		UserId:        userId,
		Amount:        body.Amount,
		Remarks:       body.Remarks,
		BalanceBefore: user.Balance,
		BalanceAfter:  user.Balance - body.Amount,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	result := <-uc.repo.Payment(doc)
	if result.Error != nil {
		return model.Result{
			Error: result.Error,
		}
	}

	return model.Result{
		Data: doc,
	}
}
func (uc *UsecaseImpl) Transfer(ctx context.Context, body model.TransferRequest) model.Result {
	userId, ok := ctx.Value("user").(string)
	if !ok {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	getUser := <-uc.repo.GetByFields(ctx, map[string]interface{}{
		"user_id": userId,
	})
	if getUser.Error != nil {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	user := getUser.Data.(entity.User)

	if user.Balance < body.Amount {
		return model.Result{
			Error: errors.New("balance is not enough"),
		}
	}

	doc := entity.Transfer{
		TransferId:    uuid.New().String(),
		TargetUserId:  body.TargetUser,
		UserId:        userId,
		TransactionType: "DEBIT",
		Amount:        body.Amount,
		Remarks:       body.Remarks,
		BalanceBefore: user.Balance,
		BalanceAfter:  user.Balance - body.Amount,
		CreatedDate:   time.Now(),
		UpdatedDate:   time.Now(),
	}

	result := <-uc.repo.Transfer(doc)
	if result.Error != nil {
		return model.Result{
			Error: result.Error,
		}
	}

	return model.Result{
		Data: doc,
	}
}

func (uc *UsecaseImpl) TransactionReport(ctx context.Context) model.Result {
	result := <- uc.repo.GetTransaction(ctx)
	if result.Error != nil {
		return model.Result{
			Error: result.Error,
		}
	}

	data := result.Data.([]entity.Transfer)

	response := make([]model.TransactionResponse, 0)
	for _, val := range data{
		response = append(response, model.TransactionResponse{
			TransferId: val.TransferId,
			UserId: val.UserId,
			TransactionType: val.TransactionType,
			BalanceBefore: val.BalanceBefore,
			BalanceAfter: val.BalanceAfter,
			Status: val.Status,
			Amount: val.Amount,
			CreatedDate: val.CreatedDate,
			Remarks: val.Remarks,
		})
	}

	return model.Result{
		Data: response,
	}
}

func (uc *UsecaseImpl) UpdateProfile(ctx context.Context, body model.UpdateProfileRequest) model.Result {
	userId, ok := ctx.Value("user").(string)
	if !ok {
		return model.Result{
			Error: errors.New("user not found"),
		}
	}

	params := map[string]interface{}{
		"user_id": userId,
	}

	doc := make(map[string]interface{})

	if body.FirstName != ""{
		doc["first_name"] = body.FirstName
	}

	if body.LastName != ""{
		doc["last_name"] = body.LastName
	}

	if body.Address != ""{
		doc["address"] = body.Address
	}

	result := <- uc.repo.UpdateUser(ctx, params, doc)
	if result.Error != nil {
		return model.Result{
			Error: result.Error,
		}
	}

	return model.Result{
		Data: body,
	}
}
