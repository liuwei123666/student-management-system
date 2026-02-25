package controllers

import (
	"math/rand"
	"net/http"
	"student-sys/config"
	"student-sys/models"

	"github.com/gin-gonic/gin"
)

// StudentPortrait 学生数字画像结构体
type StudentPortrait struct {
	BasicInfo     models.Student              `json:"basic_info"`
	Scores        []models.Score              `json:"scores"`
	AttendanceStats map[string]int           `json:"attendance_stats"`
	ActivityCount  int                       `json:"activity_count"`
	Warnings       []models.Warning           `json:"warnings"`
}

// GetStudentPortrait 获取学生数字画像
func GetStudentPortrait(c *gin.Context) {
	studentID := c.Param("id")

	// 获取学生基本信息
	var student models.Student
	if err := config.DB.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// 获取学生所有成绩
	var scores []models.Score
	config.DB.Where("student_id = ?", studentID).Find(&scores)

	// 统计考勤情况
	var attendances []models.Attendance
	config.DB.Where("student_id = ?", studentID).Find(&attendances)

	attendanceStats := make(map[string]int)
	for _, attendance := range attendances {
		attendanceStats[attendance.Status]++
	}

	// 统计参与的活动数
	var activityCount int64
	config.DB.Model(&models.CompetitionActivity{}).Where("student_id = ?", studentID).Count(&activityCount)

	// 获取预警记录
	var warnings []models.Warning
	config.DB.Where("student_id = ?", studentID).Find(&warnings)

	// 构建数字画像响应
	portrait := StudentPortrait{
		BasicInfo:      student,
		Scores:         scores,
		AttendanceStats: attendanceStats,
		ActivityCount:  int(activityCount),
		Warnings:       warnings,
	}

	c.JSON(http.StatusOK, portrait)
}

// DataSummary 数据总结结构体
type DataSummary struct {
	CollegeDistribution []CollegeDistribution `json:"college_distribution"`
	WarningTrends       []WarningTrend       `json:"warning_trends"`
}

// CollegeDistribution 学院分布结构体
type CollegeDistribution struct {
	College string `json:"college"`
	Count   int    `json:"count"`
}

// WarningTrend 预警趋势结构体
type WarningTrend struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

// GetDataSummary 获取数据总结
func GetDataSummary(c *gin.Context) {
	// 模拟学院分布数据
	collegeDistribution := []CollegeDistribution{
		{College: "计算机学院", Count: 350},
		{College: "电子工程学院", Count: 280},
		{College: "机械工程学院", Count: 220},
		{College: "经济管理学院", Count: 180},
		{College: "人文学院", Count: 150},
	}

	// 模拟预警趋势数据
	warningTrends := []WarningTrend{
		{Month: "1月", Count: 12},
		{Month: "2月", Count: 8},
		{Month: "3月", Count: 15},
		{Month: "4月", Count: 10},
		{Month: "5月", Count: 18},
		{Month: "6月", Count: 22},
		{Month: "7月", Count: 5},
		{Month: "8月", Count: 3},
		{Month: "9月", Count: 14},
		{Month: "10月", Count: 16},
		{Month: "11月", Count: 19},
		{Month: "12月", Count: 25},
	}

	summary := DataSummary{
		CollegeDistribution: collegeDistribution,
		WarningTrends:       warningTrends,
	}

	c.JSON(http.StatusOK, summary)
}

// GetTreeHolePosts 获取树洞帖子
func GetTreeHolePosts(c *gin.Context) {
	var posts []models.TreeHole
	config.DB.Order("created_at desc").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// CreateTreeHolePost 创建树洞帖子
func CreateTreeHolePost(c *gin.Context) {
	var post models.TreeHole
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 随机生成背景色
	backgroundColors := []string{"#FFF3E0", "#E3F2FD", "#E8F5E8", "#F3E5F5", "#FFFDE7"}
	post.BackgroundColor = backgroundColors[rand.Intn(len(backgroundColors))]
	post.Likes = 0

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// LikeTreeHolePost 点赞树洞帖子
func LikeTreeHolePost(c *gin.Context) {
	id := c.Param("id")
	var post models.TreeHole

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Likes++
	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes"})
		return
	}

	c.JSON(http.StatusOK, post)
}
