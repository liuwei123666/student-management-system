<template>
  <div class="tree-hole">
    <!-- 发布区域 -->
    <el-card class="publish-card">
      <template #header>
        <div class="card-header">
          <span>匿名发布</span>
        </div>
      </template>
      <div class="publish-container">
        <el-input
          v-model="newPost.content"
          type="textarea"
          :rows="4"
          placeholder="写下你的想法..."
          maxlength="500"
          show-word-limit
        />
        <div class="publish-actions">
          <el-button type="primary" @click="publishPost" :loading="publishing">
            匿名发布
          </el-button>
        </div>
      </div>
    </el-card>

    <!-- 帖子列表 -->
    <div class="posts-container">
      <el-row :gutter="20">
        <el-col :xs="24" :sm="12" :md="8" v-for="post in posts" :key="post.id">
          <el-card 
            class="post-card" 
            :style="{ backgroundColor: post.background_color }"
          >
            <div class="post-content">{{ post.content }}</div>
            <div class="post-footer">
              <span class="post-date">{{ formatDate(post.created_at) }}</span>
              <el-button 
                type="text" 
                icon="el-icon-star-on" 
                @click="likePost(post.id)"
                :loading="likingPosts.includes(post.id)"
              >
                {{ post.likes }}
              </el-button>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <!-- 加载更多 -->
    <div v-if="loading" class="loading-more">
      <el-skeleton :rows="3" animated />
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-container">
      <el-alert
        title="加载失败"
        :description="error"
        type="error"
        show-icon
        center
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { treeholeAPI } from '../api/index'

const posts = ref([])
const loading = ref(false)
const publishing = ref(false)
const error = ref('')
const newPost = ref({ content: '' })
const likingPosts = ref([])

// 加载帖子
const loadPosts = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await treeholeAPI.getTreeHolePosts()
    posts.value = response
  } catch (err) {
    error.value = err.message || '加载失败'
  } finally {
    loading.value = false
  }
}

// 发布帖子
const publishPost = async () => {
  if (!newPost.value.content.trim()) {
    return
  }

  publishing.value = true
  
  try {
    const response = await treeholeAPI.createTreeHolePost({
      content: newPost.value.content
    })
    // 将新帖子添加到列表顶部
    posts.value.unshift(response)
    // 清空输入框
    newPost.value.content = ''
  } catch (err) {
    error.value = err.message || '发布失败'
  } finally {
    publishing.value = false
  }
}

// 点赞帖子
const likePost = async (postId) => {
  likingPosts.value.push(postId)
  
  try {
    const response = await treeholeAPI.likeTreeHolePost(postId)
    // 更新帖子点赞数
    const index = posts.value.findIndex(post => post.id === postId)
    if (index !== -1) {
      posts.value[index] = response
    }
  } catch (err) {
    error.value = err.message || '点赞失败'
  } finally {
    likingPosts.value = likingPosts.value.filter(id => id !== postId)
  }
}

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleString()
}

// 初始化
onMounted(() => {
  loadPosts()
})
</script>

<style scoped>
.tree-hole {
  padding: 20px;
}

.publish-card {
  margin-bottom: 30px;
}

.card-header {
  font-weight: bold;
  font-size: 16px;
}

.publish-container {
  margin-top: 10px;
}

.publish-actions {
  margin-top: 15px;
  display: flex;
  justify-content: flex-end;
}

.posts-container {
  margin-top: 20px;
}

.post-card {
  margin-bottom: 20px;
  transition: all 0.3s ease;
}

.post-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.post-content {
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 15px;
  word-break: break-word;
}

.post-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  color: #666;
}

.post-date {
  flex: 1;
}

.loading-more {
  margin: 20px 0;
}

.error-container {
  margin: 20px 0;
}

@media (max-width: 768px) {
  .tree-hole {
    padding: 10px;
  }
  
  .post-card {
    margin-bottom: 15px;
  }
}
</style>
