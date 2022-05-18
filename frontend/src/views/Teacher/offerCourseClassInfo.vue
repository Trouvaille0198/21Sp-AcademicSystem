<template>
  <div>

    <div class="grid grid-cols-5">
      <div v-for="(item, index) in totalCourses" class="m-4">
        <div class="cursor-pointer hover:bg-blue-50 p-5 shadow-md rounded-lg bg-white flex-col space-y-3"
          @click="get_offered_course_info_by_button(item.number, item.term)">
          <div>{{ item.number }}</div>
          <div>{{ item.name }}</div>
          <div>{{ item.term }}</div>
        </div>
      </div>
    </div>

    <!-- <el-container>
      <el-main>
        <div class="p-9 shadow-md rounded-lg bg-white">
          <el-form :inline="true" :model="ruleForm" :rules="rules" ref="ruleForm" class="demo-ruleForm">
            <el-form-item label="课程号" prop="cnumber">
              <el-input v-model.number="ruleForm.cnumber"></el-input>
            </el-form-item>
            <el-form-item label="学期">
              <el-select v-model="value" style="" @change="courseTermChange" placeholder="请选择">
                <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="get_offered_course_info('ruleForm')">查看班级情况</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-main>
    </el-container> -->

    <div class="m-5 p-9 shadow-md rounded-lg bg-white">

      <template>
        <el-table :data="tableData" border show-header stripe style="width: 100%">
          <el-table-column prop="number" label="学号" width="150">
          </el-table-column>
          <el-table-column prop="name" label="姓名" width="150">
          </el-table-column>
          <el-table-column prop="sex" label="性别" width="150">
          </el-table-column>
          <el-table-column prop="department_name" label="学院" width="150">
          </el-table-column>
          <el-table-column prop="usual_score" label="平时分" width="150">
          </el-table-column>
          <el-table-column prop="exam_score" label="考试分" width="150">
          </el-table-column>
        </el-table>
      </template>
    </div>

  </div>
</template>
<script>

export default {
  data() {
    return {
      ruleForm: {
        cnumber: null,
        term: null,
      },

      tableData: null,
      tmpList: null,
      totalCourses: [{
        "id": 3,
        "name": "算法设计与分析",
        "number": "j0001",
        "credit": 4,
        "teacher_name": "计算机老师A",
        "department": "计算机学院",
        "term": "21-秋季学期"
      },],

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

  mounted: function () {
    let tid = sessionStorage.getItem("id");
    let that = this
    axios.get('http://1.15.130.83:8080/api/v1/teacher/' + tid + '/offered_course').then(function (resp) {
      that.totalCourses = resp.data.data
    })
  },

  methods: {

    courseTermChange(changeTerm) {
      const that = this
      console.log("下拉框更改了Term：" + changeTerm)  // 测试
      this.ruleForm.term = changeTerm
      console.log("this.ruleForm.term：" + this.ruleForm.term)  // 测试

    },

    async get_offered_course_info(formName) {
      let tid = sessionStorage.getItem("id");
      console.log("测试teacher id ：" + tid)  // 测试
      let courseNumber = this.ruleForm.cnumber
      console.log("测试courseNumber ：" + courseNumber)  // 测试
      let courseTerm = this.ruleForm.term
      console.log("测试courseTerm ：" + courseTerm)  // 测试
      let courseID = 0
      await axios.get('http://1.15.130.83:8080/api/v1/teacher/' + tid + '/offered_course').then(function (resp) {
        console.log("测试获取教师所开的课 resp.data.code：" + resp.data.code)  // 测试
        for (let i = 0; i < resp.data.data.length; i++) {
          if (resp.data.data[i].number === courseNumber && resp.data.data[i].term === courseTerm) {
            courseID = resp.data.data[i].id
            break
          }
        }
      })

      console.log("测试course id ：" + courseID)  // 测试
      if (courseID === 0) {
        this.tableData = null
      }
      else {
        const that = this
        // 获取班级信息
        await axios.get('http://1.15.130.83:8080/api/v1/offered_course/' + courseID + '/student').then(function (resp) {
          console.log("测试获取班级信息resp.data.msg：" + resp.data.msg)  // 测试
          console.log("测试班级信息resp.data.data[1].name：" + resp.data.data[1].name)  // 测试
          // that.tmpList = resp.data.data
          // that.tableData = that.tmpList
          that.tableData = resp.data.data
        })
      }
    },

    async get_offered_course_info_by_button(number, term) {
      let tid = sessionStorage.getItem("id");
      let courseID = 0
      this.tableData = []
      await axios.get('http://1.15.130.83:8080/api/v1/teacher/' + tid + '/offered_course').then(function (resp) {
        console.log("测试获取教师所开的课 resp.data.code：" + resp.data.code)  // 测试
        for (let i = 0; i < resp.data.data.length; i++) {
          if (resp.data.data[i].number === number && resp.data.data[i].term === term) {
            courseID = resp.data.data[i].id
            break
          }
        }
      })

      console.log("测试course id ：" + courseID)  // 测试
      if (courseID === 0) {
        this.tableData = null
      }
      else {
        const that = this
        // 获取班级信息
        await axios.get('http://1.15.130.83:8080/api/v1/offered_course/' + courseID + '/student').then(function (resp) {
          console.log("测试获取班级信息resp.data.msg：" + resp.data.msg)  // 测试
          console.log("测试班级信息resp.data.data[1].name：" + resp.data.data[1].name)  // 测试
          // that.tmpList = resp.data.data
          // that.tableData = that.tmpList
          that.tableData = resp.data.data
        })
      }
    }

  }
}
</script>