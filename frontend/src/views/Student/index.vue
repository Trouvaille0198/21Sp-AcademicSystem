<template>
  <div>
    <el-container>
      <r-aside></r-aside>
      <el-container>
        <el-header class="bg-white">
          <r-header></r-header>
        </el-header>
        <el-container>
          <el-main>
            <router-view></router-view>
          </el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import RAside from "@/components/r-aside";
import RHeader from "@/components/r-header";
export default {
  name: "index",
  components: { RHeader, RAside },
}
</script>

<style scoped>
</style>