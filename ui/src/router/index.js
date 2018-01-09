import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Redirect from '@/components/Redirect'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/:id',
      name: 'Redirect',
      component: Redirect
    }
  ]
})
