import Vue from 'vue'
import Router from 'vue-router'

import store from '@/store'
import storage from '@/model/storage'

import {Message} from 'element-ui'

Vue.use(Router)


// 解决ElementUI导航栏中的vue-router在3.0版本以上重复点菜单报错问题
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

//前台页面路由
const main = {
  path: '/main',
  name: 'MainPage',
  component: () => import('@/views/MainPage'),
  redirect:'/main/home',
  children:[
    {
      path: 'home',
      name: 'HomePage',
      component: () => import('@/views/page/HomePage')
    },
    {
      path: 'current',
      name: 'CurrentPage',
      component: () => import('@/views/page/CurrentPage')
    },
    {
      path: 'archives',
      name: 'ArchivesPage',
      component: () => import('@/views/page/ArchivesPage')
    },
    {
      path: 'current',
      name: 'CurrentPage',
      component: () => import('@/views/page/CurrentPage')
    },
    {
      path: 'journal',
      name: 'JournalPage',
      component: () => import('@/views/page/JournalPage')
    },
    {
      path: 'article',
      name: 'ArticlePage',
      component: () => import('@/views/page/ArticlePage')
    },
    {
      path: 'announcements',
      name: 'AnnouncementPage',
      component:  () => import('@/views/page/AnnouncementPage')
    },
    {
      path: 'search',
      name: 'SearchPage',
      component: () => import('@/views/page/SearchPage')
    },
    {
      path: 'about',
      name: 'AboutPage',
      component: () => import('@/views/page/AboutPage')
    },
    {
      path: 'submission',
      name: 'SubmissionPage',
      component: () => import('@/views/page/SubmissionPage')
    },
    {
      path: 'privacy',
      name: 'PrivacyPage',
      component: () => import('@/views/page/PrivacyPage')
    },
    {
      path: 'editorial',
      name: 'EditorialPage',
      component: () => import('@/views/page/EditorialPage')
    },
    {
      path: 'contact',
      name: 'ContactPage',
      component: () => import('@/views/page/ContactPage')
    },
    {
      path: 'login',
      name: 'LoginPage',
      component: () => import('@/views/page/LoginPage')
    },
    {
      path: 'register',
      name: 'RegisterPage',
      component: () => import('@/views/page/RegisterPage')
    }
  ]
}

//后台页面路由
const admin = {
  path: '/admin',
  name: 'AdminPage',
  component: () => import('@/views/Admin/AdminPage'),
  redirect:'/admin/submission',
  children:[
    {
      path: 'submission',
      name: 'NewSubmission',
      component: () => import('@/views/Admin/Submisson/NewSubmission')
    },
    {
      path:'historySubmission',
      name:'ProfileSubmission',
      component: () => import('@/views/Admin/Submisson/ProfileSubmission')
    },
    {
      path:'profileSubmissionDetail',
      name:'ProfileSubmissionDetail',
      component: () => import('@/views/Admin/Submisson/ProfileSubmissionDetail')
    },
    {
      path:'historyReview',
      name:'ProfileReview',
      component: () => import('@/views/Admin/Submisson/ProfileReview')
    },
    {
      path: 'issue',
      name: 'AdminIssue',
      component: () => import('@/views/Admin/management/AdminIssue')
    },
    {
      path: 'issuedetail',
      name: 'IssueDetail',
      component: () => import('@/views/Admin/management/IssueDetail')
    },
    {
      path: 'submissonboard',
      name: 'SubmissonBoard',
      component: () => import('@/views/Admin/management/SubmissonBoard')
    },
    {
      path: 'submissondetail',
      name: 'SubmissonDetail',
      props:true,
      component: () => import('@/views/Admin/management/SubmissonDetail')
    },
    {
      path: 'allusers',
      name: 'UserBoard',
      component: () => import('@/views/Admin/user/UserBoard')
    },
    {
      path: 'profile',
      name: 'ProfilePage',
      component: () => import('@/views/Admin/ProfilePage')
    }
  ]
}

const routes = [
  {
    path: '/',
    redirect:'/main'
  },
  main,
  admin
]

const router = new Router({
  mode: 'history',
  routes
})

// 设置路由守卫，在进页面之前，判断有token，才进入页面，否则返回登录页面
if (storage.get("token")) {
  store.commit("setToken", storage.get("token"));
}
// 导航守卫
// 使用 router.beforeEach 注册一个全局前置守卫，判断用户是否登陆
router.beforeEach((to, from, next) => {
  if (to.path === '/main/login' || to.path === '/main/register') {
    next();
  } else {
    let token =store.getters.getToken;
    if (token === undefined || token === null || token === '') {
      Message({
        message: '请先登录',
        type: 'warning'
      });
      next('/main/login');
    } else {
      next();
    }
  }
});

export default router;