package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/models"
	"student-sys/controllers"
)

// CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// 初始化数据库连接
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库表结构
	if err := config.AutoMigrate(); err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	// 初始化 Gin 引擎
	r := gin.Default()

	// 应用 CORS 中间件
	r.Use(corsMiddleware())

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 学生管理 API 路由组
	students := r.Group("/api/students")
	{
		students.GET("", getStudents)       // 获取学生列表（支持分页和搜索）
		students.GET("/:id", getStudent)    // 获取单个学生详情
		students.POST("", createStudent)    // 创建新学生
		students.PUT("/:id", updateStudent) // 更新学生信息
		students.DELETE("/:id", deleteStudent) // 删除学生
	}

	// 成绩管理 API 路由组
	scores := r.Group("/api/scores")
	{
		scores.GET("", controllers.GetScores)       // 获取成绩列表（支持分页和student_id过滤）
		scores.GET("/:id", controllers.GetScore)    // 获取单个成绩详情
		scores.POST("", controllers.CreateScore)    // 创建新成绩
		scores.PUT("/:id", controllers.UpdateScore) // 更新成绩信息
		scores.DELETE("/:id", controllers.DeleteScore) // 删除成绩
	}

	// 荣誉管理 API 路由组
	honors := r.Group("/api/honors")
	{
		honors.GET("", controllers.GetHonors)       // 获取荣誉列表（支持分页和student_id过滤）
		honors.GET("/:id", controllers.GetHonor)    // 获取单个荣誉详情
		honors.POST("", controllers.CreateHonor)    // 创建新荣誉
		honors.PUT("/:id", controllers.UpdateHonor) // 更新荣誉信息
		honors.DELETE("/:id", controllers.DeleteHonor) // 删除荣誉
	}

	// 考勤管理 API 路由组
	attendances := r.Group("/api/attendances")
	{
		attendances.GET("", controllers.GetAttendances)       // 获取考勤列表（支持分页和student_id过滤）
		attendances.GET("/:id", controllers.GetAttendance)    // 获取单个考勤详情
		attendances.POST("", controllers.CreateAttendance)    // 创建新考勤
		attendances.PUT("/:id", controllers.UpdateAttendance) // 更新考勤信息
		attendances.DELETE("/:id", controllers.DeleteAttendance) // 删除考勤
	}

	// 赛事活动管理 API 路由组
	competitions := r.Group("/api/competitions")
	{
		competitions.GET("", controllers.GetCompetitions)       // 获取赛事活动列表（支持分页和student_id过滤）
		competitions.GET("/:id", controllers.GetCompetition)    // 获取单个赛事活动详情
		competitions.POST("", controllers.CreateCompetition)    // 创建新赛事活动
		competitions.PUT("/:id", controllers.UpdateCompetition) // 更新赛事活动信息
		competitions.DELETE("/:id", controllers.DeleteCompetition) // 删除赛事活动
	}

	// 就业管理 API 路由组
	employments := r.Group("/api/employments")
	{
		employments.GET("", controllers.GetEmployments)       // 获取就业列表（支持分页和student_id过滤）
		employments.GET("/:id", controllers.GetEmployment)    // 获取单个就业详情
		employments.POST("", controllers.CreateEmployment)    // 创建新就业
		employments.PUT("/:id", controllers.UpdateEmployment) // 更新就业信息
		employments.DELETE("/:id", controllers.DeleteEmployment) // 删除就业
	}

	// 预警管理 API 路由组
	warnings := r.Group("/api/warnings")
	{
		warnings.GET("/dashboard", controllers.GetWarningDashboard) // 获取预警统计概况
		warnings.GET("", controllers.GetWarnings)                  // 获取预警列表
	}

	// 消息管理 API 路由组
	messages := r.Group("/api/messages")
	{
		messages.POST("", controllers.SendMessage)       // 发送服务消息
		messages.GET("", controllers.GetMessages)        // 获取消息列表
	}

	// 学生数字画像 API
	r.GET("/api/students/:id/portrait", controllers.GetStudentPortrait) // 获取学生数字画像

	// 数据总结 API
	r.GET("/api/data-summary", controllers.GetDataSummary) // 获取数据总结

	// 树洞 API 路由组
	treehole := r.Group("/api/treehole")
	{
		treehole.GET("", controllers.GetTreeHolePosts)       // 获取树洞帖子
		treehole.POST("", controllers.CreateTreeHolePost)    // 创建树洞帖子
		treehole.POST("/:id/like", controllers.LikeTreeHolePost) // 点赞树洞帖子
	}

	// 启动服务器
	port := 8080
	fmt.Printf("Server running on port %d\n", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getStudents 获取学生列表（支持分页、搜索和重点关注过滤）
func getStudents(c *gin.Context) {
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
		c.JSON(500, gin.H{"error": "Failed to count students"})
		return
	}

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&students).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get students"})
		return
	}

	c.JSON(200, gin.H{
		"data":  students,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// getStudent 获取单个学生详情
func getStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(200, student)
}

// createStudent 创建新学生
func createStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(201, student)
}

// updateStudent 更新学生信息
func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// 查找学生
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 保存更新
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(200, student)
}

// deleteStudent 删除学生
func deleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(200, gin.H{"message": "Student deleted successfully"})
}
