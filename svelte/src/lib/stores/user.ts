import type { User } from '$lib/types';
import { writable } from 'svelte/store';

export const initialUser: User = {
	username: '',
	sub: '',
	name: '',
	role: '',
	email: ''
};

export const user = writable<User>(initialUser);

export function setUser(data: User) {
	user.update((currentState) => ({
		...currentState,
		...data
	}));
}

export function clearUser() {
	user.update(() => ({
		username: '',
		sub: '',
		name: '',
		role: '',
		email: ''
	}));
}
