package model

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
}

func (m *Feed) TableName() string {
    return "feed"
}
