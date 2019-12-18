package model

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	URL   string `gorm:"unique_index;not null"`
	User User `gorm:"association_foreignkey:Username"`
}

