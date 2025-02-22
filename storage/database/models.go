package database

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Hostname string `gorm:"unique"`
	Secret   string
}
