package impl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"mime/multipart"

	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"github.com/MrWhok/FP-MBD-BACKEND/service"
)

type paymentServiceImpl struct {
	repo repository.PaymentRepository
}

func NewPaymentServiceImpl(repo repository.PaymentRepository) service.PaymentService {
	return &paymentServiceImpl{repo: repo}
}

func (s *paymentServiceImpl) UploadPaymentProof(ctx context.Context, reservationID int, file *multipart.FileHeader) error {
	dir := "media/payment"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	filename := fmt.Sprintf("proof_%d_%d%s", reservationID, time.Now().Unix(), filepath.Ext(file.Filename))
	savePath := filepath.Join(dir, filename)

	// Simpan file menggunakan Fiber context
	// Karena tidak ada ctx *fiber.Ctx di parameter, perlu disediakan dari luar (controller)
	// Jadi: simpan logika simpan file di controller
	// Atau → inject ctx.FiberCtx jika perlu di signature UploadPaymentProof

	// Tapi sebagai alternatif langsung dari multipart.FileHeader:
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Update path di DB
	return s.repo.UpdatePaymentProof(ctx, reservationID, savePath)
}

func (s *paymentServiceImpl) ConfirmPayment(ctx context.Context, reservationID int) error {
	return s.repo.ConfirmPayment(ctx, reservationID)
}
