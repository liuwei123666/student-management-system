<template>
  <div class="activity-manage">
    <h1>赛事活动管理</h1>
    <div class="action-bar">
      <el-button type="success" @click="showAddDialog">添加赛事活动</el-button>
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
        <el-form-item label="类型">
          <el-select v-model="formData.type" placeholder="请选择类型">
            <el-option label="竞赛" value="竞赛" />
            <el-option label="活动" value="活动" />
          </el-select>
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="formData.name" placeholder="请输入赛事/活动名称" />
        </el-form-item>
        <el-form-item label="担任角色">
          <el-input v-model="formData.role" placeholder="请输入担任角色" />
        </el-form-item>
        <el-form-item label="参与时间">
          <el-date-picker
            v-model="formData.date"
            type="datetime"
            placeholder="选择日期时间"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="获奖情况">
          <el-input v-model="formData.award" placeholder="请输入获奖情况/活动时长" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveActivity">保存</el-button>
        </span>
      </template>
    </el-dialog>
    <el-table :data="activityList" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80"></el-table-column>
      <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
      <el-table-column prop="type" label="类型" width="100"></el-table-column>
      <el-table-column prop="name" label="赛事/活动名称" width="200"></el-table-column>
      <el-table-column prop="role" label="担任角色" width="150"></el-table-column>
      <el-table-column prop="date" label="参与时间" width="180">
        <template #default="scope">
          {{ new Date(scope.row.date).toLocaleString() }}
        </template>
      </el-table-column>
      <el-table-column prop="award" label="获奖情况" width="180"></el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button size="small" type="primary" @click="editActivity(scope.row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteActivity(scope.row.id)">删除</el-button>
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

const activityList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const dialogVisible = ref(false);
const dialogTitle = ref('添加赛事活动');
const formData = ref({
  student_id: '',
  type: '竞赛',
  name: '',
  role: '',
  date: new Date(),
  award: ''
});
const editingId = ref(null);

// 显示添加对话框
const showAddDialog = () => {
  dialogTitle.value = '添加赛事活动';
  formData.value = {
    student_id: '',
    type: '竞赛',
    name: '',
    role: '',
    date: new Date(),
    award: ''
  };
  editingId.value = null;
  dialogVisible.value = true;
};

// 显示编辑对话框
const editActivity = (activity) => {
  dialogTitle.value = '编辑赛事活动';
  formData.value = {
    student_id: activity.student_id,
    type: activity.type,
    name: activity.name,
    role: activity.role,
    date: new Date(activity.date),
    award: activity.award
  };
  editingId.value = activity.ID;
  dialogVisible.value = true;
};

// 保存赛事活动
const saveActivity = async () => {
  try {
    let response;
    if (editingId.value) {
      // 更新操作
      response = await fetch(`http://localhost:8080/api/competitions/${editingId.value}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData.value)
      });
    } else {
      // 添加操作
      response = await fetch('http://localhost:8080/api/competitions', {
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
      loadActivities();
    } else {
      ElMessage.error(editingId.value ? '更新失败' : '添加失败');
    }
  } catch (error) {
    console.error('Failed to save activity:', error);
    ElMessage.error('操作失败');
  }
};

// 加载赛事活动列表
const loadActivities = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/competitions?page=${currentPage.value}&pageSize=${pageSize.value}`);
    const data = await response.json();
    activityList.value = data.data;
    total.value = data.total;
  } catch (error) {
    console.error('Failed to load activities:', error);
    ElMessage.error('加载赛事活动失败');
  }
};

// 删除赛事活动
const deleteActivity = async (id) => {
  try {
    const response = await fetch(`http://localhost:8080/api/competitions/${id}`, {
      method: 'DELETE'
    });
    if (response.ok) {
      ElMessage.success('删除成功');
      loadActivities();
    } else {
      ElMessage.error('删除失败');
    }
  } catch (error) {
    console.error('Failed to delete activity:', error);
    ElMessage.error('删除失败');
  }
};

// 分页处理
const handleSizeChange = (size) => {
  pageSize.value = size;
  loadActivities();
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
  loadActivities();
};

// 组件挂载时加载数据
onMounted(() => {
  loadActivities();
});
</script>

<style scoped>
.activity-manage {
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
</style>
