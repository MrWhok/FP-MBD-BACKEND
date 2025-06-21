package impl

import (
	"github.com/MrWhok/FP-MBD-BACKEND/entity"
	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
)

type unpaidPaymentServiceImpl struct {
	repo repository.UnpaidPaymentRepository
}

func NewUnpaidPaymentServiceImpl(repo repository.UnpaidPaymentRepository) service.UnpaidPaymentService {
	return &unpaidPaymentServiceImpl{repo: repo}
}

func (s *unpaidPaymentServiceImpl) GetAllUnpaidPayments() ([]entity.UnpaidPaymentView, error) {
	return s.repo.FindAllUnpaidPayments()
}
