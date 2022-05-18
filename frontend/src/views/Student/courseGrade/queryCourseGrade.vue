<template>
  <div>
    <el-form >
      <el-form-item label="选择学期">
        <el-select v-model="term" @change="currTermChange" placeholder="请选择学期">
          <el-option v-for="(item, index) in termList" :key="index" :label="item" :value="item"></el-option>
        </el-select>
      </el-form-item>
    </el-form>
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
            prop="department"
            label="学院"
            width="150">
        </el-table-column>
        <el-table-column
            prop="teacher_name"
            label="教师名"
            width="150">
        </el-table-column>
        <el-table-column
            prop="credit"
            label="学分"
            width="150">
        </el-table-column>
        <el-table-column
            prop="score"
            label="绩点"
            width="150">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ caculateGPA(scope.row.score) }}</span>
            </template>
        </el-table-column>
        <el-table-column
            prop="score"
            label="成绩"
            width="150">
        </el-table-column>
      </el-table>
      <p>
        平均成绩：{{ avg }}
      </p>
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
    changePage(page) {
      page = page - 1
      const that = this
      let start = page * that.pageSize, end = that.pageSize * (page + 1)
      let length = that.tmpList.length
      let ans = (end < length) ? end : length
      that.tableData = that.tmpList.slice(start, ans)
    },

    currTermChange(changeTerm) {
      console.log("method中更改了selectedTerm：" + changeTerm)  // 测试
      sessionStorage.setItem("selectedTerm", changeTerm)
      
    },
    caculateGPA(score)
    {
      let ans = 0
      if(score >= 90) ans = 4
      else if(score >= 85) ans = 3.7
      else if(score >= 82) ans = 3.3
      else if(score >= 78) ans = 3.0
      else if(score >= 75) ans = 2.7
      else if(score >= 72) ans = 2.3
      else if(score >= 68) ans = 2.0
      else if(score >= 66) ans = 1.7
      else if(score >= 64) ans = 1.5
      else if(score >= 60) ans = 1.0
      else ans = 0
      return ans.toFixed(1)
    },
  },
  data() {
    return {
      tableData: null,
      pageSize: 10,
      total: null,
      tmpList: [],
      avg: 0,
      term: sessionStorage.getItem('selectedTerm'),
      termList: null
    }
  },
  created() {
    const that = this
    that.termList = ["21-秋季学期", "21-冬季学期", "21-春季学期", "22-秋季学期", "22-冬季学期", "22-春季学期"] 
  },

  watch: {
    term: {
      handler(newTerm, oldTerm) {
        const snumber = sessionStorage.getItem('number')
        const sid = sessionStorage.getItem('id')
        console.log("测试snumber：" + snumber)  // 测试
        console.log("测试sid：" + sid)  // 测试
        const that = this

        // const term = sessionStorage.getItem('currentTerm')
        let selectedTerm = sessionStorage.getItem('selectedTerm')
        console.log("测试selectedTerm：" + selectedTerm)  // 测试
      
        // console.log("测试term：" + term)  // 测试
        
        this.tmpList = []
        axios.get('http://1.15.130.83:8080/api/v1/student/' + sid + '/selected_course?hasScore=true').then(function (resp) {
          
          for (let i = 0; i < resp.data.data.length; i++) {
            if (resp.data.data[i].term === selectedTerm) {
              console.log("测试resp.data.data[i].term：" + resp.data.data[i].term + "    ,i:" + i)  // 测试
              resp.data.data[i]
              that.tmpList.push(resp.data.data[i])
              // console.log("第二次测试selectedTerm：" + selectedTerm)  // 测试
            }
          }


          // that.tmpList = resp.data.data
          // console.log("测试code：" + resp.data.code)  // 测试
          that.total = that.tmpList.length
          console.log("测试total：" + that.total)  // 测试
          let start = 0, end = that.pageSize
          let length = that.tmpList.length
          console.log("测试课程的数量length：" + length)  // 测试
          let ans = (end < length) ? end : length
          that.tableData = that.tmpList.slice(start, end)
          let totalScore = 0
          that.avg = 0
          for (let i = 0; i < that.total; i++) {
            totalScore += that.tmpList[i].credit
            that.avg += that.tmpList[i].credit * that.tmpList[i].score
          }
          console.log("测试totalScore：" + totalScore)  // 测试
          console.log("测试avg：" + that.avg)  // 测试
          if (totalScore === 0)
            that.avg = 0
          else
            that.avg /= totalScore
        })
      },
      immediate: true
    }
  }
}
</script>

<!--
  TODO：
  1. 管理员：
    1. 学生选课管理
    2. 成绩管理（只能当前学期）
  2. 学生：成绩排名
  3. 教师：成绩管理（？使用弹框）（只能当前学期）

-->