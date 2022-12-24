package router

import (
	"ddd-atm-simulation/internal/delivery/handler"
	"ddd-atm-simulation/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Router(route *echo.Echo, usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) {
	h := handler.NewHandler(usecaseUser, usecaseTransaction)

	v1 := route.Group("")
	{
		v1.POST("/create-transaction", h.HandlerCreateTransaction)
	}

}
