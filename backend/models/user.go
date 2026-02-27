package models

// User 用户模型
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Phone    string `gorm:"uniqueIndex;size:20;not null" json:"phone"`
	Password string `gorm:"size:100;not null" json:"-"`
	Role     string `gorm:"size:10;not null" json:"role"` // 'teacher' 或 'student'
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
