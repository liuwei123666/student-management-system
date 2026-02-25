<template>
  <div class="student-portrait">
    <el-card class="search-card">
      <div class="search-container">
        <el-input v-model="studentId" placeholder="请输入学号" style="width: 300px" />
        <el-button type="primary" @click="searchStudent" style="margin-left: 10px">搜索</el-button>
      </div>
    </el-card>

    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="10" animated />
    </div>

    <div v-else-if="portrait" class="portrait-container">
      <!-- 左侧基本资料 -->
      <div class="basic-info-card">
        <el-card class="avatar-card">
          <div class="avatar-container">
            <el-avatar :size="150" :src="portrait.basic_info.Avatar || defaultAvatar" />
            <h2>{{ portrait.basic_info.Name }}</h2>
            <p class="student-id">{{ portrait.basic_info.StudentID }}</p>
          </div>
          <el-divider />
          <div class="info-item">
            <span class="label">性别：</span>
            <span class="value">{{ portrait.basic_info.Gender }}</span>
          </div>
          <div class="info-item">
            <span class="label">学院：</span>
            <span class="value">{{ portrait.basic_info.College }}</span>
          </div>
          <div class="info-item">
            <span class="label">专业：</span>
            <span class="value">{{ portrait.basic_info.Major }}</span>
          </div>
          <div class="info-item">
            <span class="label">年级：</span>
            <span class="value">{{ portrait.basic_info.Grade }}</span>
          </div>
          <div class="info-item">
            <span class="label">班级：</span>
            <span class="value">{{ portrait.basic_info.Class }}</span>
          </div>
          <div class="info-item">
            <span class="label">联系电话：</span>
            <span class="value">{{ portrait.basic_info.Phone }}</span>
          </div>
          <el-divider />
          <div class="info-item">
            <span class="label">参与活动数：</span>
            <span class="value">{{ portrait.activity_count }}</span>
          </div>
          <div class="info-item">
            <span class="label">考勤统计：</span>
            <div class="attendance-stats">
              <span v-for="(count, status) in portrait.attendance_stats" :key="status" class="attendance-item">
                {{ status }}: {{ count }}
              </span>
            </div>
          </div>
        </el-card>
      </div>

      <!-- 右侧图表 -->
      <div class="charts-container">
        <!-- 雷达图 -->
        <el-card class="chart-card">
          <template #header>
            <div class="card-header">
              <span>能力分布雷达图</span>
            </div>
          </template>
          <v-chart class="chart" :option="radarOption" />
        </el-card>

        <!-- 折线图 -->
        <el-card class="chart-card">
          <template #header>
            <div class="card-header">
              <span>成绩趋势折线图</span>
            </div>
          </template>
          <v-chart class="chart" :option="lineOption" />
        </el-card>

        <!-- 预警记录 -->
        <el-card class="warning-card">
          <template #header>
            <div class="card-header">
              <span>预警记录</span>
            </div>
          </template>
          <el-table :data="portrait.warnings" style="width: 100%">
            <el-table-column prop="warning_type" label="预警类型" />
            <el-table-column prop="level" label="预警等级" />
            <el-table-column prop="description" label="预警原因" />
            <el-table-column prop="status" label="处理状态" />
          </el-table>
        </el-card>
      </div>
    </div>

    <div v-else-if="error" class="error-container">
      <el-alert
        title="搜索失败"
        :description="error"
        type="error"
        show-icon
        center
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const studentId = ref('')
const portrait = ref(null)
const loading = ref(false)
const error = ref('')
const defaultAvatar = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'

// 搜索学生
const searchStudent = async () => {
  if (!studentId.value) {
    error.value = '请输入学号'
    return
  }

  loading.value = true
  error.value = ''
  
  try {
    const response = await axios.get(`http://localhost:8080/api/students/${studentId.value}/portrait`)
    portrait.value = response.data
    updateCharts()
  } catch (err) {
    error.value = err.response?.data?.error || '搜索失败，请检查学号是否正确'
  } finally {
    loading.value = false
  }
}

