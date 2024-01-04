package migrate

import "gorm.io/gorm"

type Companys struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null"`
}
