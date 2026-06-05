import Index from '@/components/Index/index.vue';
import Login from '@/components/Login/login.vue';
import reg from '@/components/Register/reg.vue';
import { createRouter, createWebHistory } from 'vue-router';
import { useUserInfo } from '@/stores/userInfo';

const routes = [
	{ path: '/', name: 'Index', component: Index },
	{ path: '/login', name: 'Login', component: Login },
	{ path: '/reg', name: 'Register', component: reg },
];

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: routes,
});

router.beforeEach(async (to, from) => {
	const userInfoStore = useUserInfo();
	if (
		// 检查用户是否已登录
		userInfoStore.isLogin &&
		// ❗️ 避免无限重定向
		to.name !== 'Login'
	) {
		// 将用户重定向到登录页面
		return { name: 'Login' };
	}
});

export default router;
