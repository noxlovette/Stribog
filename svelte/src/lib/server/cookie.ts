import { parseISO } from 'date-fns';
import type { Cookies } from '@sveltejs/kit';

export function setTokenCookie(
	cookies: Cookies,
	name: string,
	token: { token: string; expiresAt: string }
) {
	const expiry = parseISO(token.expiresAt);
	const maxAge = Math.floor((expiry.getTime() - Date.now()) / 1000);

	cookies.set(name, token.token, {
		path: '/',
		maxAge
	});
}
