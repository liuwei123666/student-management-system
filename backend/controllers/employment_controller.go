package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetEmployments 获取就业列表（支持分页和student_id过滤）
func GetEmployments(c *gin.Context) {
	var employments []models.Employment
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取student_id参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.Employment{})

	// 添加student_id过滤条件
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count employments"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&employments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employments,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetEmployment 获取单个就业详情
func GetEmployment(c *gin.Context) {
	id := c.Param("id")
	var employment models.Employment

	if err := config.DB.First(&employment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employment not found"})
		return
	}

	c.JSON(http.StatusOK, employment)
}

// CreateEmployment 创建新就业
func CreateEmployment(c *gin.Context) {
	var employment models.Employment

	if err := c.ShouldBindJSON(&employment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&employment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employment"})
		return
	}

	c.JSON(http.StatusCreated, employment)
}

// UpdateEmployment 更新就业信息
func UpdateEmployment(c *gin.Context) {
	id := c.Param("id")
	var employment models.Employment

	// 查找就业
	if err := config.DB.First(&employment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employment not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&employment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&employment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employment"})
		return
	}

	c.JSON(http.StatusOK, employment)
}

// DeleteEmployment 删除就业
func DeleteEmployment(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Employment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employment deleted successfully"})
}
