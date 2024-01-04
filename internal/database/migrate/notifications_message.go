package migrate

import "gorm.io/gorm"

type NotificationsMessage struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Priority int    `gorm:"type:int;not null"`
	Content  string `gorm:"type:varchar(300);not null"`
	UsersID  int64  `gorm:"type:int;not null"`
	Users    Users  `gorm:"foreignKey:UsersID"`
}
