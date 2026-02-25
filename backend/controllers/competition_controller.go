package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
)

// GetCompetitions 获取赛事活动列表（支持分页和student_id过滤）
func GetCompetitions(c *gin.Context) {
	var competitions []models.CompetitionActivity
	var total int64

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// 获取student_id参数
	studentID := c.Query("student_id")

	// 构建查询
	query := config.DB.Model(&models.CompetitionActivity{})

	// 添加student_id过滤条件
	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count competitions"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&competitions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get competitions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  competitions,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetCompetition 获取单个赛事活动详情
func GetCompetition(c *gin.Context) {
	id := c.Param("id")
	var competition models.CompetitionActivity

	if err := config.DB.First(&competition, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Competition not found"})
		return
	}

	c.JSON(http.StatusOK, competition)
}

// CreateCompetition 创建新赛事活动
func CreateCompetition(c *gin.Context) {
	var competition models.CompetitionActivity

	if err := c.ShouldBindJSON(&competition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&competition).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create competition"})
		return
	}

	c.JSON(http.StatusCreated, competition)
}

// UpdateCompetition 更新赛事活动信息
func UpdateCompetition(c *gin.Context) {
	id := c.Param("id")
	var competition models.CompetitionActivity

	// 查找赛事活动
	if err := config.DB.First(&competition, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Competition not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&competition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&competition).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update competition"})
		return
	}

	c.JSON(http.StatusOK, competition)
}

// DeleteCompetition 删除赛事活动
func DeleteCompetition(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.CompetitionActivity{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete competition"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Competition deleted successfully"})
}
