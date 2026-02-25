package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetHonors 获取荣誉列表（支持分页和student_id过滤）
func GetHonors(c *gin.Context) {
	var honors []models.Honor
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取student_id参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.Honor{})

	// 添加student_id过滤条件
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count honors"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&honors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get honors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  honors,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetHonor 获取单个荣誉详情
func GetHonor(c *gin.Context) {
	id := c.Param("id")
	var honor models.Honor

	if err := config.DB.First(&honor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Honor not found"})
		return
	}

	c.JSON(http.StatusOK, honor)
}

// CreateHonor 创建新荣誉
func CreateHonor(c *gin.Context) {
	var honor models.Honor

	if err := c.ShouldBindJSON(&honor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&honor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create honor"})
		return
	}

	c.JSON(http.StatusCreated, honor)
}

// UpdateHonor 更新荣誉信息
func UpdateHonor(c *gin.Context) {
	id := c.Param("id")
	var honor models.Honor

	// 查找荣誉
	if err := config.DB.First(&honor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Honor not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&honor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&honor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update honor"})
		return
	}

	c.JSON(http.StatusOK, honor)
}

// DeleteHonor 删除荣誉
func DeleteHonor(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Honor{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete honor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Honor deleted successfully"})
}
