<template>
  <div>
    <el-container>
      <el-main>
        <el-card style="margin-top: 10px">
          <select-course-list :rule-form="ruleForm"></select-course-list>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>
<script>
import SelectCourseList from "@/views/Student/selectCourse/selectCourseList";
export default {
  components: { SelectCourseList },
  data() {
    return {
      ruleForm: {
        teacherNumber: null,
        teacherID: null,
        courseNumber: null,
        courseID: null,
      },
    };
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>