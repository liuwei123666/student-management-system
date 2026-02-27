import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '../store/user'
import { ElMessage } from 'element-plus'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/basic-info',
    name: 'BasicInfo',
    component: () => import('../views/BasicInfo.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/student-management',
    name: 'StudentManagement',
    component: () => import('../views/StudentManage.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/focus-students',
    name: 'FocusStudents',
    component: () => import('../views/StudentManage.vue'),
    meta: { requiresAuth: true, role: 'teacher', focusOnly: true }
  },
  {
    path: '/score-management',
    name: 'ScoreManagement',
    component: () => import('../views/ScoreManagement.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/score-manage',
    name: 'ScoreManage',
    component: () => import('../views/ScoreManage.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/attendance-manage',
    name: 'AttendanceManage',
    component: () => import('../views/AttendanceManage.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/activity-manage',
    name: 'ActivityManage',
    component: () => import('../views/ActivityManage.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/employment-manage',
    name: 'EmploymentManage',
    component: () => import('../views/EmploymentManage.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/warning-dashboard',
    name: 'WarningDashboard',
    component: () => import('../views/WarningDashboard.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/mental-health',
    name: 'MentalHealth',
    component: () => import('../views/MentalHealth.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/student-portrait',
    name: 'StudentPortrait',
    component: () => import('../views/StudentPortrait.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/data-summary',
    name: 'DataSummary',
    component: () => import('../views/DataSummary.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/tree-hole',
    name: 'TreeHole',
    component: () => import('../views/TreeHole.vue'),
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/student',
    name: 'Student',
    component: () => import('../views/StudentPlaceholder.vue'),
    meta: { requiresAuth: true, role: 'student' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  // 检查路由是否需要认证
  if (to.meta.requiresAuth) {
    // 检查是否有token
    if (!userStore.token) {
      // 无token，跳转到登录页
      next('/login')
      return
    }
    
    // 检查角色是否匹配
    if (to.meta.role && to.meta.role !== userStore.role) {
      // 角色不匹配，提示无权限并跳转到对应首页
      ElMessage.error('权限不足，无法访问该资源')
      if (userStore.role === 'teacher') {
        next('/')
      } else if (userStore.role === 'student') {
        next('/student')
      } else {
        next('/login')
      }
      return
    }
  } else if (to.path === '/login') {
    // 已登录用户访问登录页，根据角色重定向到对应首页
    if (userStore.token) {
      if (userStore.role === 'teacher') {
        next('/')
      } else if (userStore.role === 'student') {
        next('/student')
      } else {
        next()
      }
      return
    }
  }
  
  next()
})

export default router
