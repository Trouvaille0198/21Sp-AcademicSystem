<template>
  <div>
    <el-container>
      <el-main>
        <h1>我开设的所有课程</h1>
        <el-card>
          <el-table
              :data="tableData"
              border
              stripe
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
                prop="credit"
                label="学分"
                width="150">
            </el-table-column>
            <el-table-column
              prop="department"
              label="学院"
              width="150">
            </el-table-column>
            <el-table-column
              prop="term"
              label="学期"
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
    this.tid = sessionStorage.getItem("id");
    this.term = sessionStorage.getItem("currentTerm");
    console.log("测试teacher id ：" + this.tid)  // 测试
    console.log("测试term：" + this.term)  // 测试

    const that = this
    axios.get('http://1.15.130.83:8080/api/v1/teacher/' + this.tid + '/offered_course').then(function (resp) {

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