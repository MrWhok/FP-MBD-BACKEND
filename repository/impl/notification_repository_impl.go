package impl

import (
	"github.com/MrWhok/FP-MBD-BACKEND/entity"
	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

type notificationRepositoryImpl struct {
	*gorm.DB
}

func NewNotificationRepositoryImpl(db *gorm.DB) repository.NotificationRepository {
	return &notificationRepositoryImpl{DB: db}
}

func (r *notificationRepositoryImpl) GetUnsentNotifications() ([]entity.Notification, error) {
	var notifications []entity.Notification

	sql := `
        SELECT notification.*, customer.email 
        FROM notification 
        JOIN customer ON notification.customer_id = customer.customer_id 
        WHERE is_sent = FALSE
    `
	err := r.DB.Raw(sql).Scan(&notifications).Error
	return notifications, err
}

func (r *notificationRepositoryImpl) MarkAsSent(notificationID int) error {
	sql := `UPDATE notification SET is_sent = TRUE WHERE notification_id = ?`

	db := r.Session(&gorm.Session{
		NewDB:       true,
		PrepareStmt: false,
	})

	return db.Exec(sql, notificationID).Error
}
