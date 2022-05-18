<template>
  <div>
    <el-card>
      <el-form style="width: 60%" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

        <el-form-item label="学号" prop="snumber">
          <el-input v-model="ruleForm.snumber" :value="ruleForm.snumber" ></el-input>
        </el-form-item>

        <el-form-item label="姓名" prop="sname">
          <el-input v-model="ruleForm.sname" :value="ruleForm.sname" ></el-input>
        </el-form-item>

        <el-form-item label="年龄" prop="age">
          <el-input v-model="ruleForm.age" :value="ruleForm.age" ></el-input>
        </el-form-item>

        <el-form-item label="性别" prop="sex">
          <el-select v-model="value1" style="width: 100%" @change="sexChange" placeholder="请选择">
                <el-option
                v-for="item in options1"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
              </el-select>
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

        <el-form-item label="密码" prop="password">
          <el-input v-model="ruleForm.password" :value="ruleForm.password" ></el-input>
        </el-form-item>


        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">添加学生</el-button>
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
        snumber: null,
        sname: null,
        age: null,
        sex: null,
        departmentID: null,
        password: null,
      },

      options1: [{
          value: '男',
          label: '男'
        }, {
          value: '女',
          label: '女',
        }
        ],
      value1: '',

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
    sexChange(sex) {
        const that = this
        console.log("下拉框更改了sex：" + sex)  // 测试
        this.ruleForm.sex = sex
        console.log("this.ruleForm.sex：" + this.ruleForm.sex)  // 测试
    },

    departmentChange(departmentID) {
        const that = this
        console.log("下拉框更改了departmentID：" + departmentID)  // 测试
        this.ruleForm.departmentID = departmentID
        console.log("this.ruleForm.departmentID：" + this.ruleForm.departmentID)  // 测试
    },

    

    async submitForm(formName) {
          const that = this
      // 第1步：获取添加学生必要的属性
          let newage = this.ruleForm.age
          console.log("测试age ：" + newage)  // 测试

          let newdepartmentID = this.ruleForm.departmentID
          console.log("测试departmentID ：" + newdepartmentID)  // 测试

          let newname = this.ruleForm.sname
          console.log("测试name ：" + newname)  // 测试

          let newnumber = this.ruleForm.snumber
          console.log("测试number ：" + newnumber)  // 测试

          let newpasword = this.ruleForm.password
          console.log("测试pasword ：" + newpasword)  // 测试

          let newsex = this.ruleForm.sex
          console.log("测试sex ：" + newsex)  // 测试
      
      // 第2步：通过api添加学生
          var config = {
            method: 'post',
            url: 'http://1.15.130.83:8080/api/v1/student',
            data : {
              age : parseInt(newage),
              departmentID : parseInt(newdepartmentID),
              name : newname,
              number : newnumber,
              password : newpasword, 
              sex : newsex, 
            },
            headers: { 'content-type': 'application/json',}
          };

          await axios(config).then(function (resp) {
            console.log("添加学生的resp.data.msg：" + resp.data.msg)  // 测试
            if (resp.data.code === 0) {
              that.$message({
                showClose: true,
                message: '添加学生成功',
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