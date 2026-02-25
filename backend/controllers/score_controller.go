package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetScores 获取成绩列表（支持分页和student_id过滤）
func GetScores(c *gin.Context) {
	var scores []models.Score
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取student_id参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.Score{})

	// 添加student_id过滤条件
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count scores"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&scores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  scores,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetScore 获取单个成绩详情
func GetScore(c *gin.Context) {
	id := c.Param("id")
	var score models.Score

	if err := config.DB.First(&score, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	c.JSON(http.StatusOK, score)
}

// CreateScore 创建新成绩
func CreateScore(c *gin.Context) {
	var score models.Score

	if err := c.ShouldBindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create score"})
		return
	}

	c.JSON(http.StatusCreated, score)
}

// UpdateScore 更新成绩信息
func UpdateScore(c *gin.Context) {
	id := c.Param("id")
	var score models.Score

	// 查找成绩
	if err := config.DB.First(&score, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update score"})
		return
	}

	c.JSON(http.StatusOK, score)
}

// DeleteScore 删除成绩
func DeleteScore(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Score{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete score"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Score deleted successfully"})
}