// 雷达图配置
const radarOption = ref({
  title: {
    text: '能力分布'
  },
  tooltip: {},
  radar: {
    indicator: [
      { name: '学习成绩', max: 100 },
      { name: '考勤情况', max: 100 },
      { name: '活动参与', max: 100 },
      { name: '心理健康', max: 100 },
      { name: '综合表现', max: 100 }
    ]
  },
  series: [
    {
      type: 'radar',
      data: [
        {
          value: [0, 0, 0, 0, 0],
          name: '能力值'
        }
      ]
    }
  ]
})

// 折线图配置
const lineOption = ref({
  title: {
    text: '成绩趋势'
  },
  tooltip: {
    trigger: 'axis'
  },
  xAxis: {
    type: 'category',
    data: []
  },
  yAxis: {
    type: 'value',
    min: 0,
    max: 100
  },
  series: [
    {
      data: [],
      type: 'line'
    }
  ]
})

// 更新图表数据
const updateCharts = () => {
  if (!portrait.value) return

  // 计算平均成绩
  const scores = portrait.value.scores
  const avgScore = scores.length > 0 ? scores.reduce((sum, score) => sum + score.mark, 0) / scores.length : 0

  // 计算考勤率
  const attendanceStats = portrait.value.attendance_stats
  const totalAttendance = Object.values(attendanceStats).reduce((sum, count) => sum + count, 0)
  const normalAttendance = attendanceStats['正常'] || 0
  const attendanceRate = totalAttendance > 0 ? (normalAttendance / totalAttendance) * 100 : 0

  // 计算活动参与度
  const activityCount = portrait.value.activity_count
  const activityLevel = Math.min((activityCount / 10) * 100, 100) // 假设10个活动为满分

  // 更新雷达图
  radarOption.value.series[0].data[0].value = [
    avgScore,
    attendanceRate,
    activityLevel,
    80, // 假设心理健康默认值
    (avgScore + attendanceRate + activityLevel) / 3 // 综合表现
  ]

  // 处理成绩趋势数据
  const termScores = {}
  scores.forEach(score => {
    if (!termScores[score.term]) {
      termScores[score.term] = []
    }
    termScores[score.term].push(score.mark)
  })

  // 计算每个学期的平均成绩
  const terms = Object.keys(termScores).sort()
  const termAvgScores = terms.map(term => {
    const termScoreList = termScores[term]
    return termScoreList.reduce((sum, score) => sum + score, 0) / termScoreList.length
  })

  // 更新折线图
  lineOption.value.xAxis.data = terms
  lineOption.value.series[0].data = termAvgScores
}

// 初始化
onMounted(() => {
  // 可以在这里设置默认学号进行测试
  // studentId.value = '2023001'
  // searchStudent()
})
</script>

<style scoped>
.student-portrait {
  padding: 20px;
}

.search-card {
  margin-bottom: 20px;
}

.search-container {
  display: flex;
  align-items: center;
  justify-content: center;
}

.loading-container {
  margin: 20px 0;
}

.portrait-container {
  display: flex;
  gap: 20px;
  margin-top: 20px;
}

.basic-info-card {
  flex: 0 0 300px;
}

.avatar-card {
  width: 100%;
}

.avatar-container {
  text-align: center;
  margin-bottom: 20px;
}

.student-id {
  color: #666;
  margin-top: 5px;
}

.info-item {
  margin: 10px 0;
}

.label {
  font-weight: bold;
  margin-right: 10px;
  color: #666;
}

.value {
  color: #333;
}

.attendance-stats {
  margin-top: 5px;
}

.attendance-item {
  display: block;
  margin: 5px 0;
  color: #333;
}

.charts-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.chart-card {
  margin-bottom: 20px;
}

.chart {
  width: 100%;
  height: 400px;
}

.warning-card {
  margin-top: 20px;
}

.card-header {
  font-weight: bold;
  font-size: 16px;
}

.error-container {
  margin-top: 20px;
}

@media (max-width: 768px) {
  .portrait-container {
    flex-direction: column;
  }
  
  .basic-info-card {
    flex: 1;
  }
}
</style>
