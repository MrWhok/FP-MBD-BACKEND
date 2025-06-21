package repository

import "github.com/MrWhok/FP-MBD-BACKEND/entity"

type UnpaidPaymentRepository interface {
	FindAllUnpaidPayments() ([]entity.UnpaidPaymentView, error)
}
