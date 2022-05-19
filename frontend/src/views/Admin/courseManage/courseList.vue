<template>
  <div>
    <el-form :inline="true">
      <el-form-item label="搜索">
        <el-input v-model.trim="query.number" placeholder="输入课号"></el-input>
        <el-input v-model.trim="query.name" placeholder="输入课名"></el-input>
        <el-input v-model.trim="query.credit" placeholder="输入学分"></el-input>
      </el-form-item>
      <!-- <el-button type="primary" icon="el-icon-search" @click="handlesearch"></el-button> -->
    </el-form>
    <el-table :data="filteredTableData" border stripe style="width: 100%">
      <el-table-column fixed prop="number" label="课程号" width="150">
      </el-table-column>
      <el-table-column prop="name" label="课程名" width="150">
      </el-table-column>
      <el-table-column prop="credit" label="学分" width="150">
      </el-table-column>
      <el-table-column prop="department_name" label="学院" width="150"
          :filters="[{ text: '理学院', value: '理学院' }, { text: '计算机学院', value: '计算机学院' }, 
          { text: '通信学院', value: '通信学院' }, { text: '外国语学院', value: '外国语学院' }]"
          :filter-method="filterHandler">
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template slot-scope="scope">
          <el-popconfirm confirm-button-text='删除' cancel-button-text='取消' icon="el-icon-info" icon-color="red" title="删除不可复原"
              @confirm="deleteCourse(scope.row)">
            <el-button slot="reference" type="text" size="small">删除</el-button>
          </el-popconfirm>
          <el-button @click="editor(scope.row)" type="text" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        background
        layout="prev, pager, next"
        :total="total"
        :page-size="pageSize"
        @current-change="changePage"
    >
    </el-pagination>
  </div>
</template>

<script>
import { debug } from 'console'

