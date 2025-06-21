package impl

import (
	"github.com/MrWhok/FP-MBD-BACKEND/entity"
	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

type unpaidPaymentRepositoryImpl struct {
	*gorm.DB
}

func NewUnpaidPaymentRepositoryImpl(db *gorm.DB) repository.UnpaidPaymentRepository {
	return &unpaidPaymentRepositoryImpl{DB: db}
}

func (r *unpaidPaymentRepositoryImpl) FindAllUnpaidPayments() ([]entity.UnpaidPaymentView, error) {
	var results []entity.UnpaidPaymentView
	err := r.Table("unpaid_payments_view").Find(&results).Error
	return results, err
}
