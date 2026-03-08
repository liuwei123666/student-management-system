package config

import (
	"fmt"
	"student-sys/models"
)

// SeedData 初始化种子数据
func SeedData() error {
	// 初始化默认教师账号
	if err := seedDefaultTeacher(); err != nil {
		return err
	}

	// 可以在这里添加其他种子数据
	// if err := seedOtherData(); err != nil {
	// 	return err
	// }

	fmt.Println("Seed data completed successfully")
	return nil
}

// seedDefaultTeacher 初始化默认教师账号
func seedDefaultTeacher() error {
	var teacherCount int64
	DB.Model(&models.User{}).Where("role = ?", "teacher").Count(&teacherCount)
	if teacherCount == 0 {
		// 加密密码
		hashedPassword, err := models.HashPassword("123456")
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		// 创建默认教师账号
		teacher := models.User{
			Phone:    "admin",
			Password: hashedPassword,
			Role:     "teacher",
		}

		if err := DB.Create(&teacher).Error; err != nil {
			return fmt.Errorf("failed to create default teacher: %w", err)
		}

		fmt.Println("默认教师账号已创建")
	}
	return nil
}

// seedOtherData 初始化其他种子数据
func seedOtherData() error {
	// 在这里添加其他种子数据
	// 例如：默认学院、专业、课程等
	return nil
}
