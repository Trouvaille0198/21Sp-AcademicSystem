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
        <el-table-column 
            fixed prop="number" label="课程号" width="150">
        </el-table-column>

        <el-table-column 
            prop="name" label="课程名" width="150">
        </el-table-column>

        <el-table-column 
            prop="credit" label="学分" width="150">
        </el-table-column>

        <el-table-column 
            prop="teacher_name" label="任课教师" width="150">
        </el-table-column>

        <el-table-column 
            prop="department" label="学院" width="150">
        </el-table-column>

        <el-table-column 
            prop="term" label="开课学期" width="150">
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
      tableData: [
        {
          "id": 1,
          "number": "0121",
          "name": "朝承恩",
          "credit": 4,
          "department_name": "计算机学院",
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