package controllers

import (
	"student-sys/config"
	"student-sys/models"

	"github.com/gin-gonic/gin"
)

// SendMessage 向特定 StudentID 推送服务消息
func SendMessage(c *gin.Context) {
	var message models.ServiceMessage

	// 绑定请求数据
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 设置默认值
	message.IsRead = false

	// 保存消息到数据库
	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(201, message)
}

// GetMessages 获取学生的服务消息列表
func GetMessages(c *gin.Context) {
	var messages []models.ServiceMessage
	var total int64

	// 获取学生ID参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.ServiceMessage{})

	// 添加学生ID过滤
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	query.Count(&total)

	// 执行查询，按创建时间倒序
	query.Order("created_at DESC").Find(&messages)

	c.JSON(200, gin.H{
		"data":  messages,
		"total": total,
	})
}
