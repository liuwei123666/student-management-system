<template>
  <div class="attendance-manage">
    <h1>安全考勤管理</h1>
    <div class="action-bar">
      <el-button type="success" @click="showAddDialog">添加考勤</el-button>
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
        <el-form-item label="考勤日期">
          <el-date-picker
            v-model="formData.date"
            type="datetime"
            placeholder="选择日期时间"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="formData.status" placeholder="请选择状态">
            <el-option label="正常" value="正常" />
            <el-option label="迟到" value="迟到" />
            <el-option label="缺勤" value="缺勤" />
          </el-select>
        </el-form-item>
        <el-form-item label="打卡位置">
          <el-input v-model="formData.location" placeholder="请输入打卡位置" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formData.description" placeholder="请输入备注" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveAttendance">保存</el-button>
        </span>
      </template>
    </el-dialog>
    <el-table :data="attendanceList" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80"></el-table-column>
      <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
      <el-table-column prop="date" label="考勤日期" width="180">
        <template #default="scope">
          {{ new Date(scope.row.date).toLocaleString() }}
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="120">
        <template #default="scope">
          <el-tag :type="getStatusType(scope.row.status)">
            {{ scope.row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="location" label="打卡位置" width="180"></el-table-column>
      <el-table-column prop="description" label="备注" width="200"></el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button size="small" type="primary" @click="editAttendance(scope.row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteAttendance(scope.row.id)">删除</el-button>
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

const attendanceList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const dialogVisible = ref(false);
const dialogTitle = ref('添加考勤');
const formData = ref({
  student_id: '',
  date: new Date(),
  status: '正常',
  location: '',
  description: ''
});
const editingId = ref(null);

// 获取状态对应的标签类型
const getStatusType = (status) => {
  switch (status) {
    case '正常':
      return 'success';
    case '迟到':
      return 'warning';
    case '缺勤':
      return 'danger';
    default:
      return '';
  }
};

// 显示添加对话框
const showAddDialog = () => {
  dialogTitle.value = '添加考勤';
  formData.value = {
    student_id: '',
    date: new Date(),
    status: '正常',
    location: '',
    description: ''
  };
  editingId.value = null;
  dialogVisible.value = true;
};

// 显示编辑对话框
const editAttendance = (attendance) => {
  dialogTitle.value = '编辑考勤';
  formData.value = {
    student_id: attendance.student_id,
    date: new Date(attendance.date),
    status: attendance.status,
    location: attendance.location,
    description: attendance.description
  };
  editingId.value = attendance.ID;
  dialogVisible.value = true;
};

// 保存考勤
const saveAttendance = async () => {
  try {
    let response;
    if (editingId.value) {
      // 更新操作
      response = await fetch(`http://localhost:8080/api/attendances/${editingId.value}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData.value)
      });
    } else {
      // 添加操作
      response = await fetch('http://localhost:8080/api/attendances', {
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
      loadAttendances();
    } else {
      ElMessage.error(editingId.value ? '更新失败' : '添加失败');
    }
  } catch (error) {
    console.error('Failed to save attendance:', error);
    ElMessage.error('操作失败');
  }
};

// 加载考勤列表
const loadAttendances = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/attendances?page=${currentPage.value}&pageSize=${pageSize.value}`);
    const data = await response.json();
    attendanceList.value = data.data;
    total.value = data.total;
  } catch (error) {
    console.error('Failed to load attendances:', error);
    ElMessage.error('加载考勤失败');
  }
};

// 删除考勤
const deleteAttendance = async (id) => {
  try {
    const response = await fetch(`http://localhost:8080/api/attendances/${id}`, {
      method: 'DELETE'
    });
    if (response.ok) {
      ElMessage.success('删除成功');
      loadAttendances();
    } else {
      ElMessage.error('删除失败');
    }
  } catch (error) {
    console.error('Failed to delete attendance:', error);
    ElMessage.error('删除失败');
  }
};

// 分页处理
const handleSizeChange = (size) => {
  pageSize.value = size;
  loadAttendances();
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
  loadAttendances();
};

// 组件挂载时加载数据
onMounted(() => {
  loadAttendances();
});
</script>

<style scoped>
.attendance-manage {
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
