package entitys

import (
	"time"

	"gorm.io/gorm"
)

type Groups struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (g *Groups) BeforeCreate(db *gorm.DB) (err error) {
	g.CreatedAt = time.Now().UTC().Local()
	return nil
}

func (g *Groups) BeforeUpdate(db *gorm.DB) (err error) {
	g.UpdatedAt = time.Now().UTC().Local()
	return nil
}
