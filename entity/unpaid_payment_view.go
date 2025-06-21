package entity

import "time"

type UnpaidPaymentView struct {
	ReservationID int       `json:"reservation_id"`
	CustomerID    int       `json:"customer_id"`
	CustomerName  string    `json:"customer_name"`
	Amount        float64   `json:"amount"`
	PaymentDue    time.Time `json:"payment_due"`
	PaymentProof  *string   `json:"payment_proof"`
}
