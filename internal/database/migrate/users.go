package migrate

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID        int64    `gorm:"primaryKey;autoIncrement"`
	Name      string   `gorm:"type:varchar(255);not null"`
	Email     string   `gorm:"type:varchar(255);not null"`
	Password  string   `gorm:"type:varchar(255);not null"`
	CompanyID int64    `gorm:"type:varchar(255);not null"`
	GroupID   int64    `gorm:"type:varchar(255);not null"`
	Companys  Companys `gorm:"foreignKey:CompanyID"`
	Groups    Groups   `gorm:"foreignKey:GroupID"`
}
