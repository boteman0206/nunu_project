package model

// Feed 用户投稿表
type Feed struct {
	ID           int64  `json:"id" gorm:"id"`                       // auto increment primary key
	UserId       int64  `json:"user_id" gorm:"user_id"`             // 用户id
	Title        string `json:"title" gorm:"title"`                 // 投稿标题
	Tag          string `json:"tag" gorm:"tag"`                     // 标签[]
	Description  string `json:"description" gorm:"description"`     // 投稿内容
	VoteCount    int64  `json:"vote_count" gorm:"vote_count"`       // 点赞数量
	ShareCount   int64  `json:"share_count" gorm:"share_count"`     // 分享数量
	CommonCount  int64  `json:"common_count" gorm:"common_count"`   // 评论数量
	ReportCount  int64  `json:"report_count" gorm:"report_count"`   // 举报数量
	ImageList    string `json:"image_list" gorm:"image_list"`       // multi json
	OperateName  string `json:"operate_name" gorm:"operate_name"`   // 后台操作 人
	DeleteStatus int8   `json:"delete_status" gorm:"delete_status"` // 删除状态，0=正常
	Type         int8   `json:"type" gorm:"type"`                   // 投稿类型,1为普通投稿,2图片投稿3为视频投稿
	DbTime       int64  `json:"db_time" gorm:"db_time"`             // 创建时间
	UpdateTime   int64  `json:"update_time" gorm:"update_time"`     // 更新时间
}

// TableName 表名称
func (*Feed) TableName() string {
	return "feed"
}
