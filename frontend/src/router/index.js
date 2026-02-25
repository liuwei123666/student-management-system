import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/basic-info',
    name: 'BasicInfo',
    component: () => import('../views/BasicInfo.vue')
  },
  {
    path: '/student-management',
    name: 'StudentManagement',
    component: () => import('../views/StudentManage.vue')
  },
  {
    path: '/focus-students',
    name: 'FocusStudents',
    component: () => import('../views/StudentManage.vue'),
    meta: { focusOnly: true }
  },
  {
    path: '/score-management',
    name: 'ScoreManagement',
    component: () => import('../views/ScoreManagement.vue')
  },
  {
    path: '/score-manage',
    name: 'ScoreManage',
    component: () => import('../views/ScoreManage.vue')
  },
  {
    path: '/attendance-manage',
    name: 'AttendanceManage',
    component: () => import('../views/AttendanceManage.vue')
  },
  {
    path: '/activity-manage',
    name: 'ActivityManage',
    component: () => import('../views/ActivityManage.vue')
  },
  {
    path: '/employment-manage',
    name: 'EmploymentManage',
    component: () => import('../views/EmploymentManage.vue')
  },
  {
    path: '/warning-dashboard',
    name: 'WarningDashboard',
    component: () => import('../views/WarningDashboard.vue')
  },
  {
    path: '/mental-health',
    name: 'MentalHealth',
    component: () => import('../views/MentalHealth.vue')
  },
  {
    path: '/student-portrait',
    name: 'StudentPortrait',
    component: () => import('../views/StudentPortrait.vue')
  },
  {
    path: '/data-summary',
    name: 'DataSummary',
    component: () => import('../views/DataSummary.vue')
  },
  {
    path: '/tree-hole',
    name: 'TreeHole',
    component: () => import('../views/TreeHole.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
