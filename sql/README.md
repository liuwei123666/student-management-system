# 数据库变更管理

## 概述

本目录用于管理学生管理系统的数据库变更。团队成员在进行数据库结构或数据变更时，需要按照本规范进行操作，以确保数据库变更能够正确同步给其他团队成员。

## 变更管理规范

### 1. 数据库结构变更

当需要修改数据库结构时（如添加表、修改字段等），请按照以下步骤操作：

1. **创建变更文件**：在 `sql` 目录下创建一个新的 SQL 文件，文件名格式为 `{序号}_{描述}.sql`，例如 `000002_add_new_field.sql`。

2. **编写变更脚本**：在文件中编写 SQL 语句，确保使用 `IF NOT EXISTS` 等语句保证脚本的幂等性。

3. **更新模型文件**：在 `backend/models` 目录中更新相应的模型文件。

4. **更新迁移代码**：在 `backend/config/database.go` 的 `AutoMigrate` 函数中添加新模型或修改现有模型。

5. **提交代码**：将变更文件和相关代码提交到 GitHub。

6. **通知团队**：告知团队成员需要执行数据库变更操作。

### 2. 种子数据变更

当需要添加或修改种子数据时，请按照以下步骤操作：

1. **更新种子数据代码**：在 `backend/config/seed.go` 文件中添加或修改种子数据逻辑。

2. **提交代码**：将变更提交到 GitHub。

3. **通知团队**：告知团队成员需要执行种子数据初始化操作。

## 执行变更

### 1. 首次初始化数据库

对于新环境，需要执行初始数据库结构脚本：

```bash
# 登录 MySQL
mysql -u student_sys -p student_system

# 执行初始脚本
source /path/to/student-management-system/sql/000001_initial_schema.sql;
```

### 2. 执行后续变更

当有新的数据库变更时，按照以下步骤执行：

1. **拉取代码**：从 GitHub 拉取最新代码。

2. **执行变更脚本**：按照序号顺序执行新的变更脚本。

3. **启动应用**：应用启动时会自动执行 GORM 的自动迁移和种子数据初始化。

## 注意事项

1. **备份数据库**：在执行数据库变更前，确保对数据库进行备份。

2. **测试变更**：在本地环境测试变更脚本，确保它们能正确执行。

3. **保持顺序**：按照序号顺序执行变更脚本，确保变更的正确性。

4. **幂等性**：确保变更脚本是幂等的，多次执行不会导致错误。

5. **文档化**：详细记录每一次数据库变更的内容和原因。

## 示例

### 添加新表示例

```sql
-- 添加新表
-- 执行顺序：2

CREATE TABLE IF NOT EXISTS `new_table` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_new_table_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='新表';
```

### 修改字段示例

```sql
-- 修改字段
-- 执行顺序：3

ALTER TABLE `students` ADD COLUMN IF NOT EXISTS `new_field` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '新字段';
```
