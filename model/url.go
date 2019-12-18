package model

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	User User `gorm:"association_foreignkey:Username"`
	URL   string `gorm:"unique_index;not null"`
	ShortURL string `gorm:"unique_index;not null"`
}

