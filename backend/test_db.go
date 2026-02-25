package main

import (
	"fmt"
	"log"

	"student-sys/config"
	"student-sys/models"
)

func main() {
	// 初始化数据库连接
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	fmt.Println("Database connection test: SUCCESS")

	// 测试插入数据 - 使用新的 StudentID 避免重复键错误
	fmt.Println("\nTesting data insertion...")
	student := models.Student{
		StudentID: "20230002", // 使用新的 StudentID
		Name:      "李四",
		Gender:    "女",
		College:   "电子工程学院",
		Major:     "电子信息工程",
		Grade:     "2023级",
		Class:     "电信1班",
		Phone:     "13900139000",
		IsFocus:   true,
		Avatar:    "https://example.com/avatar2.jpg",
	}

	if err := config.DB.Create(&student).Error; err != nil {
		log.Fatalf("Failed to create student: %v", err)
	}

	fmt.Println("Data insertion test: SUCCESS")

	// 测试查询数据
	fmt.Println("\nTesting data query...")
	var result models.Student
	if err := config.DB.Where("student_id = ?", "20230002").First(&result).Error; err != nil {
		log.Fatalf("Failed to query student: %v", err)
	}

	fmt.Printf("Query result: StudentID=%s, Name=%s\n", result.StudentID, result.Name)
	fmt.Println("Data query test: SUCCESS")

	// 清空数据库 - 使用 GORM 的正确方法
	fmt.Println("\nClearing database...")

	// 使用 Delete 方法并传入一个空对象来删除所有记录
	if err := config.DB.Unscoped().Delete(&models.Student{}).Error; err != nil {
		log.Printf("Failed to clear student table: %v", err)
	} else {
		fmt.Println("Cleared student table successfully")
	}

	// 验证数据已清空
	var count int64
	config.DB.Model(&models.Student{}).Count(&count)
	fmt.Printf("\nStudent table count after clearing: %d\n", count)

	if count == 0 {
		fmt.Println("All tests passed successfully!")
	} else {
		fmt.Println("Warning: Database not cleared properly")
	}
}
