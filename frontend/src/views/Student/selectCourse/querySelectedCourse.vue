<!-- 学生/选课管理/查询课表 -->

<template>
  <div>
    <el-card>
      <el-table
          :data="tableData"
          border
          style="width: 100%">
        <el-table-column
            fixed
            prop="number"
            label="课程号"
            width="150">
        </el-table-column>
        <el-table-column
            prop="name"
            label="课程名"
            width="150">
        </el-table-column>
        <el-table-column
            prop="teacher_name"
            label="教师名"
            width="150">
        </el-table-column>
        <el-table-column
            prop="offered_course_id"
            label="课程ID"
            width="150">
        </el-table-column>
        <el-table-column
            prop="credit"
            label="学分"
            width="150">
        </el-table-column>
        <el-table-column
            label="操作"
            width="100">

          <template slot-scope="scope">
            <el-popconfirm
                confirm-button-text='退课'
                cancel-button-text='取消'
                icon="el-icon-info"
                title="确定退课？"
                @confirm="deleteSCT(scope.row)"
            >
              <el-button slot="reference" type="text" size="small">退课</el-button>
            </el-popconfirm>
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
    </el-card>
  </div>
</template>

<script>
export default {
  methods: {
    deleteSCT(row) {
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

    },
    changePage(page) {
      page = page - 1
      const that = this
      let start = page * that.pageSize, end = that.pageSize * (page + 1)
      let length = that.tmpList.length
      let ans = (end < length) ? end : length
      that.tableData = that.tmpList.slice(start, ans)
    },
  },
  data() {
    return {
      tableData: null,
      pageSize: 10,
      total: null,
      tmpList: [],
      type: sessionStorage.getItem('type')
    }
  },
  created() {
    const sid = sessionStorage.getItem('id')
    console.log("测试sid：" + sid)  // 测试
    const term = sessionStorage.getItem('currentTerm')
    console.log("测试term：" + term)  // 测试
    const that = this
    axios.get('http://1.15.130.83:8080/api/v1/student/' + sid + '/selected_course').then(function (resp) {
      console.log("测试resp.data.data.length：" + resp.data.data.length)  // 测试
      for (let i = 0; i < resp.data.data.length; i++) {
        if (resp.data.data[i].term === term) {
          console.log("测试resp.data.data[i].term：" + resp.data.data[i].term + "    ,i:" + i)  // 测试
          that.tmpList.push(resp.data.data[i])
        }
          
      }
      console.log("测试that.tmpList.length：" + that.tmpList.length)  // 测试
      
      that.total = that.tmpList.length
      let start = 0, end = that.pageSize
      let length = that.tmpList.length
      let ans = (end < length) ? end : length
      that.tableData = that.tmpList.slice(start, end)
    })
  },
}
</script>