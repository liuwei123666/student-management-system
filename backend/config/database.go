package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"student-sys/models"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	// 从环境变量获取数据库连接信息
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "student_sys") // 使用新创建的用户
	password := getEnv("DB_PASSWORD", "student_sys") // 使用新创建的密码
	dbname := getEnv("DB_NAME", "student_system")

	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Database connected successfully")
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	// 自动迁移所有模型
	err := DB.AutoMigrate(
		&models.Student{},
		&models.Score{},
		&models.Honor{},
		&models.Attendance{},
		&models.CompetitionActivity{},
		&models.Employment{},
		&models.MentalHealth{},
		&models.Warning{},
		&models.ServiceMessage{},
		&models.TreeHole{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	fmt.Println("Database migration completed successfully")
	return nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
