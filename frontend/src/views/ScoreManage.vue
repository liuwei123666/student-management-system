<template>
  <div class="score-manage">
    <h1>成绩管理</h1>
    <div class="action-bar">
      <el-button type="primary" @click="importScores">模拟导入成绩</el-button>
      <el-button type="success" @click="showAddDialog">添加成绩</el-button>
    </div>

    <!-- 添加/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
    >
      <el-form :model="formData" label-width="100px">
        <el-form-item label="学号">
          <el-input v-model="formData.student_id" placeholder="请输入学号" />
        </el-form-item>
        <el-form-item label="课程名称">
          <el-input v-model="formData.course_name" placeholder="请输入课程名称" />
        </el-form-item>
        <el-form-item label="学分">
          <el-input-number v-model="formData.credit" :min="0" :max="10" step="0.5" />
        </el-form-item>
        <el-form-item label="分数">
          <el-input-number v-model="formData.mark" :min="0" :max="100" step="1" />
        </el-form-item>
        <el-form-item label="学期">
          <el-input v-model="formData.term" placeholder="请输入学期" />
        </el-form-item>
        <el-form-item label="是否挂科">
          <el-switch v-model="formData.is_failed" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveScore">保存</el-button>
        </span>
      </template>
    </el-dialog>
    <el-table :data="scoreList" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80"></el-table-column>
      <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
      <el-table-column prop="course_name" label="课程名称" width="180"></el-table-column>
      <el-table-column prop="credit" label="学分" width="80"></el-table-column>
      <el-table-column prop="mark" label="分数" width="80"></el-table-column>
      <el-table-column prop="term" label="学期" width="120"></el-table-column>
      <el-table-column prop="is_failed" label="是否挂科" width="100">
        <template #default="scope">
          <span :class="{ 'failed': scope.row.is_failed }">
            {{ scope.row.is_failed ? '是' : '否' }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button size="small" type="primary" @click="editScore(scope.row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteScore(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';

const scoreList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const dialogVisible = ref(false);
const dialogTitle = ref('添加成绩');
const formData = ref({
  student_id: '',
  course_name: '',
  credit: 0,
  mark: 0,
  term: '',
  is_failed: false
});
const editingId = ref(null);

// 模拟导入成绩
const importScores = () => {
  ElMessage.success('导入成功');
};

// 显示添加对话框
const showAddDialog = () => {
  dialogTitle.value = '添加成绩';
  formData.value = {
    student_id: '',
    course_name: '',
    credit: 0,
    mark: 0,
    term: '',
    is_failed: false
  };
  editingId.value = null;
  dialogVisible.value = true;
};

// 显示编辑对话框
const editScore = (score) => {
  dialogTitle.value = '编辑成绩';
  formData.value = {
    student_id: score.student_id,
    course_name: score.course_name,
    credit: score.credit,
    mark: score.mark,
    term: score.term,
    is_failed: score.is_failed
  };
  editingId.value = score.ID;
  dialogVisible.value = true;
};

// 保存成绩
const saveScore = async () => {
  try {
    let response;
    if (editingId.value) {
      // 更新操作
      response = await fetch(`http://localhost:8080/api/scores/${editingId.value}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData.value)
      });
    } else {
      // 添加操作
      response = await fetch('http://localhost:8080/api/scores', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData.value)
      });
    }
    
    if (response.ok) {
      ElMessage.success(editingId.value ? '更新成功' : '添加成功');
      dialogVisible.value = false;
      loadScores();
    } else {
      ElMessage.error(editingId.value ? '更新失败' : '添加失败');
    }
  } catch (error) {
    console.error('Failed to save score:', error);
    ElMessage.error('操作失败');
  }
};

// 加载成绩列表
const loadScores = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/scores?page=${currentPage.value}&pageSize=${pageSize.value}`);
    const data = await response.json();
    scoreList.value = data.data;
    total.value = data.total;
  } catch (error) {
    console.error('Failed to load scores:', error);
    ElMessage.error('加载成绩失败');
  }
};

// 删除成绩
const deleteScore = async (id) => {
  try {
    const response = await fetch(`http://localhost:8080/api/scores/${id}`, {
      method: 'DELETE'
    });
    if (response.ok) {
      ElMessage.success('删除成功');
      loadScores();
    } else {
      ElMessage.error('删除失败');
    }
  } catch (error) {
    console.error('Failed to delete score:', error);
    ElMessage.error('删除失败');
  }
};

// 分页处理
const handleSizeChange = (size) => {
  pageSize.value = size;
  loadScores();
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
  loadScores();
};

// 组件挂载时加载数据
onMounted(() => {
  loadScores();
});
</script>

<style scoped>
.score-manage {
  padding: 20px;
}

.action-bar {
  margin-bottom: 20px;
  display: flex;
  gap: 10px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.failed {
  color: red;
  font-weight: bold;
}
</style>
