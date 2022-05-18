<template>
  <div>
    <el-container>
      <el-main>
        <h1>学院展示</h1>
        <el-card>
          <el-table
              :data="tableData"
              border
              stripe
              style="width: 100%">
            <el-table-column
                prop="number"
                label="学院号"
                width="150">
            </el-table-column>

            <el-table-column
                prop="name"
                label="学院名"
                width="150">
            </el-table-column>

            <el-table-column
                prop="course_num"
                label="课程数"
                width="150">
            </el-table-column>

            <el-table-column
                prop="student_num"
                label="学生数"
                width="150">
            </el-table-column>

            <el-table-column
                prop="teacher_num"
                label="教师数"
                width="150">
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
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  methods: {
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
      tmpList: null,
      tid: null,
      term: null
    }
  },
  created() {
    const that = this
    axios.get('http://1.15.130.83:8080/api/v1/department').then(function (resp) {

      that.tmpList = resp.data.data
      that.total = resp.data.data.length
      let start = 0, end = that.pageSize
      let length = that.tmpList.length
      let ans = (end < length) ? end : length
      that.tableData = that.tmpList.slice(start, ans)

    })

  }

}
</script>