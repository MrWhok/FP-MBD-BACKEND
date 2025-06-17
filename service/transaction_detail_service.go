package service

import (
	"context"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

type TransactionDetailService interface {
	FindById(ctx context.Context, id string) model.TransactionDetailModel
}
