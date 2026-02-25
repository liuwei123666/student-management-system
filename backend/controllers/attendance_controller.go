package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetAttendances 获取考勤列表（支持分页和student_id过滤）
func GetAttendances(c *gin.Context) {
	var attendances []models.Attendance
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取student_id参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.Attendance{})

	// 添加student_id过滤条件
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count attendances"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&attendances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attendances"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetAttendance 获取单个考勤详情
func GetAttendance(c *gin.Context) {
	id := c.Param("id")
	var attendance models.Attendance

	if err := config.DB.First(&attendance, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance not found"})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// CreateAttendance 创建新考勤
func CreateAttendance(c *gin.Context) {
	var attendance models.Attendance

	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attendance"})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

// UpdateAttendance 更新考勤信息
func UpdateAttendance(c *gin.Context) {
	id := c.Param("id")
	var attendance models.Attendance

	// 查找考勤
	if err := config.DB.First(&attendance, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// DeleteAttendance 删除考勤
func DeleteAttendance(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Attendance{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance deleted successfully"})
}
