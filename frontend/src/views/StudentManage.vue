<template>
  <div class="student-manage-container">
    <el-page-header>
      <template #title>学生管理</template>
    </el-page-header>

    <el-card class="student-manage-card">
      <!-- 搜索和添加按钮 -->
      <div class="search-and-add">
        <el-input
          v-model="searchQuery"
          placeholder="按学号或姓名搜索"
          clearable
          prefix-icon="el-icon-search"
          style="width: 300px; margin-right: 10px"
          @keyup.enter="loadStudents"
        />
        <el-button type="primary" @click="openAddDialog">
          <el-icon><Plus /></el-icon>
          新增学生
        </el-button>
      </div>

      <!-- 学生列表 -->
      <el-table
        v-loading="loading"
        :data="students"
        style="width: 100%"
        @row-click="handleRowClick"
      >
        <el-table-column prop="StudentID" label="学号" width="180" />
        <el-table-column prop="Name" label="姓名" width="120" />
        <el-table-column prop="Gender" label="性别" width="80" />
        <el-table-column prop="College" label="学院" />
        <el-table-column prop="Major" label="专业" />
        <el-table-column prop="Grade" label="年级" width="100" />
        <el-table-column prop="Class" label="班级" />
        <el-table-column prop="Phone" label="联系电话" />
        <el-table-column prop="IsFocus" label="重点关注" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.IsFocus" type="danger">是</el-tag>
            <el-tag v-else>否</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="openEditDialog(scope.row)">
              <el-icon><Edit /></el-icon>
              编辑
            </el-button>
            <el-button size="small" type="danger" @click="deleteStudent(scope.row)">
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 新增/编辑学生对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
    >
      <el-form :model="form" :rules="rules" ref="formRef">
        <el-form-item label="学号" prop="StudentID">
          <el-input v-model="form.StudentID" placeholder="请输入学号" />
        </el-form-item>
        <el-form-item label="姓名" prop="Name">
          <el-input v-model="form.Name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="性别" prop="Gender">
          <el-select v-model="form.Gender" placeholder="请选择性别">
            <el-option label="男" value="男" />
            <el-option label="女" value="女" />
          </el-select>
        </el-form-item>
        <el-form-item label="学院" prop="College">
          <el-input v-model="form.College" placeholder="请输入学院" />
        </el-form-item>
        <el-form-item label="专业" prop="Major">
          <el-input v-model="form.Major" placeholder="请输入专业" />
        </el-form-item>
        <el-form-item label="年级" prop="Grade">
          <el-input v-model="form.Grade" placeholder="请输入年级" />
        </el-form-item>
        <el-form-item label="班级" prop="Class">
          <el-input v-model="form.Class" placeholder="请输入班级" />
        </el-form-item>
        <el-form-item label="联系电话" prop="Phone">
          <el-input v-model="form.Phone" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="重点关注">
          <el-switch v-model="form.IsFocus" />
        </el-form-item>
        <el-form-item label="头像URL" prop="Avatar">
          <el-input v-model="form.Avatar" placeholder="请输入头像URL" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveStudent">保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, computed } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Plus, Edit, Delete } from '@element-plus/icons-vue'

// 获取路由信息
const route = useRoute()

// 响应式数据
const students = ref([])
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchQuery = ref('')
const dialogVisible = ref(false)
const dialogTitle = ref('新增学生')
const form = reactive({
  StudentID: '',
  Name: '',
  Gender: '',
  College: '',
  Major: '',
  Grade: '',
  Class: '',
  Phone: '',
  IsFocus: false,
  Avatar: ''
})
const formRef = ref(null)

// 表单验证规则
const rules = {
  StudentID: [
    { required: true, message: '请输入学号', trigger: 'blur' }
  ],
  Name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  Gender: [
    { required: true, message: '请选择性别', trigger: 'change' }
  ],
  College: [
    { required: true, message: '请输入学院', trigger: 'blur' }
  ],
  Major: [
    { required: true, message: '请输入专业', trigger: 'blur' }
  ],
  Grade: [
    { required: true, message: '请输入年级', trigger: 'blur' }
  ],
  Class: [
    { required: true, message: '请输入班级', trigger: 'blur' }
  ],
  Phone: [
    { required: true, message: '请输入联系电话', trigger: 'blur' }
  ]
}

// 加载学生数据
const loadStudents = async () => {
  loading.value = true
  try {
    // 构建查询参数
    const params = new URLSearchParams()
    params.append('page', currentPage.value)
    params.append('pageSize', pageSize.value)
    params.append('search', searchQuery.value)
    
    // 检查是否需要过滤重点关注学生
    if (route.meta.focusOnly) {
      params.append('focusOnly', 'true')
    }
    
    const response = await fetch(`http://localhost:8080/api/students?${params.toString()}`)
    if (!response.ok) throw new Error('Failed to load students')
    const data = await response.json()
    students.value = data.data
    total.value = data.total
  } catch (error) {
    ElMessage.error('加载学生数据失败')
    console.error('Error loading students:', error)
  } finally {
    loading.value = false
  }
}

// 打开新增对话框
const openAddDialog = () => {
  // 重置表单
  Object.keys(form).forEach(key => {
    form[key] = ''
  })
  form.IsFocus = false
  dialogTitle.value = '新增学生'
  dialogVisible.value = true
}

// 打开编辑对话框
const openEditDialog = (row) => {
  // 复制数据到表单
  Object.assign(form, row)
  dialogTitle.value = '编辑学生'
  dialogVisible.value = true
}

// 保存学生
const saveStudent = async () => {
  if (!formRef.value) return
  
  // 验证表单
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    try {
      let response
      if (form.ID) {
        // 编辑学生
        response = await fetch(`http://localhost:8080/api/students/${form.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(form)
        })
      } else {
        // 新增学生
        response = await fetch('http://localhost:8080/api/students', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(form)
        })
      }
      
      if (!response.ok) throw new Error('Failed to save student')
      
      ElMessage.success(form.ID ? '学生信息更新成功' : '学生添加成功')
      dialogVisible.value = false
      loadStudents()
    } catch (error) {
      ElMessage.error('保存失败，请重试')
      console.error('Error saving student:', error)
    }
  })
}

// 删除学生
const deleteStudent = async (row) => {
  if (!row.ID) return
  
  try {
    const response = await fetch(`http://localhost:8080/api/students/${row.ID}`, {
      method: 'DELETE'
    })
    
    if (!response.ok) throw new Error('Failed to delete student')
    
    ElMessage.success('学生删除成功')
    loadStudents()
  } catch (error) {
    ElMessage.error('删除失败，请重试')
    console.error('Error deleting student:', error)
  }
}

// 处理分页大小变化
const handleSizeChange = (size) => {
  pageSize.value = size
  currentPage.value = 1
  loadStudents()
}

// 处理当前页变化
const handleCurrentChange = (current) => {
  currentPage.value = current
  loadStudents()
}

// 处理行点击
const handleRowClick = (row) => {
  console.log('Row clicked:', row)
}

// 初始加载
onMounted(() => {
  loadStudents()
})
</script>

<style scoped>
.student-manage-container {
  padding: 20px;
}

.student-manage-card {
  margin-top: 20px;
}

.search-and-add {
  margin-bottom: 20px;
  display: flex;
  align-items: center;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>