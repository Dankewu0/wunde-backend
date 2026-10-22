package main

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title string
	Slug  string `gorm:"uniqueIndex"`
}
