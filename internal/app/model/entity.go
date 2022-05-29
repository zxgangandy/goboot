package model

import (
	"time"
)

type User struct {
	Id       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name     string    `gorm:"column:name"`
	Age      int       `gorm:"column:age"`
	Address  string    `gorm:"column:address"`
	Password string    `gorm:"column:password"`
	Modified time.Time `gorm:"column:modified"`
	Created  time.Time `gorm:"column:created"`
}
