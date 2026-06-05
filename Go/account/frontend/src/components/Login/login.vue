<!-- eslint-disable vue/multi-word-component-names -->
<template>
	<div class="container">
		<div class="login-container">
			<div class="login-header">
				<h2>欢迎您来</h2>
			</div>
			<div class="login-body">
				<div class="login-item">
					<span>账号:</span>
					<n-input
						v-model:value="userInfo.username"
						type="text"
						placeholder="输入账号"
					/>
				</div>
				<div class="login-item">
					<span>密码:</span>
					<n-input
						v-model:value="userInfo.password"
						type="password"
						placeholder="输入密码"
					/>
				</div>
				<div class="register" @click="goToregister">
					<span>没有账户?去注册!</span>
				</div>
			</div>
			<div class="login-footer">
				<n-button block type="primary" @click="login">登录</n-button>
			</div>
		</div>
	</div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { NInput, NButton } from 'naive-ui';
import { useRouter } from 'vue-router';
const router = useRouter();
type USERINFO = {
	username: string;
	password: string;
};

const userInfo = ref<USERINFO>({
	username: '',
	password: '',
});
const goToregister = function () {
	router.push({ name: 'Register' });
};
const login = async function () {
	console.log(userInfo.value);
	const url: string = '/api/login';
	try {
		const rep = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(userInfo.value),
		});
		console.log(rep.status);
	} catch (error) {
		console.error(error);
	}
};
</script>
<style lang="css">
@import url('@/static/css/login.css');
</style>
