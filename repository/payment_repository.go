package repository

import (
	"context"
)

type PaymentRepository interface {
	UpdatePaymentProof(ctx context.Context, reservationID int, proofPath string) error
	ConfirmPayment(ctx context.Context, reservationID int) error
}
