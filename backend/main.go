package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"student-sys/config"
	"student-sys/controllers"
	"student-sys/middleware"
)

// CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许的前端域名
		allowedOrigins := []string{
			"http://localhost:5173", // 前端开发服务器
			"http://localhost:3000", // 可能的其他前端端口
		}

		// 检查请求来源是否在允许列表中
		origin := c.Request.Header.Get("Origin")
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		// 设置其他CORS头
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400") // 24小时

		// 安全头信息
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")

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

	// 初始化种子数据
	if err := config.SeedData(); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	// 初始化 Gin 引擎
	r := gin.Default()

	// 应用 CORS 中间件
	r.Use(corsMiddleware())
	
	// 应用错误处理中间件
	r.Use(middleware.ErrorHandler())

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 认证 API 路由组 - 公开
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register) // 注册
		auth.POST("/login", controllers.Login)     // 登录
	}

	// 受JWT保护的业务路由组
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())       // JWT鉴权
	api.Use(middleware.TeacherAuthMiddleware()) // 教师角色鉴权
	{
		// 学生管理 API 路由组
		students := api.Group("/students")
		{
			students.GET("", controllers.GetStudents)       // 获取学生列表（支持分页和搜索）
			students.GET("/:id", controllers.GetStudent)    // 获取单个学生详情
			students.POST("", controllers.CreateStudent)    // 创建新学生
			students.PUT("/:id", controllers.UpdateStudent) // 更新学生信息
			students.DELETE("/:id", controllers.DeleteStudent) // 删除学生
		}

		// 成绩管理 API 路由组
		scores := api.Group("/scores")
		{
			scores.GET("", controllers.GetScores)       // 获取成绩列表（支持分页和student_id过滤）
			scores.GET("/:id", controllers.GetScore)    // 获取单个成绩详情
			scores.POST("", controllers.CreateScore)    // 创建新成绩
			scores.PUT("/:id", controllers.UpdateScore) // 更新成绩信息
			scores.DELETE("/:id", controllers.DeleteScore) // 删除成绩
		}

		// 荣誉管理 API 路由组
		honors := api.Group("/honors")
		{
			honors.GET("", controllers.GetHonors)       // 获取荣誉列表（支持分页和student_id过滤）
			honors.GET("/:id", controllers.GetHonor)    // 获取单个荣誉详情
			honors.POST("", controllers.CreateHonor)    // 创建新荣誉
			honors.PUT("/:id", controllers.UpdateHonor) // 更新荣誉信息
			honors.DELETE("/:id", controllers.DeleteHonor) // 删除荣誉
		}

		// 考勤管理 API 路由组
		attendances := api.Group("/attendances")
		{
			attendances.GET("", controllers.GetAttendances)       // 获取考勤列表（支持分页和student_id过滤）
			attendances.GET("/:id", controllers.GetAttendance)    // 获取单个考勤详情
			attendances.POST("", controllers.CreateAttendance)    // 创建新考勤
			attendances.PUT("/:id", controllers.UpdateAttendance) // 更新考勤信息
			attendances.DELETE("/:id", controllers.DeleteAttendance) // 删除考勤
		}

		// 赛事活动管理 API 路由组
		competitions := api.Group("/competitions")
		{
			competitions.GET("", controllers.GetCompetitions)       // 获取赛事活动列表（支持分页和student_id过滤）
			competitions.GET("/:id", controllers.GetCompetition)    // 获取单个赛事活动详情
			competitions.POST("", controllers.CreateCompetition)    // 创建新赛事活动
			competitions.PUT("/:id", controllers.UpdateCompetition) // 更新赛事活动信息
			competitions.DELETE("/:id", controllers.DeleteCompetition) // 删除赛事活动
		}

		// 就业管理 API 路由组
		employments := api.Group("/employments")
		{
			employments.GET("", controllers.GetEmployments)       // 获取就业列表（支持分页和student_id过滤）
			employments.GET("/:id", controllers.GetEmployment)    // 获取单个就业详情
			employments.POST("", controllers.CreateEmployment)    // 创建新就业
			employments.PUT("/:id", controllers.UpdateEmployment) // 更新就业信息
			employments.DELETE("/:id", controllers.DeleteEmployment) // 删除就业
		}

		// 预警管理 API 路由组
		warnings := api.Group("/warnings")
		{
			warnings.GET("/dashboard", controllers.GetWarningDashboard) // 获取预警统计概况
			warnings.GET("", controllers.GetWarnings)                  // 获取预警列表
		}

		// 消息管理 API 路由组
		messages := api.Group("/messages")
		{
			messages.POST("", controllers.SendMessage)       // 发送服务消息
			messages.GET("", controllers.GetMessages)        // 获取消息列表
		}

		// 学生数字画像 API
		api.GET("/students/:id/portrait", controllers.GetStudentPortrait) // 获取学生数字画像

		// 数据总结 API
		api.GET("/data-summary", controllers.GetDataSummary) // 获取数据总结

		// 树洞 API 路由组
		treehole := api.Group("/treehole")
		{
			treehole.GET("", controllers.GetTreeHolePosts)       // 获取树洞帖子
			treehole.POST("", controllers.CreateTreeHolePost)    // 创建树洞帖子
			treehole.POST("/:id/like", controllers.LikeTreeHolePost) // 点赞树洞帖子
		}
	}

	// 启动服务器
	port := 8080
	fmt.Printf("Server running on port %d\n", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}


