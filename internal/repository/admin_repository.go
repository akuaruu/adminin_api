package repository

import (
	"context"

	"gorm.io/gorm"
	"github.com/akuaruu/adminin_api/internal/model"
)

type AdminRepository interface {
	Create(ctx context.Context, admin *model.Admin) error
	FindByEmail(ctx context.Context, email string) (*model.Admin, error)

}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository{
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) Create(ctx context.Context, admin *model.Admin) error {
	return r.db.WithContext(ctx).Create(admin).Error
}

func (r *adminRepository) FindByEmail(ctx context.Context, email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}