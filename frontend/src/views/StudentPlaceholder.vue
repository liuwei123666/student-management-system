<template>
  <div class="student-placeholder">
    <h1>学生端功能建设中，敬请期待</h1>
    <el-button type="primary" @click="handleLogout">退出登录</el-button>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '../store/user'

const router = useRouter()
const userStore = useUserStore()

const handleLogout = () => {
  // 清除用户状态
  userStore.logout()
  // 提示退出成功
  ElMessage.success('退出成功')
  // 跳转到登录页
  router.push('/login')
}
</script>

<style scoped>
.student-placeholder {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f7fa;
}

.student-placeholder h1 {
  font-size: 24px;
  color: #303133;
  margin-bottom: 40px;
  text-align: center;
}
</style>
