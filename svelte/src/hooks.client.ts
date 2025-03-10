// hooks.client.ts
import { setProfile, setUser } from '$lib/stores';
import type { ClientInit } from '@sveltejs/kit';

export const init: ClientInit = async () => {
	console.log('client init');
	const user = localStorage.getItem('user') || '';
	const profile = localStorage.getItem('profile') || '';

	if (user) {
		setUser(JSON.parse(user));
	}

	if (profile) {
		setProfile(JSON.parse(profile));
	}
};
