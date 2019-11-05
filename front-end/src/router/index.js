import Vue from 'vue'
import Router from 'vue-router'
import Ping from '@/components/Ping'
import Home from '@/components/Home'
import Login from '@/components/Login'
import Register from '@/components/Register'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/ping',
      name: "Ping",
      component: Ping
    },
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/login',
      name: 'Login',
      component: Login

    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    }
  ]
})
