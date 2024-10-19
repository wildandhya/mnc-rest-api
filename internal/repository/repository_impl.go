package repository

import (
	"context"
	"fmt"

	"github.com/wildandhya/mnc-rest-api/internal/entity"
	"github.com/wildandhya/mnc-rest-api/internal/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

func (r *RepositoryImpl) GetByFields(ctx context.Context, filters map[string]interface{}) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		var user entity.User
		query := r.DB
		for field, value := range filters {
			query = query.Where(fmt.Sprintf("%s = ?", field), value)
		}

		result := query.First(&user)
		if result.Error != nil {
			output <- model.Result{
				Error: result.Error,
				Data:  nil,
			}
		}

		output <- model.Result{
			Data: user,
		}
	}()

	return output
}

func (r *RepositoryImpl) Save(ctx context.Context, data entity.User) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		result := r.DB.Create(&data)
		if result.Error != nil {
			output <- model.Result{
				Error: result.Error,
			}
		}

		output <- model.Result{
			Data: result.RowsAffected,
		}
	}()

	return output
}

func (r *RepositoryImpl) TopUp(data entity.TopUp) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		err := r.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&data).Error; err != nil {
				return err
			}

			if err := tx.Model(&entity.User{}).Where("user_id = ?", data.UserId).Update("balance", data.BalanceAfter).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			output <- model.Result{
				Error: err,
			}
		}

		output <- model.Result{
			Data: "ok",
		}
	}()

	return output
}

func (r *RepositoryImpl) Payment(data entity.Payment) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		err := r.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&data).Error; err != nil {
				return err
			}

			if err := tx.Model(&entity.User{}).Where("user_id = ?", data.UserId).Update("balance", data.BalanceAfter).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			output <- model.Result{
				Error: err,
			}
		}

		output <- model.Result{
			Data: "ok",
		}
	}()

	return output
}

func (r *RepositoryImpl) Transfer(data entity.Transfer) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		err := r.DB.Transaction(func(tx *gorm.DB) error {
			// increase user target balance
			var targetUser entity.User
			getTargetUser := tx.Where("user_id = ?", data.TargetUserId).First(&targetUser)
			if getTargetUser.Error != nil {
				return fmt.Errorf("target user not found")
			}

			
			newBalanceTargetUser := targetUser.Balance + data.Amount
			if err := tx.Model(&entity.User{}).Where("user_id = ?", data.TargetUserId).Update("balance", newBalanceTargetUser).Error; err != nil {
				return err
			}

			// insert transactions history
			if err := tx.Create(&data).Error; err != nil {
				return err
			}

			// decrease user balance
			if err := tx.Model(&entity.User{}).Where("user_id = ?", data.UserId).Update("balance", data.BalanceAfter).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			output <- model.Result{
				Error: err,
			}
		}

		output <- model.Result{
			Data: "ok",
		}
	}()

	return output
}


func (r *RepositoryImpl) GetTransaction(ctx context.Context) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		var data []entity.Transfer
		result := r.DB.Find(&data)
		if result.Error != nil {
			output <- model.Result{
				Error: result.Error,
				Data:  nil,
			}
		}

		if result.RowsAffected == 0 {
			output <- model.Result{
				Error: fmt.Errorf("no data found"),
				Data:  nil,
			}
			return
		}
		output <- model.Result{
			Data: data,
		}
	}()

	return output
}

func (r *RepositoryImpl) UpdateUser(ctx context.Context, params map[string]interface{},  body map[string]interface{}) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		query := r.DB.Model(&entity.User{})
		for field, value := range params {
			query = query.Where(fmt.Sprintf("%s = ?", field), value)
		}
		result := query.Updates(body)
		if result.Error != nil {
			output <- model.Result{
				Error: result.Error,
			}
		}

		output <- model.Result{
			Data: result.RowsAffected,
		}
	}()

	return output
}