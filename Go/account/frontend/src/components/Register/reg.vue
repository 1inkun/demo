<!-- eslint-disable vue/multi-word-component-names -->
<template>
	<div class="container">
		<div class="login-container">
			<div class="login-header">
				<h2>注册新账户</h2>
			</div>
			<div class="login-body">
				<div class="login-item">
					<div>账号:</div>
					<n-input
						v-model:value="newUserInfo.username"
						type="text"
						placeholder="输入账号"
					/>
				</div>
				<div class="login-item">
					<div>昵称:</div>
					<n-input
						v-model:value="newUserInfo.nickname"
						type="text"
						placeholder="输入昵称"
					/>
				</div>
				<div class="login-item">
					<div>密码:</div>
					<n-input
						v-model:value="newUserInfo.password"
						type="password"
						placeholder="输入密码"
					/>
				</div>
				<div class="login-item">
					<div>确认密码:</div>
					<n-input
						v-model:value="checkPasswd"
						type="password"
						placeholder="确认密码"
					/>
				</div>
			</div>
			<div class="login-footer">
				<n-button
					block
					type="primary"
					:disabled="!isSamePasswd"
					@click="submit"
					>注册</n-button
				>
			</div>
		</div>
	</div>
</template>
<script setup lang="ts">
import { NInput, NButton } from 'naive-ui';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
type NEWUSERINFO = {
	nickname: string;
	username: string;
	password: string;
};
const newUserInfo = ref<NEWUSERINFO>({
	nickname: '',
	username: '',
	password: '',
});
const router = useRouter();
const checkPasswd = ref('');
const isSamePasswd = computed(() => {
	return checkPasswd.value === newUserInfo.value.password ? true : false;
});
const submit = async function () {
	console.log(newUserInfo.value);
	const url: string = '/api/reg';
	if (
		newUserInfo.value.username === '' ||
		newUserInfo.value.password === ''
	) {
		alert('请检查输入的信息');
		newUserInfo.value = {
			nickname: '',
			username: '',
			password: '',
		};
	}
	try {
		const rep = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(newUserInfo.value),
		});
		console.log(rep.status);
		if (200 <= rep.status && rep.status < 300) {
			newUserInfo.value = {
				nickname: '',
				username: '',
				password: '',
			};
			router.push({ name: 'Login' });
			alert('注册成功!');
		}
	} catch (error) {
		console.error(error);
	}
};
</script>
<style lang="css">
@import url('@/static/css/reg.css');
</style>
