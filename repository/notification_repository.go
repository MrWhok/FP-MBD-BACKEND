package repository

import "github.com/MrWhok/FP-MBD-BACKEND/entity"

type NotificationRepository interface {
	GetUnsentNotifications() ([]entity.Notification, error)
	MarkAsSent(id int) error
}
