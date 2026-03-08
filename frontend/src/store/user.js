import { defineStore } from 'pinia'

// 定义用户状态管理store
export const useUserStore = defineStore('user', {
  state: () => ({
    // 从localStorage中获取初始状态
    token: localStorage.getItem('token') || '',
    phone: localStorage.getItem('phone') || '',
    role: localStorage.getItem('role') || ''
  }),
  getters: {
    // 判断用户是否已登录
    isLoggedIn: (state) => !!state.token,
    // 获取用户角色
    getUserRole: (state) => state.role
  },
  actions: {
    // 登录成功后设置用户信息
    setUserInfo(token, phone, role) {
      this.token = token
      this.phone = phone
      this.role = role
      
      // 持久化到localStorage
      localStorage.setItem('token', token)
      localStorage.setItem('phone', phone)
      localStorage.setItem('role', role)
    },
    // 登出，清除用户信息
    logout() {
      this.token = ''
      this.phone = ''
      this.role = ''
      
      // 从localStorage中清除
      localStorage.removeItem('token')
      localStorage.removeItem('phone')
      localStorage.removeItem('role')
    }
  }
})
