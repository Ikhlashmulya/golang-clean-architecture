package repository

import (
	"context"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	CountByUsername(ctx context.Context, username string) (int64, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := new(domain.User)
	if err := r.DB.WithContext(ctx).Where("username = ?", username).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CountByUsername(ctx context.Context, username string) (int64, error) {
	var countUser int64
	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("username = ?", username).Count(&countUser).Error; err != nil {
		return 0, err
	}
	return countUser, nil
}