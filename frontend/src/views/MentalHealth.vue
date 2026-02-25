<template>
  <div class="mental-health">
    <h2>心理健康评估记录</h2>
    
    <el-card class="mental-health-card">
      <template #header>
        <div class="card-header">
          <span>心理评估记录列表</span>
        </div>
      </template>
      <el-table :data="mentalHealthRecords" style="width: 100%">
        <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
        <el-table-column prop="assessment_date" label="评估日期" width="180"></el-table-column>
        <el-table-column prop="level" label="心理评级" width="120">
          <template #default="scope">
            <el-tag :type="getLevelType(scope.row.level)">{{ scope.row.level }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="counselor" label="负责辅导员/心理老师" width="200"></el-table-column>
        <el-table-column prop="notes" label="备注"></el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'MentalHealth',
  data() {
    return {
      mentalHealthRecords: []
    }
  },
  mounted() {
    this.getMentalHealthRecords()
  },
  methods: {
    // 获取心理健康评估记录
    getMentalHealthRecords() {
      // 这里需要实现获取心理健康评估记录的API调用
      // 由于后端可能还没有实现对应的API，暂时使用模拟数据
      this.mentalHealthRecords = [
        {
          id: 1,
          student_id: '2021001',
          assessment_date: '2026-02-20',
          level: '正常',
          counselor: '张老师',
          notes: '情绪稳定，状态良好',
          created_at: '2026-02-20 10:00:00'
        },
        {
          id: 2,
          student_id: '2021002',
          assessment_date: '2026-02-19',
          level: '关注',
          counselor: '李老师',
          notes: '近期压力较大，建议适当放松',
          created_at: '2026-02-19 14:30:00'
        },
        {
          id: 3,
          student_id: '2021003',
          assessment_date: '2026-02-18',
          level: '高危',
          counselor: '王老师',
          notes: '情绪波动较大，需要密切关注',
          created_at: '2026-02-18 09:15:00'
        }
      ]
    },
    // 根据心理评级获取标签类型
    getLevelType(level) {
      switch (level) {
        case '正常':
          return 'success'
        case '关注':
          return 'warning'
        case '高危':
          return 'danger'
        default:
          return 'info'
      }
    }
  }
}
</script>

<style scoped>
.mental-health {
  padding: 20px;
}

.mental-health-card {
  margin-top: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
