import Vue from 'vue';
import VueRouter from 'vue-router';
import login from '../views/login/index';
import admin from '../views/Admin/index';
import adminHome from '../views/Admin/home';
import studentManage from '../views/Admin/studentManage/index'
import addStudent from "@/views/Admin/studentManage/addStudent";
import studentList from "@/views/Admin/studentManage/studentList";
import teacherManage from "@/views/Admin/teacherManage/index"
import addTeacher from "@/views/Admin/teacherManage/addTeacher";
import courseManage from "@/views/Admin/courseManage/index";
import addCourse from "@/views/Admin/courseManage/addCourse";
import teacher from "@/views/Teacher/index";
import teacherList from "@/views/Admin/teacherManage/teacherList";
import student from "@/views/Student/index";
import courseList from "@/views/Admin/courseManage/courseList";
import offerCourseClassInfo from "@/views/Teacher/offerCourseClassInfo";
import teacherHome from "@/views/Teacher/home";
import setCourse from "@/views/Teacher/setCourse";
import studentHome from "@/views/Student/home";
import myOfferCourse from "@/views/Teacher/myOfferCourse";

import CourseTeacherManage from "@/views/Admin/selectCourseManage/index";
import offeredCourseManage from "@/views/Admin/selectCourseManage/offeredCourseManage";
import departmentManage from "@/views/Admin/departmentManage/index";
import departmentList from "@/views/Admin/departmentManage/departmentList";

import studentSelectCourseManage from "@/views/Student/selectCourse/index";
import selectCourse from "@/views/Student/selectCourse/selectCourse";
import querySelectedCourse from "@/views/Student/selectCourse/querySelectedCourse";
import studentCourseGrade from "@/views/Student/courseGrade/index";
import queryCourseGrade from "@/views/Student/courseGrade/queryCourseGrade";
import teacherGradeCourseManage from "@/views/Teacher/teacherGradeCourseManage/index";
import teacherEditorGradeCourse from "@/views/Teacher/teacherGradeCourseManage/teacherEditorGradeCourse";
import updateInfo from "@/components/updateInfo";

Vue.use(VueRouter)

const routes = [
  {
    // ??????
    path: '/',
    name: 'index',
    component: login,
    redirect: '/login'
  },
  {
    // ?????????
    path: '/login',
    name: 'login',
    component: login
  },
  {
    // admin ?????????
    path: '/admin',
    name: 'admin',
    redirect: '/adminHome',
    component: admin,
    meta: {requireAuth: true},
    children: [
      {
        path: '/adminHome',
        name: '?????????',
        component: adminHome,
        meta: {requireAuth: true},
        children: [
          {
            path: '/adminHome',
            name: '???????????????',
            component: adminHome,
            meta: {requireAuth: true},
          }
        ]
      },
      {
        path: '/studentManage',
        name: '????????????',
        component: studentManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addStudent',
            name: '????????????',
            component: addStudent,
            meta: {requireAuth: true}
          },
          {
            path: '/studentList',
            name: '????????????',
            component: studentList,
            meta: {requireAuth: true},
          },
          
        ]
      },
      {
        path: '/teacherManage',
        name: '????????????',
        component: teacherManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addTeacher',
            name: '????????????',
            component: addTeacher,
            meta: {requireAuth: true}
          },
          {
            path: '/teacherList',
            name: '????????????',
            component: teacherList,
            meta: {requireAuth: true},
            
          },
        ]
      },
      {
        path: '/courseManage',
        name: '????????????',
        component: courseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addCourse',
            name: '????????????',
            component: addCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/courseList',
            name: '????????????',
            component: courseList,
            meta: {requireAuth: true},
          },
          
        ]
      },
      {
        path: '/CourseTeacher',
        name: '????????????',
        component: CourseTeacherManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/offeredCourseManage',
            name: '????????????',
            component: offeredCourseManage,
            meta: {requireAuth: true},
          }
        ]
      },

      {
        path: '/departmentManage',
        name: '????????????',
        component: departmentManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/departmentList',
            name: '????????????',
            component: departmentList,
            meta: {requireAuth: true},
          }
        ]
      },

      
    ]
  },
  {
    path: '/teacher',
    name: 'teacher',
    component: teacher,
    redirect: '/teacherHome',
    meta: {requireAuth: true},
    children: [
      {
        path: '/teacherHome',
        name: '?????????',
        meta: {requireAuth: true},
        component: teacherHome,
        children: [
          {
            path: '/teacherHome',
            name: '????????????',
            meta: {requireAuth: true},
            component: teacherHome
          },
        ]
      },

      {
        path: '/courseManage',
        name: '????????????',
        meta: {requireAuth: true},
        component: setCourse,
        children: [
          {
            path: '/myOfferCourse',
            name: '??????????????????',
            component: myOfferCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/offerCourseClassInfo',
            name: '??????????????????',
            component: offerCourseClassInfo,
            meta: {requireAuth: true}
          },
        ]
      },
      {
        name: '??????????????????',
        path: '/teacherQueryGradeCourseManage',
        component: teacherGradeCourseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/teacherEditorGradeCourse',
            name: '????????????',
            component: teacherEditorGradeCourse,
            meta: {requireAuth: true}
          }
        ]
      }
    ]
  },
  {
    path: '/student',
    name: 'student',
    component: student,
    redirect: '/studentHome',
    meta: {requireAuth: true},
    children: [
      {
        path: '/student',
        name: '?????????',
        component: studentHome,
        meta: {requireAuth: true},
        children: [
          {
            path: '/studentHome',
            name: '????????????',
            component: studentHome,
            meta: {requireAuth: true},
          },
        ],
      },
      
      {
        path: '/studentSelectCourseManage',
        name: '????????????',
        component: studentSelectCourseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/studentSelectCourse',
            name: '??????',
            component: selectCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/querySelectedCourse',
            name: '????????????',
            component: querySelectedCourse,
            meta: {requireAuth: true}
          }
        ]
      },
      {
        path: '/courseGrade',
        name: '??????????????????',
        component: studentCourseGrade,
        meta: {requireAuth: true},
        children: [
          {
            path: '/queryCourseGrade',
            name: '????????????',
            component: queryCourseGrade,
            meta: {requireAuth: true}
          },
        ]
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

/*
  session ?????????
    1. token
    2. name
    3. type
    4. tid
    5. sid
    5. ???????????? info
 */
router.beforeEach((to, from, next) => {
  console.log(from.path + ' ====> ' + to.path)
  if (to.meta.requireAuth) { // ???????????????????????????????????????
    if (sessionStorage.getItem("token") === 'true') { // ????????????????????????token
      next()
    } else {
      // ?????????,?????????????????????
      next({
        path: '/login',
        query: {redirect: to.fullPath}
      })
    }
  } else {
    // ????????????????????????????????????????????????????????????????????????
    if(sessionStorage.getItem("token") === 'true'){
      console.log('???????????????????????????????????????????????????')
      const t = sessionStorage.getItem("type")
      next('/' + t);
    }else{
      next();
    }
  }
});