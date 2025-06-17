package repository

import (
	"context"
	"github.com/MrWhok/FP-MBD-BACKEND/entity"
)

type TransactionDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}
