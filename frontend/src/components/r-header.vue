<template>
  <div class="flex justify-between">

    <div class="flex items-center">
      <div class="flex items-center space-x-2">
        <i class="el-icon-magic-stick"></i>
        <div>{{ currentTerm }}</div>
      </div>

      <el-divider direction="vertical" class="m-4"></el-divider>
      <div class="flex items-center space-x-2">
        <i class="el-icon-user"></i>
        <div>{{ name }}</div>
      </div>

    </div>
    <el-button icon="el-icon-switch-button" class="m-3" round @click="out()"></el-button>
  </div>
</template>

<script>
export default {
  name: "r-header",
  data() {
    return {
      name: null,
      currentTerm: null
    }
  },
  created() {
    this.name = sessionStorage.getItem("name")
    this.currentTerm = sessionStorage.getItem("currentTerm")
  },
  methods: {
    out() {
      sessionStorage.clear();
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
</style>