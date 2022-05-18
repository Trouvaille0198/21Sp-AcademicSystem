<template>
  <div>
    <el-card>
      <el-form style="width: 60%" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

        <el-form-item label="课程号" prop="cnumber">
          <el-input v-model="ruleForm.cnumber" :value="ruleForm.cnumber" ></el-input>
        </el-form-item>

        <el-form-item label="学期" >
              <el-select v-model="value" style="width: 100%" @change="courseTermChange" placeholder="请选择">
                <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
              </el-select>
          </el-form-item>

        <el-form-item label="学号" prop="snumber">
          <el-input v-model="ruleForm.snumber" :value="ruleForm.snumber" ></el-input>
        </el-form-item>

        

        <el-form-item label="平时分" prop="usualScore">
          <el-input v-model.number="ruleForm.usualScore" :value="ruleForm.usualScore"></el-input>
        </el-form-item>

        <el-form-item label="考试分" prop="examScore">
          <el-input v-model.number="ruleForm.examScore" :value="ruleForm.examScore"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
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
        term:null, 
        snumber: null,
        usualScore: null,
        examScore: null,

        sid: null,
        cid: null,
        tname: null,
        
      },

      

      options: [{
          value: '21-秋季学期',
          label: '21-秋季学期'
        }, {
          value: '21-冬季学期',
          label: '21-冬季学期',
        }, {
          value: '22-春季学期',
          label: '22-春季学期'
        }],

        value: ''
    };
  },


  methods: {
    courseTermChange(changeTerm) {
        const that = this
        console.log("下拉框更改了Term：" + changeTerm)  // 测试
        this.ruleForm.term = changeTerm
        console.log("this.ruleForm.term：" + this.ruleForm.term)  // 测试
    },

    

    async submitForm(formName) {
      // 第1步：通过课程号和学期找到offeredCourseID
          let tid = sessionStorage.getItem("id");
          console.log("测试teacher id ：" + tid)  // 测试
          let courseNumber = this.ruleForm.cnumber
          console.log("测试courseNumber ：" + courseNumber)  // 测试
          let courseTerm = this.ruleForm.term
          console.log("测试courseTerm ：" + courseTerm)  // 测试
          let courseID = 0
          let studentNumber = this.ruleForm.snumber
          let studentID = 0
          let uscore = this.ruleForm.usualScore
          let escore = this.ruleForm.examScore

          await axios.get('http://1.15.130.83:8080/api/v1/teacher/' + tid + '/offered_course').then(function (resp) {
            console.log("测试获取教师所开的课 resp.data.code：" + resp.data.code)  // 测试
            for (let i = 0; i < resp.data.data.length; i++) {
              if (resp.data.data[i].number === courseNumber && resp.data.data[i].term === courseTerm) 
              {
                courseID = resp.data.data[i].id
                break
              }
            }
          })
      // 第2步：通过offeredCourseID找到班级（多个学生）
          console.log("测试course id ：" + courseID)  // 测试
          const that = this
            // 获取班级信息
          await axios.get('http://1.15.130.83:8080/api/v1/offered_course/' + courseID + '/student').then(function (resp) {
              console.log("测试获取班级信息resp.data.msg：" + resp.data.msg)  // 测试
      // 第3步：通过学号在多个学生中找到学生ID
              for (let i = 0; i < resp.data.data.length; i++) {
                if (resp.data.data[i].number === studentNumber) {
                  studentID = resp.data.data[i].id
                  break
                }
              }
          })
          console.log("测试获取的学生ID：" + studentID)  // 测试
          console.log("uscore：" + uscore)  // 测试
          console.log("escore：" + escore)  // 测试
          
      
      // 第4步：通过学生ID、offeredCourseID、平时分和考试分修改学生成绩
          var config = {
            method: 'put',
            url: 'http://1.15.130.83:8080/api/v1/selection',
            data : {
              courseID : parseInt(courseID),
              examScore : parseInt(uscore),
              studentID : parseInt(studentID),
              usualScore : parseInt(escore)
            },
            headers: { 'content-type': 'application/json',}
          };

          await axios(config).then(function (resp) {
            console.log("更改分数：courseID api：" + courseID)  // 测试
            console.log("更改分数的resp.data.msg：" + resp.data.msg)  // 测试
            if (resp.data.code === 0) {
              that.$message({
                showClose: true,
                message: '编辑成绩成功',
                type: 'success'
              });
              // window.location.reload()
            }
            else {
              that.$message({
                showClose: true,
                message: resp.data.msg,
                type: '错误！请检查输入信息，或联系管理员'
              });
            }
          })
        
    }
    
  }
}
</script>