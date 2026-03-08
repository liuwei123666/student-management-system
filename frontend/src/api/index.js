// API服务，统一管理所有API请求

import { useUserStore } from '../store/user'
import router from '../router'
import { ElMessage } from 'element-plus'

// API基础URL
const API_BASE_URL = 'http://localhost:8080/api';

// 通用请求函数
async function request(url, options = {}) {
  try {
    // 获取用户store
    const userStore = useUserStore()
    
    // 请求拦截器：添加Authorization头
    const headers = {
      'Content-Type': 'application/json',
      ...options.headers,
    }
    
    // 如果有token，添加到请求头
    if (userStore.token) {
      headers['Authorization'] = `Bearer ${userStore.token}`
    }

    const response = await fetch(`${API_BASE_URL}${url}`, {
      ...options,
      headers,
    });

    // 响应拦截器：处理错误
    if (!response.ok) {
      // 处理401错误（未登录或Token失效）
      if (response.status === 401) {
        // 清除用户状态
        userStore.logout()
        // 跳转到登录页
        router.push('/login')
        throw new Error('未登录或登录已过期，请重新登录')
      }
      
      // 处理403错误（无权限）
      if (response.status === 403) {
        // 提示用户权限不足，不跳转登录页
        ElMessage.error('权限不足，无法访问该资源')
        throw new Error('权限不足')
      }
      
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error('API request error:', error);
    throw error;
  }
}

// 学生相关API
export const studentAPI = {
  // 获取学生列表
  getStudents: async (params = {}) => {
    const queryParams = new URLSearchParams();
    if (params.page) queryParams.append('page', params.page);
    if (params.pageSize) queryParams.append('pageSize', params.pageSize);
    if (params.search) queryParams.append('search', params.search);
    if (params.focusOnly) queryParams.append('focusOnly', params.focusOnly);

    const queryString = queryParams.toString();
    const url = queryString ? `/students?${queryString}` : '/students';
    return request(url);
  },

  // 获取单个学生详情
  getStudent: async (id) => {
    return request(`/students/${id}`);
  },

  // 创建新学生
  createStudent: async (studentData) => {
    return request('/students', {
      method: 'POST',
      body: JSON.stringify(studentData),
    });
  },

  // 更新学生信息
  updateStudent: async (id, studentData) => {
    return request(`/students/${id}`, {
      method: 'PUT',
      body: JSON.stringify(studentData),
    });
  },

  // 删除学生
  deleteStudent: async (id) => {
    return request(`/students/${id}`, {
      method: 'DELETE',
    });
  },

  // 获取学生数字画像
  getStudentPortrait: async (id) => {
    return request(`/students/${id}/portrait`);
  },
};

// 成绩相关API
export const scoreAPI = {
  // 获取成绩列表
  getScores: async (params = {}) => {
    const queryParams = new URLSearchParams();
    if (params.page) queryParams.append('page', params.page);
    if (params.pageSize) queryParams.append('pageSize', params.pageSize);
    if (params.student_id) queryParams.append('student_id', params.student_id);

    const queryString = queryParams.toString();
    const url = queryString ? `/scores?${queryString}` : '/scores';
    return request(url);
  },

  // 获取单个成绩详情
  getScore: async (id) => {
    return request(`/scores/${id}`);
  },

  // 创建新成绩
  createScore: async (scoreData) => {
    return request('/scores', {
      method: 'POST',
      body: JSON.stringify(scoreData),
    });
  },

  // 更新成绩信息
  updateScore: async (id, scoreData) => {
    return request(`/scores/${id}`, {
      method: 'PUT',
      body: JSON.stringify(scoreData),
    });
  },

  // 删除成绩
  deleteScore: async (id) => {
    return request(`/scores/${id}`, {
      method: 'DELETE',
    });
  },
};

// 考勤相关API
export const attendanceAPI = {
  // 获取考勤列表
  getAttendances: async (params = {}) => {
    const queryParams = new URLSearchParams();
    if (params.page) queryParams.append('page', params.page);
    if (params.pageSize) queryParams.append('pageSize', params.pageSize);
    if (params.student_id) queryParams.append('student_id', params.student_id);

    const queryString = queryParams.toString();
    const url = queryString ? `/attendances?${queryString}` : '/attendances';
    return request(url);
  },

  // 获取单个考勤详情
  getAttendance: async (id) => {
    return request(`/attendances/${id}`);
  },

  // 创建新考勤
  createAttendance: async (attendanceData) => {
    return request('/attendances', {
      method: 'POST',
      body: JSON.stringify(attendanceData),
    });
  },

  // 更新考勤信息
  updateAttendance: async (id, attendanceData) => {
    return request(`/attendances/${id}`, {
      method: 'PUT',
      body: JSON.stringify(attendanceData),
    });
  },

  // 删除考勤
  deleteAttendance: async (id) => {
    return request(`/attendances/${id}`, {
      method: 'DELETE',
    });
  },
};

// 赛事活动相关API
export const competitionAPI = {
  // 获取赛事活动列表
  getCompetitions: async (params = {}) => {
    const queryParams = new URLSearchParams();
    if (params.page) queryParams.append('page', params.page);
    if (params.pageSize) queryParams.append('pageSize', params.pageSize);
    if (params.student_id) queryParams.append('student_id', params.student_id);

    const queryString = queryParams.toString();
    const url = queryString ? `/competitions?${queryString}` : '/competitions';
    return request(url);
  },

  // 获取单个赛事活动详情
  getCompetition: async (id) => {
    return request(`/competitions/${id}`);
  },

  // 创建新赛事活动
  createCompetition: async (competitionData) => {
    return request('/competitions', {
      method: 'POST',
      body: JSON.stringify(competitionData),
    });
  },

  // 更新赛事活动信息
  updateCompetition: async (id, competitionData) => {
    return request(`/competitions/${id}`, {
      method: 'PUT',
      body: JSON.stringify(competitionData),
    });
  },

  // 删除赛事活动
  deleteCompetition: async (id) => {
    return request(`/competitions/${id}`, {
      method: 'DELETE',
    });
  },
};

// 就业相关API
export const employmentAPI = {
  // 获取就业列表
  getEmployments: async (params = {}) => {
    const queryParams = new URLSearchParams();
    if (params.page) queryParams.append('page', params.page);
    if (params.pageSize) queryParams.append('pageSize', params.pageSize);
    if (params.student_id) queryParams.append('student_id', params.student_id);

    const queryString = queryParams.toString();
    const url = queryString ? `/employments?${queryString}` : '/employments';
    return request(url);
  },

  // 获取单个就业详情
  getEmployment: async (id) => {
    return request(`/employments/${id}`);
  },

  // 创建新就业
  createEmployment: async (employmentData) => {
    return request('/employments', {
      method: 'POST',
      body: JSON.stringify(employmentData),
    });
  },

  // 更新就业信息
  updateEmployment: async (id, employmentData) => {
    return request(`/employments/${id}`, {
      method: 'PUT',
      body: JSON.stringify(employmentData),
    });
  },

  // 删除就业
  deleteEmployment: async (id) => {
    return request(`/employments/${id}`, {
      method: 'DELETE',
    });
  },
};

// 预警相关API
export const warningAPI = {
  // 获取预警统计概况
  getWarningDashboard: async () => {
    return request('/warnings/dashboard');
  },

  // 获取预警列表
  getWarnings: async () => {
    return request('/warnings');
  },
};

// 消息相关API
export const messageAPI = {
  // 发送服务消息
  sendMessage: async (messageData) => {
    return request('/messages', {
      method: 'POST',
      body: JSON.stringify(messageData),
    });
  },

  // 获取消息列表
  getMessages: async () => {
    return request('/messages');
  },
};

// 树洞相关API
export const treeholeAPI = {
  // 获取树洞帖子
  getTreeHolePosts: async () => {
    return request('/treehole');
  },

  // 创建树洞帖子
  createTreeHolePost: async (postData) => {
    return request('/treehole', {
      method: 'POST',
      body: JSON.stringify(postData),
    });
  },

  // 点赞树洞帖子
  likeTreeHolePost: async (id) => {
    return request(`/treehole/${id}/like`, {
      method: 'POST',
    });
  },
};

// 数据总结API
export const dataSummaryAPI = {
  // 获取数据总结
  getDataSummary: async () => {
    return request('/data-summary');
  },
};

// 认证相关API
export const authAPI = {
  // 注册
  register: async (userData) => {
    return request('/auth/register', {
      method: 'POST',
      body: JSON.stringify(userData),
    });
  },
  // 登录
  login: async (userData) => {
    return request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(userData),
    });
  },
};

export default {
  student: studentAPI,
  score: scoreAPI,
  attendance: attendanceAPI,
  competition: competitionAPI,
  employment: employmentAPI,
  warning: warningAPI,
  message: messageAPI,
  treehole: treeholeAPI,
  dataSummary: dataSummaryAPI,
  auth: authAPI,
};
