package repositories

import (
	"notification-system/internal/entitys"

	"gorm.io/gorm"
)

type IGroupRepository interface {
	Create(group *entitys.Groups) (*entitys.Groups, error)
	Get(id int64) (*entitys.Groups, error)
	Update(group *entitys.Groups) (*entitys.Groups, error)
	All() (*entitys.Groups, error)
}

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db}
}

func (r *GroupRepository) Create(group *entitys.Groups) (*entitys.Groups, error) {
	return group, r.db.Create(&group).Error
}

func (r *GroupRepository) Get(id int64) (*entitys.Groups, error) {
	var group entitys.Groups
	err := r.db.Where("id =?", id).First(&group).Error
	return &group, err
}

func (r *GroupRepository) Update(group *entitys.Groups) (*entitys.Groups, error) {
	err := r.db.Save(group).Error
	return group, err
}

func (r *GroupRepository) All() (*entitys.Groups, error) {
	var groups entitys.Groups
	err := r.db.Find(&groups).Error
	return &groups, err
}
