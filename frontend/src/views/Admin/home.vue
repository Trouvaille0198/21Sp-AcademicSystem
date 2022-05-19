<template>
  <div>
    <el-container>
      <el-main>
        <div class="p-9 shadow-md rounded-lg bg-white">
          <div class="text-lg semibold mb-4">用户信息
          </div>
          <div class="flex-col space-y-3">
            <div><i class="el-icon-s-management mr-2"></i>账号：admin</div>
            <div><i class="el-icon-s-custom mr-2"></i>姓名：admin</div>
            <div><i class="el-icon-s-ticket mr-2"></i>类型：
              <el-tag>admin</el-tag>
            </div>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Logout from "@/components/logout";
export default {
  name: "home",
  components: { Logout },
}
</script>

<style scoped>
</style>