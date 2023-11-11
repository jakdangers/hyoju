package entity

import "gorm.io/gorm"

/*
Group
*/
type Group struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	Code        string
}
