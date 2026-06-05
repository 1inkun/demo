import { fileURLToPath, URL } from 'node:url';

import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';

// https://vite.dev/config/
export default defineConfig({
	plugins: [vue(), vueDevTools()],
	resolve: {
		alias: {
			'@': fileURLToPath(new URL('./src', import.meta.url)),
		},
	},
	server: {
		proxy: {
			// 将前端 /api 开头的请求代理到后端 8080 端口
			'/api': {
				target: 'http://localhost:8080', // 后端服务地址
				changeOrigin: true, // 支持跨域（修改请求头中的 origin）
				rewrite: (path) => path.replace(/^\/api/, ''), // 如果需要去掉 /api 前缀，取消注释
			},
		},
	},
});
