package service

import (
	"context"
	"mime/multipart"
)

type PaymentService interface {
	UploadPaymentProof(ctx context.Context, reservationID int, file *multipart.FileHeader) error
	ConfirmPayment(ctx context.Context, reservationID int) error
}
