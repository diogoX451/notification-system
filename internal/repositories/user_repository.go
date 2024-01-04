package repositories

import (
	"notification-system/internal/entitys"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *entitys.Users) (*entitys.Users, error)
	Get(id int64) (*entitys.Users, error)
	All() (*entitys.Users, error)
	Update(id int64, user *entitys.Users) (*entitys.Users, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entitys.Users) (*entitys.Users, error) {
	return user, r.db.Create(user).Error
}

func (r *UserRepository) Get(id int64) (*entitys.Users, error) {
	var user entitys.Users
	return &user, r.db.Where("id =?", id).First(&user).Error
}

func (r *UserRepository) All() (*entitys.Users, error) {
	var user entitys.Users
	return &user, r.db.Find(&user).Error
}

func (r *UserRepository) Update(id int64, user *entitys.Users) (*entitys.Users, error) {
	return user, r.db.Model(user).Where("id =?", id).Updates(user).Error
}
