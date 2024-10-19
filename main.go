package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/wildandhya/mnc-rest-api/db"
	"github.com/wildandhya/mnc-rest-api/internal/delivery"
	"github.com/wildandhya/mnc-rest-api/internal/repository"
	"github.com/wildandhya/mnc-rest-api/internal/usecase"
	"github.com/wildandhya/mnc-rest-api/pkg/helper"
)

func main() {
	e := echo.New()

	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	db := db.NewDatabase()
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handlers := delivery.NewDelivery(usecase)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	e.POST("/topup", handlers.TopUp, helper.ValidateBearer())
	e.POST("/pay", handlers.Payment, helper.ValidateBearer())
	e.POST("/transfer", handlers.Transfer, helper.ValidateBearer())
	e.GET("/transactions", handlers.Transactions, helper.ValidateBearer())
	e.PUT("/profile", handlers.UpdateProfile, helper.ValidateBearer())
	e.Logger.Fatal(e.Start(":8080"))
}
