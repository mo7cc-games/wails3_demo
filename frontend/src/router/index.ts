import { createRouter } from 'vue-router';
import { createWebHashHistory } from 'vue-router';
// import { createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('@src/views/HomeView.vue'),
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@src/views/AboutView.vue'),
  },
  {
    path: '/test',
    name: 'test',
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('@src/views/TestView.vue'),
  },
];

const router = createRouter({
  // history: createWebHistory(import.meta.env.BASE_URL),
  history: createWebHashHistory(), //必须使用 hash 模式 否则生产环境 刷新会 404
  routes,
});

export default router;
