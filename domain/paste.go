package domain

import "gorm.io/gorm"

type Paste struct {
	gorm.Model
	ShortHash string `gorm:"type:char(7);not NULL;unique"`
	ExpireMin int    `gorm:"default:0"`
	PastPath  string `gorm:"type:varchar(255)"`
}
