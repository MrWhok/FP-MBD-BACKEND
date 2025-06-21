package impl

import (
	"context"
	"errors"

	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	*gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepositoryImpl{DB: db}
}

func (r *paymentRepositoryImpl) UpdatePaymentProof(ctx context.Context, reservationID int, proofPath string) error {
	sql := `UPDATE payment SET payment_proof = ? WHERE reservation_id = ?`
	err := r.DB.WithContext(ctx).Exec(sql, proofPath, reservationID).Error
	if err != nil {
		return errors.New("failed to update payment proof: " + err.Error())
	}
	return nil
}

func (r *paymentRepositoryImpl) ConfirmPayment(ctx context.Context, reservationID int) error {
	err := r.WithContext(ctx).Exec(`CALL confirm_payment(?)`, reservationID).Error
	if err != nil {
		return errors.New("failed to confirm payment: " + err.Error())
	}
	return nil
}
