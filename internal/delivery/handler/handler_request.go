package handler

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/enum"
)

type ReqCreateReviews []ReqCreateReview

type ReqCreateProduct struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type ReqCreateReview struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
}

type ReqCreateUser struct {
	Name   string  `json:"name" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

func (r *ReqCreateUser) Aggregate() *aggregate.User {
	user, _ := aggregate.NewUser(r.Name, r.Amount)
	return user
}

type ReqTransaction struct {
	UserID        uint64  `json:"user_id" validate:"required"`
	Flag          int     `json:"flag" validate:"required"`
	UserReceiveID uint64  `json:"user_receive_id" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}

func (r *ReqTransaction) Aggregate() *aggregate.Transaction {
	transaction, _ := aggregate.NewTransaction(r.UserID, enum.Flag(r.Flag), r.UserReceiveID, r.Nominal)
	return transaction
}
