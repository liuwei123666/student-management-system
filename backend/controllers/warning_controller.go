package controllers

import (
	"student-sys/config"
	"student-sys/models"

	"github.com/gin-gonic/gin"
)

// GetWarningDashboard 获取预警统计概况
func GetWarningDashboard(c *gin.Context) {
	// 统计学业预警总数
	var academicWarningCount int64
	config.DB.Model(&models.Warning{}).Where("warning_type = ?", "学业").Count(&academicWarningCount)

	// 统计安全预警总数
	var safetyWarningCount int64
	config.DB.Model(&models.Warning{}).Where("warning_type = ?", "安全").Count(&safetyWarningCount)

	// 统计心理预警总数
	var mentalWarningCount int64
	config.DB.Model(&models.Warning{}).Where("warning_type = ?", "心理").Count(&mentalWarningCount)

	// 统计综合预警总数（所有类型）
	var totalWarningCount int64
	config.DB.Model(&models.Warning{}).Count(&totalWarningCount)

	// 统计处理状态
	var statusCounts []struct {
		Status string
		Count  int64
	}
	config.DB.Model(&models.Warning{}).Select("status, count(*) as count").Group("status").Find(&statusCounts)

	// 构造响应
	response := gin.H{
		"academicWarningCount": academicWarningCount,
		"safetyWarningCount":   safetyWarningCount,
		"mentalWarningCount":   mentalWarningCount,
		"totalWarningCount":    totalWarningCount,
		"statusCounts":         statusCounts,
	}

	c.JSON(200, response)
}

// GetWarnings 获取预警列表
func GetWarnings(c *gin.Context) {
	var warnings []models.Warning
	var total int64

	// 获取分页参数
	page := 1
	pageSize := 10

	// 构建查询
	query := config.DB.Model(&models.Warning{})

	// 计算总数
	query.Count(&total)

	// 执行分页查询，按创建时间倒序
	query.Order("created_at DESC").Limit(pageSize).Find(&warnings)

	c.JSON(200, gin.H{
		"data":  warnings,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
