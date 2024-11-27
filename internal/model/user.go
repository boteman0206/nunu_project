package model

type User struct {
	// gorm.Model
	CommonData
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
}

func (u *User) TableName() string {
	return "users"
}
