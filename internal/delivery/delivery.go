package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildandhya/mnc-rest-api/internal/model"
	"github.com/wildandhya/mnc-rest-api/internal/usecase"
)

type Delivery struct {
	uc usecase.Usecase
}

func NewDelivery(uc usecase.Usecase) Delivery {
	return Delivery{
		uc: uc,
	}
}

func (d *Delivery) Register(ctx echo.Context) error {
	var body model.RegisterRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	result := d.uc.Register(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) Login(ctx echo.Context) error {
	var body model.LoginRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	result := d.uc.Login(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) TopUp(ctx echo.Context) error {
	var body model.TopUpRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	result := d.uc.TopUp(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) Payment(ctx echo.Context) error {
	var body model.PaymentRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	result := d.uc.Payment(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) Transfer(ctx echo.Context) error {
	var body model.TransferRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	result := d.uc.Transfer(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) Transactions(ctx echo.Context) error {
	result := d.uc.TransactionReport(ctx.Request().Context())
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}

func (d *Delivery) UpdateProfile(ctx echo.Context) error {
	var body model.UpdateProfileRequest
	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: err.Error(),
		})
	}
	result := d.uc.UpdateProfile(ctx.Request().Context(), body)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status: "FAILURE",
			Message: result.Error.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status: "SUCCESS",
		Result: result.Data,
	})
}