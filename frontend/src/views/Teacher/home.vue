<template>
  <div>
    <el-container>
      <el-main>
        <info-card></info-card>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Logout from "@/components/logout";
import InfoCard from "@/components/infoCard";
export default {
  name: "home",
  components: { InfoCard, Logout },
}
</script>

<style scoped>
</style>