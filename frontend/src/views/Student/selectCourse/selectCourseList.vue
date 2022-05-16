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
              <el-button slot="reference" type="text" size="small">选择</el-button>
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
    select(row) {
      console.log(row)
      const cid = 1
      const tid = 1
      const sid = sessionStorage.getItem('id')
      const term = sessionStorage.getItem('currentTerm')
      const sct = {
        cid: cid,
        tid: tid,
        sid: sid,
        term: term
      }
      const that = this
      axios.post('http://localhost:10086/SCT/save', sct).then(function (resp) {
        if (resp.data === '选课成功') {
          that.$message({
            showClose: true,
            message: '选课成功',
            type: 'success'
          });
        }
        else {
          that.$message({
            showClose: true,
            message: resp.data,
            type: 'error'
          });
        }
      })

    },
    deleteCourseTeacher(row) {
      const that = this
      axios.post('http://localhost:10086/courseTeacher/deleteById', row).then(function (resp) {
        if (resp.data === true) {
          that.$message({
            showClose: true,
            message: '删除成功',
            type: 'success'
          });
          window.location.reload()
        }
        else {
          that.$message({
            showClose: true,
            message: '删除出错，请查询数据库连接',
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
        credit: ''
      }
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
    }
  },
  watch: {},

  mounted: function () {
    axios.get('http://1.15.130.83:8080/api/v1/offered_course').then((resp) => {
      console.log(resp.data)
      this.tableData = resp.data.data
      console.log("????")
      console.log(this.tableData)
    })
  }
}
</script>