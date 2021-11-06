import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Login from '@/views/Login.vue'
import UserInfo from '@/views/UserInfo.vue'
import UserInfoAll from '@/views/UserInfoAll.vue'
import Register from '@/views/Register.vue'
import Verify from '@/views/Verify.vue'
import UserInfoUpdate from '@/views/UserInfoUpdate.vue'
import UserInfoUpdateA from '@/views/UserInfoUpdateA.vue'
import UserDelete from '@/views/UserDelete.vue'
import WXLogin from '@/views/WXLogin.vue'
import WXBind from '@/views/WXBind.vue'
import UserInfoWX from '@/views/UserInfoWX.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/userinfo',
    name: 'UserInfo',
    component: UserInfo
  },
  {
    path: '/userinfoall',
    name: 'UserInfoAll',
    component: UserInfoAll
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/verify',
    name: 'Verify',
    component: Verify
  },
  {
    path: '/userInfoUpdate',
    name: 'UserInfoUpdate',
    component: UserInfoUpdate
  },
  {
    path: '/userInfoUpdateA',
    name: 'UserInfoUpdateA',
    component: UserInfoUpdateA
  },
  {
    path: '/userDelete',
    name: 'UserDelete',
    component: UserDelete
  },
  {
    path: '/wxlogin',
    name: 'WXLogin',
    component: WXLogin
  }, {
    path: '/wxbind',
    name: 'WXBind',
    component: WXBind
  },
  {
    path: '/userInfoWX',
    name: 'UserInfoWX',
    component: UserInfoWX
  }
]

const router = new VueRouter({
  routes
})

const originalPush = router.push
// 修改原型对象中的push方法
router.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

export default router
