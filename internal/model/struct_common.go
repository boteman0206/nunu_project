package model

// 公共的结构体定义
type CommonData struct {
	// gorm.Model
	ID        int64 `gorm:"primarykey"`
	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}
