<template>
  <div class="p-9 shadow-md rounded-lg bg-white">
    <div class="text-lg semibold mb-4">用户信息
    </div>
    <div class="flex-col space-y-3">
      <div><i class="el-icon-s-management mr-2"></i>账号：{{ number }}</div>
      <div><i class="el-icon-s-custom mr-2"></i>姓名：{{ name }}</div>
      <div><i class="el-icon-s-ticket mr-2"></i>类型：
        <el-tag>{{ type }}</el-tag>
      </div>
    </div>
  </div>
  <!-- <el-card>
    <el-descriptions title="用户信息">
      <el-descriptions-item label="用户账号">{{ number }}</el-descriptions-item>
      <el-descriptions-item label="用户名">{{ name }}</el-descriptions-item>
      <el-descriptions-item label="用户类型">
        <el-tag size="small">{{ type }}</el-tag>
      </el-descriptions-item>
    </el-descriptions>
  </el-card> -->
</template>

<script>
export default {
  name: "infoCard",
  data() {
    return {
      number: sessionStorage.getItem("number"),
      name: sessionStorage.getItem("name"),
      type: sessionStorage.getItem("type")
    }
  }
}
</script>

<style scoped>
</style>