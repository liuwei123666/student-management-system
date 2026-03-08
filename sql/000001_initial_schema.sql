-- 初始数据库结构
-- 执行顺序：1

-- 创建学生表
CREATE TABLE IF NOT EXISTS `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '学号',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `gender` varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT '性别',
  `college` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '学院',
  `major` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '专业',
  `grade` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '年级',
  `class` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '班级',
  `phone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '联系电话',
  `is_focus` tinyint(1) NOT NULL DEFAULT '0' COMMENT '重点关注状态',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像URL',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_students_student_id` (`student_id`),
  KEY `idx_students_name` (`name`),
  KEY `idx_students_college` (`college`),
  KEY `idx_students_major` (`major`),
  KEY `idx_students_grade` (`grade`),
  KEY `idx_students_class` (`class`),
  KEY `idx_students_is_focus` (`is_focus`),
  KEY `idx_students_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='基础信息表';

-- 创建成绩表
CREATE TABLE IF NOT EXISTS `scores` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生学号',
  `course_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '课程名称',
  `credit` float NOT NULL COMMENT '学分',
  `mark` float NOT NULL COMMENT '分数',
  `term` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '学期',
  `is_failed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否挂科',
  PRIMARY KEY (`id`),
  KEY `idx_scores_student_id` (`student_id`),
  KEY `idx_scores_course_name` (`course_name`),
  KEY `idx_scores_mark` (`mark`),
  KEY `idx_scores_term` (`term`),
  KEY `idx_scores_is_failed` (`is_failed`),
  KEY `idx_scores_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='成绩表';

-- 创建荣誉表
CREATE TABLE IF NOT EXISTS `honors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `title` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '荣誉名称',
  `level` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '级别',
  `award_date` datetime NOT NULL COMMENT '获奖时间',
  PRIMARY KEY (`id`),
  KEY `idx_honors_student_id` (`student_id`),
  KEY `idx_honors_title` (`title`),
  KEY `idx_honors_level` (`level`),
  KEY `idx_honors_award_date` (`award_date`),
  KEY `idx_honors_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='荣誉表';

-- 创建考勤表
CREATE TABLE IF NOT EXISTS `attendances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `date` datetime NOT NULL COMMENT '考勤日期',
  `status` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '状态',
  `location` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '打卡位置/宿舍号',
  `description` text COLLATE utf8mb4_general_ci COMMENT '备注说明',
  PRIMARY KEY (`id`),
  KEY `idx_attendances_student_id` (`student_id`),
  KEY `idx_attendances_date` (`date`),
  KEY `idx_attendances_status` (`status`),
  KEY `idx_attendances_location` (`location`),
  KEY `idx_attendances_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='安全考勤表';

-- 创建赛事活动表
CREATE TABLE IF NOT EXISTS `competition_activities` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `type` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '赛事/活动名称',
  `role` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '担任角色/参赛身份',
  `date` datetime NOT NULL COMMENT '参与时间',
  `award` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '获奖情况/活动时长',
  PRIMARY KEY (`id`),
  KEY `idx_competition_activities_student_id` (`student_id`),
  KEY `idx_competition_activities_type` (`type`),
  KEY `idx_competition_activities_name` (`name`),
  KEY `idx_competition_activities_date` (`date`),
  KEY `idx_competition_activities_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='赛事与活动表';

-- 创建就业表
CREATE TABLE IF NOT EXISTS `employments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `status` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '就业状态',
  `company` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '签约公司/实习单位',
  `position` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位',
  PRIMARY KEY (`id`),
  KEY `idx_employments_student_id` (`student_id`),
  KEY `idx_employments_status` (`status`),
  KEY `idx_employments_company` (`company`),
  KEY `idx_employments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='就业表';

-- 创建心理健康表
CREATE TABLE IF NOT EXISTS `mental_healths` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `assessment_date` datetime NOT NULL COMMENT '评估日期',
  `level` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '心理评级',
  `counselor` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '负责辅导员/心理老师',
  `notes` text COLLATE utf8mb4_general_ci COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_mental_healths_student_id` (`student_id`),
  KEY `idx_mental_healths_assessment_date` (`assessment_date`),
  KEY `idx_mental_healths_level` (`level`),
  KEY `idx_mental_healths_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='心理健康表';

-- 创建预警表
CREATE TABLE IF NOT EXISTS `warnings` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联学生',
  `warning_type` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '预警类型',
  `level` varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT '预警等级',
  `description` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '预警详细原因说明',
  `status` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '处理状态',
  PRIMARY KEY (`id`),
  KEY `idx_warnings_student_id` (`student_id`),
  KEY `idx_warnings_warning_type` (`warning_type`),
  KEY `idx_warnings_level` (`level`),
  KEY `idx_warnings_status` (`status`),
  KEY `idx_warnings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='学生预警表';

-- 创建服务消息表
CREATE TABLE IF NOT EXISTS `service_messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '接收方学号',
  `title` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '推送标题',
  `content` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '推送内容',
  `is_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已读',
  PRIMARY KEY (`id`),
  KEY `idx_service_messages_student_id` (`student_id`),
  KEY `idx_service_messages_title` (`title`),
  KEY `idx_service_messages_is_read` (`is_read`),
  KEY `idx_service_messages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='服务推送表';

-- 创建树洞表
CREATE TABLE IF NOT EXISTS `tree_holes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '树洞内容',
  `background_color` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '卡片背景色',
  `likes` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  PRIMARY KEY (`id`),
  KEY `idx_tree_holes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='树洞表';

-- 创建用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `phone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `password` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `role` varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- 插入默认教师账号
INSERT IGNORE INTO `users` (`phone`, `password`, `role`) VALUES ('admin', '$2a$10$Qz5i4aW9YQ5J9e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e1e', 'teacher');
