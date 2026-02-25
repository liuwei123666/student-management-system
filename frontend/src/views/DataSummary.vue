<template>
  <div class="data-summary">
    <el-card class="summary-card">
      <template #header>
        <div class="card-header">
          <span>全校数据总结</span>
        </div>
      </template>
      
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="10" animated />
      </div>

      <div v-else-if="summary" class="charts-container">
        <!-- 学院学生人数分布饼图 -->
        <el-card class="chart-card">
          <template #header>
            <div class="chart-header">
              <span>各学院学生人数分布</span>
            </div>
          </template>
          <v-chart class="chart" :option="pieOption" />
        </el-card>

        <!-- 预警数量趋势柱状图 -->
        <el-card class="chart-card">
          <template #header>
            <div class="chart-header">
              <span>全校每月预警数量趋势</span>
            </div>
          </template>
          <v-chart class="chart" :option="barOption" />
        </el-card>
      </div>

      <div v-else-if="error" class="error-container">
        <el-alert
          title="数据加载失败"
          :description="error"
          type="error"
          show-icon
          center
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const summary = ref(null)
const loading = ref(false)
const error = ref('')

// 饼图配置
const pieOption = ref({
  title: {
    text: '各学院学生人数分布',
    left: 'center'
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    orient: 'vertical',
    left: 'left',
  },
  series: [
    {
      name: '学生人数',
      type: 'pie',
      radius: '50%',
      data: [],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
})

// 柱状图配置
const barOption = ref({
  title: {
    text: '全校每月预警数量趋势',
    left: 'center'
  },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['预警数量']
  },
  xAxis: {
    type: 'category',
    data: []
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      name: '预警数量',
      type: 'bar',
      data: [],
      itemStyle: {
        color: '#409EFF'
      }
    }
  ]
})

// 加载数据
const loadData = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await axios.get('http://localhost:8080/api/data-summary')
    summary.value = response.data
    updateCharts()
  } catch (err) {
    error.value = err.response?.data?.error || '数据加载失败'
  } finally {
    loading.value = false
  }
}

// 更新图表数据
const updateCharts = () => {
  if (!summary.value) return

  // 更新饼图数据
  pieOption.value.series[0].data = summary.value.college_distribution.map(item => ({
    name: item.college,
    value: item.count
  }))

  // 更新柱状图数据
  barOption.value.xAxis.data = summary.value.warning_trends.map(item => item.month)
  barOption.value.series[0].data = summary.value.warning_trends.map(item => item.count)
}

// 初始化
onMounted(() => {
  loadData()
})
</script>

<style scoped>
.data-summary {
  padding: 20px;
}

.summary-card {
  width: 100%;
}

.card-header {
  font-weight: bold;
  font-size: 18px;
}

.loading-container {
  margin: 20px 0;
}

.charts-container {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 20px;
}

.chart-card {
  margin-bottom: 20px;
}

.chart-header {
  font-weight: bold;
  font-size: 16px;
}

.chart {
  width: 100%;
  height: 400px;
}

.error-container {
  margin: 20px 0;
}

@media (max-width: 768px) {
  .chart {
    height: 300px;
  }
}
</style>
