package service

import "github.com/MrWhok/FP-MBD-BACKEND/entity"

type UnpaidPaymentService interface {
	GetAllUnpaidPayments() ([]entity.UnpaidPaymentView, error)
}
