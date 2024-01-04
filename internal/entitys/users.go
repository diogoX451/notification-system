package entitys

import (
	"notification-system/pkg"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	Email     string    `gorm:"column:email;type:varchar(255);not null"`
	Password  string    `gorm:"column:password;type:varchar(255);not null"`
	CompanyID int64     `gorm:"column:company_id;type:varchar(255);not null"`
	GroupID   int64     `gorm:"column:group_id;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Users) BeforeCreate(db *gorm.DB) (err error) {
	c.CreatedAt = time.Now().UTC().Local()
	c.Password = pkg.HashPassword(c.Password)
	return nil
}

func (c *Users) BeforeUpdate(db *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UTC().Local()
	return nil
}

func (c *Users) TableName() string {
	return "users"
}
