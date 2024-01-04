package entitys

import (
	"time"

	"gorm.io/gorm"
)

type Companys struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Companys) BeforeCreate(db *gorm.DB) (err error) {
	c.CreatedAt = time.Now().UTC().Local()
	return nil
}

func (c *Companys) BeforeUpdate(db *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UTC().Local()
	return nil
}
