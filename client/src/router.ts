import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/Home.vue';
import About from '@/views/About.vue';
import Login from '@/views/Login.vue';
import { Dep } from '@/dep';


Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/',
      name: 'home',
      component: Home,
      beforeEnter: (to, from, next) =>  Dep.getGuardSvc().beforeEnter(to,from,next)
    },
    {
      path: '/about',
      name: 'about',
      component: About,
      beforeEnter: (to, from, next) => Dep.getGuardSvc().beforeEnter(to,from,next)
    },
  ],
});
