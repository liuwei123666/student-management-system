package models

import (
	"time"

	"gorm.io/gorm"
)

// Student 基础信息表
type Student struct {
	gorm.Model
	StudentID string `gorm:"type:varchar(20);uniqueIndex;comment:学号" json:"StudentID"`
	Name      string `gorm:"type:varchar(50);comment:姓名" json:"Name"`
	Gender    string `gorm:"type:varchar(10);comment:性别" json:"Gender"`
	College   string `gorm:"type:varchar(100);comment:学院" json:"College"`
	Major     string `gorm:"type:varchar(100);comment:专业" json:"Major"`
	Grade     string `gorm:"type:varchar(20);comment:年级" json:"Grade"`
	Class     string `gorm:"type:varchar(50);comment:班级" json:"Class"`
	Phone     string `gorm:"type:varchar(20);comment:联系电话" json:"Phone"`
	IsFocus   bool   `gorm:"comment:重点关注状态" json:"IsFocus"`
	Avatar    string `gorm:"type:varchar(255);comment:头像URL" json:"Avatar"`
}

// Score 成绩表
type Score struct {
	gorm.Model
	StudentID  string  `gorm:"type:varchar(20);index;comment:关联学生学号" json:"student_id"`
	CourseName string  `gorm:"type:varchar(100);comment:课程名称" json:"course_name"`
	Credit     float64 `gorm:"type:float;comment:学分" json:"credit"`
	Mark       float64 `gorm:"type:float;comment:分数" json:"mark"`
	Term       string  `gorm:"type:varchar(20);comment:学期" json:"term"`
	IsFailed   bool    `gorm:"comment:是否挂科" json:"is_failed"`
}

// Honor 荣誉表
type Honor struct {
	gorm.Model
	StudentID  string    `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	Title      string    `gorm:"type:varchar(100);comment:荣誉名称" json:"title"`
	Level      string    `gorm:"type:varchar(20);comment:级别" json:"level"`
	AwardDate  time.Time `gorm:"comment:获奖时间" json:"award_date"`
}

// Attendance 安全考勤表
type Attendance struct {
	gorm.Model
	StudentID   string    `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	Date        time.Time `gorm:"comment:考勤日期" json:"date"`
	Status      string    `gorm:"type:varchar(20);comment:状态" json:"status"`
	Location    string    `gorm:"type:varchar(100);comment:打卡位置/宿舍号" json:"location"`
	Description string    `gorm:"type:text;comment:备注说明" json:"description"`
}

// CompetitionActivity 赛事与活动表
type CompetitionActivity struct {
	gorm.Model
	StudentID string    `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	Type      string    `gorm:"type:varchar(20);comment:类型" json:"type"`
	Name      string    `gorm:"type:varchar(100);comment:赛事/活动名称" json:"name"`
	Role      string    `gorm:"type:varchar(50);comment:担任角色/参赛身份" json:"role"`
	Date      time.Time `gorm:"comment:参与时间" json:"date"`
	Award     string    `gorm:"type:varchar(100);comment:获奖情况/活动时长" json:"award"`
}

// Employment 就业表
type Employment struct {
	gorm.Model
	StudentID string `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	Status    string `gorm:"type:varchar(20);comment:就业状态" json:"status"`
	Company   string `gorm:"type:varchar(100);comment:签约公司/实习单位" json:"company"`
	Position  string `gorm:"type:varchar(100);comment:岗位" json:"position"`
}

// MentalHealth 心理健康表
type MentalHealth struct {
	gorm.Model
	StudentID       string    `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	AssessmentDate  time.Time `gorm:"comment:评估日期" json:"assessment_date"`
	Level           string    `gorm:"type:varchar(20);comment:心理评级" json:"level"`
	Counselor       string    `gorm:"type:varchar(50);comment:负责辅导员/心理老师" json:"counselor"`
	Notes           string    `gorm:"type:text;comment:备注" json:"notes"`
}

// Warning 学生预警表
type Warning struct {
	gorm.Model
	StudentID   string `gorm:"type:varchar(20);index;comment:关联学生" json:"student_id"`
	WarningType string `gorm:"type:varchar(20);comment:预警类型" json:"warning_type"`
	Level       string `gorm:"type:varchar(10);comment:预警等级" json:"level"`
	Description string `gorm:"type:text;comment:预警详细原因说明" json:"description"`
	Status      string `gorm:"type:varchar(20);comment:处理状态" json:"status"`
}

// ServiceMessage 服务推送表
type ServiceMessage struct {
	gorm.Model
	StudentID string `gorm:"type:varchar(20);index;comment:接收方学号" json:"student_id"`
	Title     string `gorm:"type:varchar(100);comment:推送标题" json:"title"`
	Content   string `gorm:"type:text;comment:推送内容" json:"content"`
	IsRead    bool   `gorm:"comment:是否已读" json:"is_read"`
}

// TreeHole 树洞表
type TreeHole struct {
	gorm.Model
	Content       string `gorm:"type:text;comment:树洞内容" json:"content"`
	BackgroundColor string `gorm:"type:varchar(20);comment:卡片背景色" json:"background_color"`
	Likes         int    `gorm:"comment:点赞数" json:"likes"`
}
