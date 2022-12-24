package router

import (
	"ddd-to-do-list/internal/delivery/handler"
	"ddd-to-do-list/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Router(route *echo.Echo, usecaseUser usecase.UserUsecase, usecaseTransaction usecase.TransactionUsecase) {
	h := handler.NewHandler(usecaseUser, usecaseTransaction)

	v1 := route.Group("")
	{
		v1.POST("/create-transaction", h.HandlerCreateTransaction)
	}

}
