<template>
  <div>
    <el-card>
      <el-form style="width: 60%" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

        <el-form-item label="课程号" prop="cnumber">
          <el-input v-model="ruleForm.cnumber" :value="ruleForm.cnumber" ></el-input>
        </el-form-item>

        <el-form-item label="课程名" prop="cname">
          <el-input v-model="ruleForm.cname" :value="ruleForm.cname" ></el-input>
        </el-form-item>

        <el-form-item label="学分" prop="credit">
          <el-input v-model="ruleForm.credit" :value="ruleForm.credit" ></el-input>
        </el-form-item>

        <el-form-item label="院系" >
              <el-select v-model="value2" style="width: 100%" @change="departmentChange" placeholder="请选择">
                <el-option
                v-for="item in options2"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
              </el-select>
          </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">添加课程</el-button>
        </el-form-item>

      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    
    return {
      ruleForm: {
        cnumber: null,
        cname: null,
        credit: null,
        departmentID: null,
      },

      options2: [{
          value: '1',
          label: '计算机学院'
        }, {
          value: '2',
          label: '通信学院',
        }, {
          value: '3',
          label: '理学院'
        }, {
          value: '4',
          label: '外国语学院'
        }],

        value2: ''
    };
  },


  methods: {

    departmentChange(departmentID) {
        const that = this
        console.log("下拉框更改了departmentID：" + departmentID)  // 测试
        this.ruleForm.departmentID = departmentID
        console.log("this.ruleForm.departmentID：" + this.ruleForm.departmentID)  // 测试
    },

    

    async submitForm(formName) {
          const that = this
      // 第1步：获取添加课程必要的属性

          let newdepartmentID = this.ruleForm.departmentID
          console.log("测试departmentID ：" + newdepartmentID)  // 测试

          let newname = this.ruleForm.cname
          console.log("测试name ：" + newname)  // 测试

          let newcredit = this.ruleForm.credit
          console.log("测试name ：" + newcredit)  // 测试

          let newnumber = this.ruleForm.cnumber
          console.log("测试number ：" + newnumber)  // 测试
      
      // 第2步：通过api添加课程
          var config = {
            method: 'post',
            url: 'http://1.15.130.83:8080/api/v1/course',
            data : {
              credit : parseInt(newcredit),
              departmentID : parseInt(newdepartmentID),
              name : newname,
              number : newnumber,
            },
            headers: { 'content-type': 'application/json',}
          };

          await axios(config).then(function (resp) {
            console.log("添加课程的resp.data.msg：" + resp.data.msg)  // 测试
            if (resp.data.code === 0) {
              that.$message({
                showClose: true,
                message: '添加课程成功',
                type: 'success'
              });
              // window.location.reload()
            }
            else {
              that.$message({
                showClose: true,
                message: resp.data.msg,
                type: '错误！请检查输入信息或数据库'
              });
            }
          })
        
    }
    
  }
}
</script>