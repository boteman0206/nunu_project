package model

type User struct {
	// gorm.Model
	CommonData
	Username    string `gorm:"column:username;not null"`
	Password    string `gorm:"column:password;not null"`
	PortraitUrl string `gorm:"column:portrait_url;default:''"`
}

func (u *User) TableName() string {
	return "users"
}
