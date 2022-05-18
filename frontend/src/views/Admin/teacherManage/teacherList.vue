<template>
  <div>
    <el-form :inline="true">
      <el-form-item label="教师搜索">
        <el-input v-model.trim="query.number" placeholder="输入教师号"></el-input>
        <el-input v-model.trim="query.name" placeholder="输入姓名"></el-input>
      </el-form-item>
    </el-form>
    <template>
      <el-table :data="filteredTableData" border show-header stripe style="width: 100%">
        <el-table-column 
            fixed prop="number" label="教师号" width="150">
        </el-table-column> 

        <el-table-column 
            prop="name" label="姓名" width="150">
        </el-table-column>

        <el-table-column 
              prop="sex" label="性别" width="150">
        </el-table-column>

        <el-table-column 
            prop="age" label="年龄" width="150">
        </el-table-column>

        <el-table-column 
            prop="department_name" label="学院" width="150">
        </el-table-column>

        <el-table-column 
            label="操作" width="100">
            <template slot-scope="scope">
              <el-popconfirm confirm-button-text='确认' cancel-button-text='取消' icon="el-icon-info" title="确定删除该学生？"
                @confirm="deleteStudent(scope.row)">
                <el-button slot="reference" type="text" size="small">删除教师</el-button>
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
    // 第1步 ： 找出要删除的教师的ID
        async deleteStudent(row) { 
          console.log(row)
          const studentNumber = row.number
          let studentID
          console.log("教师的studentNumber ：" + studentNumber)  // 测试

          await axios.get('http://1.15.130.83:8080/api/v1/teacher').then(function (resp) {
            console.log("找出教师IDresp.data.msg：" + resp.data.msg)  // 测试
            for (let i = 0; i < resp.data.data.length; i++) {
              if (resp.data.data[i].number === studentNumber) 
              {
                studentID = resp.data.data[i].id
                break
              }
            }
          })
          
          console.log("删除教师 studentID ：" + studentID)  // 测试
          const that = this
    // 第2步 ： 通过ID删除教师
      
        await axios.delete('http://1.15.130.83:8080/api/v1/teacher/' + studentID).then(function (resp) {
          console.log("删除教师的resp.data.msg：" + resp.data.msg)  // 测试
          if (resp.data.code === 0) {
            that.$message({
              showClose: true,
              message: '删除教师成功',
              type: 'success'
            });
            window.location.reload()
          }
          else {
            that.$message({
              showClose: true,
              message: resp.data.msg,
              type: 'error，请检查数据库'
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
      tableData: [
        {
          "id": 1,
          "number": "0121",
          "name": "朝承恩",
          "sex": "男",
          "age": 20,
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
          (this.query.name === '' || item.name.includes(this.query.name)) 
        )
          return item
      })
    },

  },

  watch: {},

  mounted: function () {
    axios.get('http://1.15.130.83:8080/api/v1/teacher').then((resp) => {
      console.log(resp.data)
      this.tableData = resp.data.data
      // console.log(this.tableData)
    })
  }
}
</script>