package service

import (
	"context"
	"errors"

	"github.com/akuaruu/adminin_api/internal/model"
	"github.com/akuaruu/adminin_api/internal/repository"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Register(ctx context.Context, admin *model.Admin) error
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{
		adminRepo: repo,
	}
}

func (a *adminService) Register(ctx context.Context, admin *model.Admin) error {

	_, err := a.adminRepo.FindByEmail(ctx, admin.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)

	return a.adminRepo.Create(ctx, admin)
}
