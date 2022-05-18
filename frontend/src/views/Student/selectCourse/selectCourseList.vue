<template>
  <div>
    <el-form :inline="true">
      <el-form-item label="课程搜索">
        <el-input v-model.trim="query.number" placeholder="输入课程号"></el-input>
        <el-input v-model.trim="query.name" placeholder="输入课程名"></el-input>
        <el-input v-model.trim="query.credit" placeholder="输入学分"></el-input>
      </el-form-item>
    </el-form>
    <template>
      <el-table :data="filteredTableData" border show-header stripe style="width: 100%">
        <el-table-column fixed prop="number" label="课程号" width="150">
        </el-table-column>
        <el-table-column prop="name" label="课程名" width="150">
        </el-table-column>
        <el-table-column prop="teacher_name" label="教师名" width="150">
        </el-table-column>
        <el-table-column prop="credit" label="学分" width="150">
        </el-table-column>
        <el-table-column prop="department" label="学院" width="150">
        </el-table-column>
        <el-table-column prop="term" label="学期" width="150"
          :filters="[{ text: '21-秋季学期', value: '21-秋季学期' }, { text: '21-冬季学期', value: '21-冬季学期' }, { text: '22-春季学期', value: '22-春季学期' }]"
          :filter-method="filterHandler">
        </el-table-column>
        <el-table-column label="操作" width="100">

          <template slot-scope="scope">
            <el-popconfirm confirm-button-text='选择' cancel-button-text='取消' icon="el-icon-info" title="确定选择该教师开设的课程？"
              @confirm="select(scope.row)">
              <el-button slot="reference" type="text" size="small">选课</el-button>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background layout="prev, pager, next" :total="total" :page-size="pageSize"
        @current-change="changePage">
      </el-pagination>
    </template>
  </div>
</template>

<script>
export default {
  methods: {
    // axios是异步操作，改为同步  标识为1，配合2使用
    async select(row) { 
      console.log(row)
      const sid = sessionStorage.getItem('id')
      const courseNumber = row.number
      const courseTeacher = row.teacher_name
      const courseTerm = row.term
      let courseID
      console.log("选课测试 courseNumber ：" + courseNumber)  // 测试
      console.log("选课测试 courseTeacher ：" + courseTeacher)  // 测试
      console.log("选课测试 courseTerm ：" + courseTerm)  // 测试

      // axios是异步操作，改为同步  标识为2，配合1使用
      await axios.get('http://1.15.130.83:8080/api/v1/offered_course').then(function (resp) {
        console.log("测试resp.data.data.length：" + resp.data.data.length)  // 测试
        for (let i = 0; i < resp.data.data.length; i++) {
          if (resp.data.data[i].number === courseNumber &&
             resp.data.data[i].teacher_name === courseTeacher &&
             resp.data.data[i].term === courseTerm) 
          {
            courseID = resp.data.data[i].id
            break
          }
        }
      })
      
      console.log("选课测试 courseID ：" + courseID)  // 测试
      const that = this

      var config = {
        method: 'post',
        url: 'http://1.15.130.83:8080/api/v1/selection',
        data : {
          offeredCourseID : parseInt(courseID),
          studentID : parseInt(sid)
        },
        headers: { 'content-type': 'application/json',}
      };

      axios(config).then(function (resp) {
        console.log("courseID api：" + courseID)  // 测试
        console.log("选课的resp.data.msg：" + resp.data.msg)  // 测试
        if (resp.data.code === 0) {
          that.$message({
            showClose: true,
            message: '选课成功',
            type: 'success'
          });
          // window.location.reload()
        }
        else {
          that.$message({
            showClose: true,
            message: resp.data.msg,
            type: 'error'
          });
        }
      })


    },

    changePage(page) {
      page = page - 1
      const that = this
      let start = page * that.pageSize, end = that.pageSize * (page + 1)
      let length = that.tmpList.length
      let ans = (end < length) ? end : length
      that.tableData = that.tmpList.slice(start, ans)
    },

    filterHandler(value, row, column) {
      const property = column['property'];
      return row[property] === value;
    },

  },

  data() {
    return {
      tableData: [
        {
          "id": 1,
          "name": "数据库原理(1)",
          "number": "0121",
          "credit": 4,
          "teacher_name": "老师A",
          "department": "计算机学院",
          "term": "21-冬季学期"
        }
      ],
      pageSize: 10,
      total: null,
      tmpList: null,
      type: sessionStorage.getItem('type'),
      query: {
        number: '',
        name: '',
        credit: '',
        teacher_name:''
      },
      
    }
  },

  props: {
    ruleForm: Object
  },

  computed: {
    filteredTableData() {
      return this.tableData.filter(item => {
        if (
          (this.query.number === '' || item.number.includes(this.query.number)) &&
          (this.query.name === '' || item.name.includes(this.query.name)) &&
          (this.query.credit === '' || item.credit == parseInt(this.query.credit))
        )
          return item
      })
    },

    // filteredSelectCourse() {
    //   return this.tableData.filter(item => {
    //     if (
    //       (this.query.number === item.number) &&
    //       (this.query.name === item.name) &&
    //       (parseInt(this.query.credit) === item.credit) &&
    //       (this.query.teacher_name === item.teacher_name)
    //     )
    //       return item
    //   })
    // }

  },

  watch: {},

  mounted: function () {
    axios.get('http://1.15.130.83:8080/api/v1/offered_course').then((resp) => {
      console.log(resp.data)
      this.tableData = resp.data.data
      // console.log(this.tableData)
    })
  }
}
</script>