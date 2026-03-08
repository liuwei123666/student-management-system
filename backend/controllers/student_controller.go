package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetStudents 获取学生列表（支持分页、搜索和重点关注过滤）
func GetStudents(c *gin.Context) {
	var students []models.Student
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取搜索参数
	search := c.Query("search")
	
	// 获取重点关注过滤参数
	focusOnly := c.Query("focusOnly")

	// 构建查询
	query := config.DB.Model(&models.Student{})

	// 添加搜索条件
	if search != "" {
		query = query.Where("student_id LIKE ? OR name LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	
	// 添加重点关注过滤条件
	if focusOnly == "true" {
		query = query.Where("is_focus = ?", true)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count students"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get students"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetStudent 获取单个学生详情
func GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// CreateStudent 创建新学生
func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// UpdateStudent 更新学生信息
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// 查找学生
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent 删除学生
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
