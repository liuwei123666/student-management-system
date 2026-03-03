package models

import "golang.org/x/crypto/bcrypt"

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

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword 验证密码
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
