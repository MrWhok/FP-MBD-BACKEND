package impl

import (
	"context"

	"github.com/MrWhok/FP-MBD-BACKEND/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (r *userRepositoryImpl) Register(ctx context.Context, nama, email, noTelp, password string) error {
	return r.DB.WithContext(ctx).Exec(`CALL register_customer(?, ?, ?, ?)`, nama, email, noTelp, password).Error
}

func (r *userRepositoryImpl) Login(ctx context.Context, email string) (string, int, string, error) {
	var hashedPassword string
	var customerID int
	var role string

	row := r.DB.WithContext(ctx).Raw(`SELECT customer_id, password, role FROM customer WHERE email = ?`, email).Row()
	err := row.Scan(&customerID, &hashedPassword, &role)
	if err != nil {
		return "", 0, "", err
	}

	return hashedPassword, customerID, role, nil
}

func (r *userRepositoryImpl) FindRolesByCustomerID(ctx context.Context, customerID int) ([]map[string]interface{}, error) {
	rows, err := r.DB.WithContext(ctx).Raw(`
		SELECT r.role
		FROM user_role ur
		JOIN role r ON ur.role_id = r.role_id
		WHERE ur.customer_id = ?
	`, customerID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []map[string]interface{}
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, map[string]interface{}{"role": role})
	}

	return roles, nil
}
