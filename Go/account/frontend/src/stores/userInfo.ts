import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useUserInfo = defineStore('userInfoStore', () => {
	const count = ref(0);
	const isLogin = ref(false);
	const logIn = function () {
		isLogin.value = true;
	};
	const logOut = function () {
		isLogin.value = false;
	};

	return { count, logIn, logOut, isLogin };
});
