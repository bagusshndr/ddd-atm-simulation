package handler

import (
	"ddd-atm-simulation/internal/shared"
	"ddd-atm-simulation/internal/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseUser        usecase.UserUsecase
	usecaseTransaction usecase.TransactionUsecase
}

func (h *handler) HandlerCreateTransaction(c echo.Context) error {
	var body ReqTransaction
	c.Bind(&body)
	err := h.usecaseTransaction.CreateTransaction(*body.Aggregate())

	if err != nil {
		log.Println(err)
		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "Success", nil, nil).JSON(c)
}

func NewHandler(usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) *handler {
	return &handler{
		usecaseUser:        usecaseUser,
		usecaseTransaction: usecaseTransaction,
	}
}