export default {
  methods: {
    select(row) {
      console.log(row)
    },
    deleteCourse(row) {
      // let courseNumber = row.number
      let courseID = 0 
      courseID = row.offered_course_id
      let sid = 0 
      sid = sessionStorage.getItem('id')

      // let sct = new FormData();

      // axios.get("http://1.15.130.83:8080/api/v1/student/" + sid + "/selected_course").then(function (resp) {
      //   console.log("测试查询课程的api状态code:"+resp.data.code)  // 测试
        // for (let i = 0; i < resp.data.data.length; i++) {
        //   if (resp.data.data[i].number === courseNumber) {
        //     courseID = resp.data.data[i].offered_course_id
        //     sct.append("offeredCourseID", resp.data.data[i].offered_course_id)
        //     console.log("courseID 1：" + courseID)  // 测试
        //     break
        //   }

        // }
      // })
      // courseID = 6
      // sct.append("offeredCourseID", 6)
      // console.log("courseID 1：" + courseID)  // 测试
      // setTimeout(function(){
      //   console.log("sleep 1 seconds");
      // },1000);


      // console.log("courseID 2：" + courseID)  // 测试
      // sct.append("studentID", sid)
      const that = this
      // console.log("student id：" + sid)  // 测试

      // console.log("courseID 3：" + courseID)  // 测试


      var config = {
        method: 'delete',
        url: 'http://1.15.130.83:8080/api/v1/selection',
        data : {
          offeredCourseID : parseInt(courseID),
          studentID : parseInt(sid)
        },
        headers: { 'content-type': 'application/json',}
      };


      axios(config).then(function (resp) {
        console.log("courseID api：" + courseID)  // 测试
        console.log("student api：" + sid)  // 测试
        console.log("退课的resp.data.msg：" + resp.data.msg)  // 测试
        if (resp.data.code === 0) {
          that.$message({
            showClose: true,
            message: '退课成功',
            type: 'success'
          });
          window.location.reload()
        }
        else {
          that.$message({
            showClose: true,
            message: '退课失败，请联系管理员',
            type: 'error'
          });
        }
      })







      // const that = this
      // var param = {id: row.id}
      // console.log(row);
      // axios.delete('http://1.15.130.83:8080/api/v1/course', {data : {id : row.id}}).then(function (resp) {
      //   console.log(resp)
      //   if (resp.data === true) {
      //     that.$message({
      //       showClose: true,
      //       message: '删除成功',
      //       type: 'success'
      //     });
      //     window.location.reload()
      //   }
      //   else {
      //     that.$message({
      //       showClose: true,
      //       message: '删除出错，请查询数据库连接',
      //       type: 'error'
      //     });
      //   }
      // }).catch(function (error) {
      //   console.log(error)
      //   that.$message({
      //     showClose: true,
      //     message: '删除出错，存在外键依赖',
      //     type: 'error'
      //   });
      // })
    },
    offer(row) {
      const tid = sessionStorage.getItem("tid")
      const cid = row.cid
      const term = sessionStorage.getItem("currentTerm")

      const that = this
      axios.get('http://localhost:10086/courseTeacher/insert/' + cid + '/' + tid + '/' + term).then(function (resp) {
        if (resp.data === true) {
          that.$message({
            showClose: true,
            message: '开设成功',
            type: 'success'
          });
          window.location.reload()
        }
        else {
          that.$message({
            showClose: true,
            message: '开设失败，请联系管理员',
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
    editor(row) {
      this.$router.push({
        path: '/editorCourse',
        query: {
          cid: row.name
        }
      })
    },
    filterHandler(value, row, column) {
      const property = column['property'];
      return row[property] === value;
    },
  },
  // created() {
  //   // console.log(this.type)
  //   const that = this;
  //   that.tmpList = []
  //   axios.get('http://1.15.130.83:8080/api/v1/course').then(function (resp) {
  //     console.log("测试resp.data.data.length：" + resp.data.data.length)  // 测试
  //     for (let i = 0; i < resp.data.data.length; i++) {
  //       //if (resp.data.data[i].term === term) {
  //         console.log("测试resp.data.data[i].name：" + resp.data.data[i].name + "    ,i:" + i);  // 测试
  //         that.tmpList.push(resp.data.data[i]);
  //       //}
          
  //     }
  //     console.log("测试that.tmpList.length：" + that.tmpList.length);  // 测试
      
  //     that.total = that.tmpList.length
  //     let start = 0, end = that.pageSize
  //     let length = that.tmpList.length
  //     let ans = (end < length) ? end : length
  //     that.tableData = that.tmpList.slice(start, end)
  //   })
  // },
  data() {
    return {
      tableData: [
        {
          "id": 1,
          "name": "数据库原理(1)",
          "number": "0121",
          "credit": 4,
          "department_name": "计算机学院",
        }
      ],
      pageSize: 10,
      total: null,
      tmpList: [],
      type: sessionStorage.getItem("type"),
      query: {
        number: '',
        name: '',
        credit: ''
      }
    }
  },
  props: {
    ruleForm: Object,
    isActive: Boolean
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
    }
  },
  watch: {
    // ruleForm: {
    //   handler(newRuleForm, oldRuleForm) {
    //     console.log("组件监听 form")
    //     console.log(newRuleForm)
    //     const that = this
    //     that.tmpList = null
    //     that.total = null
    //     that.tableData = null
    //     // axios.post("http://localhost:10086/course/findBySearch", newRuleForm).then(function (resp) {
    //     // axios.get('http://1.15.130.83:8080/api/v1/course/' + newRuleForm.number).then(function (resp) {
    //     //   console.log("查询结果:");
    //     //   console.log(resp)
    //     //   that.tmpList = resp.data
    //     //   that.total = resp.data.length
    //     //   let start = 0, end = that.pageSize
    //     //   let length = that.tmpList.length
    //     //   let ans = (end < length) ? end : length
    //     //   that.tableData = that.tmpList.slice(start, ans)
    //     // })
    //     axios.get('http://1.15.130.83:8080/api/v1/course/' + newRuleForm.number).then(function (resp) {
    //       for (let i = 0; i < resp.data.data.length; i++) {
    //         //if (resp.data.data[i].number === newRuleForm.number) {
    //           console.log("测试resp.data.data[i].name：" + resp.data.data[i].name + "    ,i:" + i);  // 测试
    //           that.tmpList.push(resp.data.data[i]);
    //         //}
            
    //       }
    //       console.log("测试that.tmpList.length：" + that.tmpList.length);  // 测试
          
    //       that.total = that.tmpList.length
    //       let start = 0, end = that.pageSize
    //       let length = that.tmpList.length
    //       let ans = (end < length) ? end : length
    //       that.tableData = that.tmpList.slice(start, end)
    //     })
    //   },
    //   deep: true,
    //   immediate: true
    // }
  },
  mounted: function () {
    axios.get('http://1.15.130.83:8080/api/v1/course').then((resp) => {
      console.log(resp.data)
      this.tableData = resp.data.data
      console.log("????")
      console.log(this.tableData)
    })
  }
}
</script>