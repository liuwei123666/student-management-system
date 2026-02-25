<template>
  <div class="warning-dashboard">
    <h2>学生预警一张表</h2>
    
    <!-- 统计卡片 -->
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="stat-card academic-card">
          <div class="card-content">
            <div class="card-title">学业告警数</div>
            <div class="card-value">{{ dashboardData.academicWarningCount || 0 }}</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card safety-card">
          <div class="card-content">
            <div class="card-title">安全考勤异常数</div>
            <div class="card-value">{{ dashboardData.safetyWarningCount || 0 }}</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card mental-card">
          <div class="card-content">
            <div class="card-title">心理关注数</div>
            <div class="card-value">{{ dashboardData.mentalWarningCount || 0 }}</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card total-card">
          <div class="card-content">
            <div class="card-title">综合预警数</div>
            <div class="card-value">{{ dashboardData.totalWarningCount || 0 }}</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 预警记录列表 -->
    <el-card class="warning-list-card" style="margin-top: 20px;">
      <template #header>
        <div class="card-header">
          <span>最近预警记录</span>
        </div>
      </template>
      <el-table :data="warnings" style="width: 100%">
        <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
        <el-table-column prop="warning_type" label="预警类型" width="120"></el-table-column>
        <el-table-column prop="level" label="预警等级" width="100"></el-table-column>
        <el-table-column prop="description" label="预警原因"></el-table-column>
        <el-table-column prop="status" label="处理状态" width="120"></el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180"></el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button type="primary" size="small" @click="openSendMessageDialog(scope.row)">发送服务推送</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 发送服务推送对话框 -->
    <el-dialog
      v-model="sendMessageDialogVisible"
      title="发送服务推送"
      width="500px"
    >
      <el-form :model="messageForm" label-width="80px">
        <el-form-item label="学号">
          <el-input v-model="messageForm.student_id" readonly></el-input>
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="messageForm.title"></el-input>
        </el-form-item>
        <el-form-item label="内容">
          <el-input type="textarea" v-model="messageForm.content" :rows="4"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="sendMessageDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="sendMessage">发送</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const dashboardData = ref({})
const warnings = ref([])
const sendMessageDialogVisible = ref(false)
const messageForm = ref({
  student_id: '',
  title: '',
  content: ''
})

// 获取预警统计概况
const getDashboardData = () => {
  fetch('http://localhost:8080/api/warnings/dashboard')
    .then(response => response.json())
    .then(data => {
      dashboardData.value = data
    })
    .catch(error => {
      console.error('获取预警统计数据失败:', error)
    })
}

// 获取预警列表
const getWarnings = () => {
  fetch('http://localhost:8080/api/warnings')
    .then(response => response.json())
    .then(data => {
      warnings.value = data.data
    })
    .catch(error => {
      console.error('获取预警列表失败:', error)
    })
}

// 打开发送消息对话框
const openSendMessageDialog = (row) => {
  messageForm.value.student_id = row.student_id
  messageForm.value.title = ''
  messageForm.value.content = ''
  sendMessageDialogVisible.value = true
}

// 发送服务消息
const sendMessage = () => {
  fetch('http://localhost:8080/api/messages', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      student_id: messageForm.value.student_id,
      title: messageForm.value.title,
      content: messageForm.value.content
    })
  })
    .then(response => response.json())
    .then(data => {
      ElMessage.success('消息发送成功')
      sendMessageDialogVisible.value = false
    })
    .catch(error => {
      console.error('发送消息失败:', error)
      ElMessage.error('消息发送失败')
    })
}

onMounted(() => {
  getDashboardData()
  getWarnings()
})
</script>

<style scoped>
.warning-dashboard {
  padding: 20px;
}

.stat-card {
  height: 150px;
}

.card-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.card-title {
  font-size: 16px;
  margin-bottom: 10px;
  color: #606266;
}

.card-value {
  font-size: 28px;
  font-weight: bold;
}

.academic-card .card-value {
  color: #409EFF;
}

.safety-card .card-value {
  color: #F56C6C;
}

.mental-card .card-value {
  color: #E6A23C;
}

.total-card .card-value {
  color: #67C23A;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.warning-list-card {
  margin-top: 20px;
}
</style>
