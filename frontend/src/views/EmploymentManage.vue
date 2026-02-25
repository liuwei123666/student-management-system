<template>
  <div class="employment-manage">
    <h1>就业管理</h1>
    <div class="action-bar">
      <el-button type="success" @click="showAddDialog">添加就业信息</el-button>
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
        <el-form-item label="就业状态">
          <el-select v-model="formData.status" placeholder="请选择就业状态">
            <el-option label="已就业" value="已就业" />
            <el-option label="未就业" value="未就业" />
            <el-option label="实习" value="实习" />
          </el-select>
        </el-form-item>
        <el-form-item label="公司/单位">
          <el-input v-model="formData.company" placeholder="请输入签约公司/实习单位" />
        </el-form-item>
        <el-form-item label="岗位">
          <el-input v-model="formData.position" placeholder="请输入岗位" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveEmployment">保存</el-button>
        </span>
      </template>
    </el-dialog>
    <el-table :data="employmentList" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80"></el-table-column>
      <el-table-column prop="student_id" label="学号" width="120"></el-table-column>
      <el-table-column prop="status" label="就业状态" width="120"></el-table-column>
      <el-table-column prop="company" label="签约公司/实习单位" width="200"></el-table-column>
      <el-table-column prop="position" label="岗位" width="180"></el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button size="small" type="primary" @click="editEmployment(scope.row)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteEmployment(scope.row.id)">删除</el-button>
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

const employmentList = ref([]);
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const dialogVisible = ref(false);
const dialogTitle = ref('添加就业信息');
const formData = ref({
  student_id: '',
  status: '已就业',
  company: '',
  position: ''
});
const editingId = ref(null);

// 显示添加对话框
const showAddDialog = () => {
  dialogTitle.value = '添加就业信息';
  formData.value = {
    student_id: '',
    status: '已就业',
    company: '',
    position: ''
  };
  editingId.value = null;
  dialogVisible.value = true;
};

// 显示编辑对话框
const editEmployment = (employment) => {
  dialogTitle.value = '编辑就业信息';
  formData.value = {
    student_id: employment.student_id,
    status: employment.status,
    company: employment.company,
    position: employment.position
  };
  editingId.value = employment.ID;
  dialogVisible.value = true;
};

// 保存就业信息
const saveEmployment = async () => {
  try {
    let response;
    if (editingId.value) {
      // 更新操作
      response = await fetch(`http://localhost:8080/api/employments/${editingId.value}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData.value)
      });
    } else {
      // 添加操作
      response = await fetch('http://localhost:8080/api/employments', {
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
      loadEmployments();
    } else {
      ElMessage.error(editingId.value ? '更新失败' : '添加失败');
    }
  } catch (error) {
    console.error('Failed to save employment:', error);
    ElMessage.error('操作失败');
  }
};

// 加载就业列表
const loadEmployments = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/employments?page=${currentPage.value}&pageSize=${pageSize.value}`);
    const data = await response.json();
    employmentList.value = data.data;
    total.value = data.total;
  } catch (error) {
    console.error('Failed to load employments:', error);
    ElMessage.error('加载就业信息失败');
  }
};

// 删除就业信息
const deleteEmployment = async (id) => {
  try {
    const response = await fetch(`http://localhost:8080/api/employments/${id}`, {
      method: 'DELETE'
    });
    if (response.ok) {
      ElMessage.success('删除成功');
      loadEmployments();
    } else {
      ElMessage.error('删除失败');
    }
  } catch (error) {
    console.error('Failed to delete employment:', error);
    ElMessage.error('删除失败');
  }
};

// 分页处理
const handleSizeChange = (size) => {
  pageSize.value = size;
  loadEmployments();
};

const handleCurrentChange = (current) => {
  currentPage.value = current;
  loadEmployments();
};

// 组件挂载时加载数据
onMounted(() => {
  loadEmployments();
});
</script>

<style scoped>
.employment-manage {
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
