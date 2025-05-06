import { createWebHistory, createRouter, RouteRecordRaw } from 'vue-router'

import HomeView from '@/views/Home.vue'
import NotFoundView from '@/views/NotFound.vue'
import SignInView from '@/views/SignIn.vue'
import SignUpView from '@/views/SignUp.vue'
import MyProfileView from '@/views/MyProfile.vue'

const routes = [
  { path: '/', name: 'home', component: HomeView, meta: { title: 'Home - eduflow' } },
  { path: '/auth', name: 'auth', children : [
    { path: 'sign-in', name: 'sign-in', component: SignInView, meta: { title: 'Signing in - eduflow' } }, 
    { path: 'sign-up', name: 'sign-up', component: SignUpView, meta: { title: 'Signing up - eduflow' } }, 
  ]},
  { path: '/profile', name: 'profile', component: MyProfileView, meta: { title: 'Profile - eduflow' } },
  { path: '/:notFound(.*)', name: 'not-found', component: NotFoundView, meta: { title: '404 not found - eduflow' } }
]

const router = createRouter({
  history: createWebHistory(),
  routes: <RouteRecordRaw[]>routes,
})

router.beforeEach((to, from) => {
  document.title = <string>to.meta?.title ?? 'eduflow'
})


export default router
