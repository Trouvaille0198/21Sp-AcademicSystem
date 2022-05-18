<template>
  <div>
    <el-container>

      <el-main>
        <div class="flex justify-center items-center h-screen">
          <el-card class="rounded-lg shadow-md px-10 py-3" shadow="hover">


            <div class="text-3xl font-bold italic text-primary text-center ">
              <!-- <i class="el-icon-s-home" style="margin-right: 25px"></i> -->

              选课管理系统

            </div>


            <div class="mt-8">
              <el-form class="p-0" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
                <el-form-item label="账号" prop="number">
                  <el-input v-model="ruleForm.number" placeholder="请输入账号" prefix-icon="el-icon-lollipop"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                  <el-input v-model="ruleForm.password" placeholder="请输入密码" show-password
                    prefix-icon="el-icon-ice-cream-round"></el-input>
                </el-form-item>
                <el-form-item label="类型" prop="type">
                  <el-radio-group v-model="ruleForm.type">
                    <el-radio label="student" value="student">学生</el-radio>
                    <el-radio label="teacher" value="teacher">老师</el-radio>
                    <el-radio label="admin" value="admin">管理员</el-radio>
                  </el-radio-group>
                </el-form-item>
                <el-form-item class="">
                  <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
                  <el-button @click="resetForm('ruleForm')">重置</el-button>
                  <!-- <el-button @click="test('ruleForm')">test</el-button> -->
                </el-form-item>
              </el-form>
            </div>
          </el-card>
        </div>
      </el-main>
    </el-container>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ruleForm: {
        number: null, // 学号
        password: null, // 密码
        type: null, // 用户类型  admin、student、teacher
      },
      rules: {
        number: [
          { required: true, message: '请输入用户 number', trigger: 'blur' },
          // { type: 'number', message: '请输入数字', trigger: 'blur' },  // 加这里是让账号不是0开头的
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '请选择', trigger: 'change' }
        ],
      }
    };
  },
  methods: {
    submitForm(formName) {
      const that = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let check = 7
          let name = null

          // axios.get('http://localhost:16/info/getCurrentTerm').then(function (resp) {
          //   sessionStorage.setItem("currentTerm", resp.data)
          // })
          sessionStorage.setItem("currentTerm", "21-冬季学期")    // 暂时写死
          sessionStorage.setItem("ForbidCourseSelection", false)  // 暂时写死
          // axios.get('http://localhost:16/info/getForbidCourseSelection').then(function (resp) {
          //   sessionStorage.setItem("ForbidCourseSelection", resp.data)
          // })

          if (that.ruleForm.type === 'admin') {
            let form = new FormData();
            form.append("number", this.ruleForm.number)
            form.append("password", this.ruleForm.password)

            console.log(form)
            axios.post("http://1.15.130.83:8080/api/v1/login/admin", form).then(function (resp) {
              console.log("管理员登陆验证信息：" + resp.data.code)
              check = resp.data.code
              if (check === 0) {
                name = "admin"

                sessionStorage.setItem("token", 'true')
                sessionStorage.setItem("type", that.ruleForm.type)
                sessionStorage.setItem("name", name)
                sessionStorage.setItem("number", "123")




                console.log('name: ' + name + ' ' + that.ruleForm.type + ' ' + "123")

                if (that.ruleForm.type === 'admin' && name === 'admin') {
                  that.$message({
                    showClose: true,
                    message: '登陆成功，欢迎 ' + name + '!',
                    type: 'success'
                  });
                  that.$router.push('/admin')
                }
                else {
                  that.$message({
                    showClose: true,
                    message: 'admin 登陆失败，检查登陆类型',
                    type: 'error'
                  });
                }
              }
              else {
                that.$message({
                  showClose: true,
                  message: '登陆失败，检查账号密码',
                  type: 'error'
                });
              }
            })
          }



          else if (that.ruleForm.type === 'teacher') {
            let form = new FormData();
            form.append("number", this.ruleForm.number)
            form.append("password", this.ruleForm.password)

            console.log(form)
            axios.post("http://1.15.130.83:8080/api/v1/login/teacher", form).then(function (resp) {
              console.log("教师登陆验证信息：" + resp.data.code)
              check = resp.data.code  // 值为0或7  ：0 true  7 false
              console.log("登录状态:" + resp.data.msg)  // 测试
              let id = resp.data.data.ID  // 这里一定是两个data，因为是从resp的data中读取名为data的字段
              console.log("登录ID:" + id)  // 测试
              if (check === 0) {
                axios.get("http://1.15.130.83:8080/api/v1/teacher/" + id).then(function (resp) {
                  console.log("登陆页正在获取用户信息" + resp.data.data.name)
                  name = resp.data.data.name
                  console.log("测试session的name:" + name)  // 测试
                  sessionStorage.setItem("id", id)
                  sessionStorage.setItem("token", 'true')
                  sessionStorage.setItem("type", that.ruleForm.type)
                  sessionStorage.setItem("name", name)
                  sessionStorage.setItem("number", resp.data.data.number)

                  if (that.ruleForm.type === 'teacher' && name !== 'admin') {
                    that.$message({
                      showClose: true,
                      message: '登陆成功，欢迎 ' + name + '!',
                      type: 'success'
                    });
                    that.$router.push('/teacher')
                  }
                  else {
                    that.$message({
                      showClose: true,
                      message: 'teacher 登陆失败，检查登陆类型',
                      type: 'error'
                    });
                  }
                })
              }
              else {
                that.$message({
                  showClose: true,
                  message: '登陆失败，检查账号密码',
                  type: 'error'
                });
              }
            })
          }


          else if (that.ruleForm.type === 'student') {
            console.log("测试：准备登录")   // 测试
            let form = new FormData();
            form.append("number", this.ruleForm.number)
            form.append("password", this.ruleForm.password)
            console.log(this.ruleForm.number)   // 测试
            console.log(this.ruleForm.password)   // 测试

            console.log(form.number)   // 测试为undefined  FormData对象只有一个.append()方法,它不会向其追加属性,而只存储提供的数据
            console.log(form.password)   // 测试为undefined  FormData对象只有一个.append()方法,它不会向其追加属性,而只存储提供的数据

            axios.post("http://1.15.130.83:8080/api/v1/login/student", form).then(function (resp) {
              console.log("学生登陆验证信息：" + resp.data.code)
              check = resp.data.code  // 值为0或7  ：0 true  7 false
              console.log("登录状态:" + resp.data.msg)  // 测试
              let id = resp.data.data.ID  // 这里一定是两个data，因为是从resp的data中读取名为data的字段
              console.log("登录ID:" + id)  // 测试
              if (check === 0) {
                axios.get("http://1.15.130.83:8080/api/v1/student/" + id).then(function (resp) {
                  console.log("登陆页正在获取用户信息" + resp.data.data.name)
                  name = resp.data.data.name
                  console.log("测试session的name:" + name)  // 测试
                  sessionStorage.setItem("id", id)
                  sessionStorage.setItem("token", 'true')
                  sessionStorage.setItem("type", that.ruleForm.type)
                  sessionStorage.setItem("name", name)
                  sessionStorage.setItem("number", resp.data.data.number)

                  that.$message({
                    showClose: true,
                    message: '登陆成功，欢迎 ' + name + '!',
                    type: 'success'
                  });

                  console.log('正在跳转：' + '/' + that.ruleForm.type)

                  // 3. 路由跳转
                  that.$router.push({
                    path: '/' + that.ruleForm.type,
                    query: {}
                  })
                })
              }
              else {
                that.$message({
                  showClose: true,
                  message: '账号密码错误，请联系管理员',
                  type: 'error'
                });
              }
            })
          }
          else {
            console.log("! error type")
          }
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    test(forName) {
      console.log(this.ruleForm)
    }
  }
}
</script>

<style>
.login-module {
  /*width: 380px;*/
  /*height: 325px;*/
  margin-top: 60px;
  /*border: none;*/
  position: absolute;
  right: 500px;
  text-align: center;
  width: 30%;
}

.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}
</style>