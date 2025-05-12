import type { User } from '$lib/types';
import { writable } from 'svelte/store';

export const initialUser: User = {
	name: '',
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
		name: '',
		email: ''
	}));
}
