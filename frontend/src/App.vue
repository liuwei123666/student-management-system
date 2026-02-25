<template>
  <div class="app-container">
    <!-- 侧边栏 -->
    <el-aside :width="isSidebarCollapsed ? '64px' : '200px'" class="sidebar" :class="{ collapsed: isSidebarCollapsed }">
      <div class="logo">
        <h2 v-if="!isSidebarCollapsed">学生管理系统</h2>
      </div>
      <el-menu
        :default-active="activeMenu"
        class="sidebar-menu"
        router
      >
        <el-menu-item index="/">
          <el-icon><House /></el-icon>
          <span v-if="!isSidebarCollapsed">首页</span>
        </el-menu-item>
        <el-menu-item index="/basic-info">
          <el-icon><User /></el-icon>
          <span v-if="!isSidebarCollapsed">基础信息</span>
        </el-menu-item>
        <el-menu-item index="/student-management">
          <el-icon><User /></el-icon>
          <span v-if="!isSidebarCollapsed">学生管理</span>
        </el-menu-item>
        <el-menu-item index="/score-manage">
          <el-icon><Document /></el-icon>
          <span v-if="!isSidebarCollapsed">成绩管理</span>
        </el-menu-item>
        <el-menu-item index="/attendance-manage">
          <el-icon><Timer /></el-icon>
          <span v-if="!isSidebarCollapsed">安全考勤</span>
        </el-menu-item>
        <el-menu-item index="/activity-manage">
          <el-icon><Trophy /></el-icon>
          <span v-if="!isSidebarCollapsed">赛事活动</span>
        </el-menu-item>
        <el-menu-item index="/employment-manage">
          <el-icon><Briefcase /></el-icon>
          <span v-if="!isSidebarCollapsed">就业管理</span>
        </el-menu-item>
        <el-menu-item index="/focus-students">
          <el-icon><Star /></el-icon>
          <span v-if="!isSidebarCollapsed">重点关注</span>
        </el-menu-item>
        <el-menu-item index="/warning-dashboard">
          <el-icon><Warning /></el-icon>
          <span v-if="!isSidebarCollapsed">学生预警一张表</span>
        </el-menu-item>
        <el-menu-item index="/mental-health">
          <el-icon><Message /></el-icon>
          <span v-if="!isSidebarCollapsed">心理健康</span>
        </el-menu-item>
        <el-menu-item index="/student-portrait">
          <el-icon><Avatar /></el-icon>
          <span v-if="!isSidebarCollapsed">数字画像</span>
        </el-menu-item>
        <el-menu-item index="/data-summary">
          <el-icon><DataAnalysis /></el-icon>
          <span v-if="!isSidebarCollapsed">数据总结</span>
        </el-menu-item>
        <el-menu-item index="/tree-hole">
          <el-icon><ChatLineRound /></el-icon>
          <span v-if="!isSidebarCollapsed">树洞</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 主内容区 -->
    <div class="main-container">
      <!-- 顶部栏 -->
      <el-header class="header">
        <div class="header-left">
          <el-icon @click="toggleSidebar"><Menu /></el-icon>
        </div>
        <div class="header-right">
          <span>欢迎，管理员</span>
        </div>
      </el-header>

      <!-- 内容区 -->
      <el-main class="content">
        <router-view />
      </el-main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { House, User, Document, Menu, Timer, Trophy, Briefcase, Warning, Message, Star, DataAnalysis, Avatar, ChatLineRound } from '@element-plus/icons-vue'

const route = useRoute()
const isSidebarCollapsed = ref(false)

const activeMenu = computed(() => {
  return route.path
})

const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value
}

// 响应式处理
const handleResize = () => {
  if (window.innerWidth < 768) {
    isSidebarCollapsed.value = true
  } else {
    isSidebarCollapsed.value = false
  }
}

onMounted(() => {
  handleResize()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app-container {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

.sidebar {
  background-color: #2c3e50;
  color: white;
  transition: width 0.3s;
  overflow: hidden;
}

.sidebar.collapsed {
  width: 64px;
}

.logo {
  padding: 20px;
  text-align: center;
  border-bottom: 1px solid #34495e;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo h2 {
  font-size: 18px;
  margin: 0;
  transition: opacity 0.3s;
}

.sidebar-menu {
  height: calc(100vh - 60px);
  background-color: #2c3e50;
}

.sidebar-menu :deep(.el-menu-item) {
  color: white;
  height: 60px;
  line-height: 60px;
}

.sidebar-menu :deep(.el-menu-item.is-active) {
  background-color: #34495e;
  color: #409eff;
}

.main-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.header {
  height: 60px;
  background-color: white;
  border-bottom: 1px solid #e0e0e0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-left {
  font-size: 20px;
  cursor: pointer;
}

.header-right {
  font-size: 14px;
  color: #666;
}

.content {
  flex: 1;
  background-color: #f5f7fa;
  padding: 20px;
  overflow-y: auto;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .app-container {
    flex-direction: column;
  }

  .sidebar {
    position: fixed;
    top: 60px;
    left: 0;
    height: calc(100vh - 60px);
    z-index: 1000;
    transform: translateX(-100%);
  }

  .sidebar:not(.collapsed) {
    transform: translateX(0);
  }

  .main-container {
    width: 100%;
  }

  .header {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1001;
  }

  .content {
    margin-top: 60px;
    padding: 10px;
  }
}
</style>
