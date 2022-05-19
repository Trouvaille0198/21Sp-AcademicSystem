<template>
  <div>
    <el-form style="width: 60%" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="课程名" prop="name">
        <el-input v-model="ruleForm.cname"></el-input>
      </el-form-item>
      <el-form-item label="课程号" prop="number">
        <el-input v-model="ruleForm.cnumber"></el-input>
      </el-form-item>
      <el-form-item label="学分" prop="credit">
        <el-input v-model="ruleForm.ccredit"></el-input>
      </el-form-item>
      <el-form-item label="学院">
        <el-select v-model="value" placeholder="请选择">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
        <el-button @click="test">test</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ruleForm: {
        name: null,
        number: null,
        credit: null,
        departmentID: null
      },
      rules: {
        // name: [
        //   { required: true, message: '请输入名称', trigger: 'blur' },
        // ],
        // credit: [
        //   // { required: true, message: '请输入学分', trigger: 'change' },
        //   { required: true, type: 'number', message: '请输入数字', trigger: 'blur' },
        // ],
        // number: [
        //   { required: true, message: '请输入学号', trigger: 'blur' },
        // ],
        // departmentID: [
        //   { required: true },
        // ]
      },
      options: [{
          value: '选项1',
          label: '计算机学院'
        }, {
          value: '选项2',
          label: '理学院'
        }, {
          value: '选项3',
          label: '外国语学院'
        }, {
          value: '选项4',
          label: '通信学院'
        }],
        value: ''
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          // 通过前端校验
          const that = this
          console.log(this.ruleForm)
          if (options.value === '计算机学院') 
              this.ruleForm.departmentID = 1;
          else if (options.value === '理学院')
              this.ruleForm.departmentID = 2;
          else if (options.value === '外国语学院')
              this.ruleForm.departmentID = 3;
          else
              this.ruleForm.departmentID = 4;
          axios.post("http://1.15.130.83:8080/api/v1/course", this.ruleForm).then(function (resp) {
            console.log(resp)
            if (resp.data === true) {
              that.$message({
                showClose: true,
                message: '插入成功',
                type: 'success'
              });
            }
            else {
              that.$message.error('插入失败，请检查数据库t');
            }
            that.$router.push("/courseList")
          })
        }
        else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    test() {
      console.log(this.ruleForm)
    }
  }
}
</script>