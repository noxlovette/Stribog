import { env } from '$env/dynamic/private';
import { getUserFromToken } from '@noxlovette/svarog';
import type { Handle, HandleFetch } from '@sveltejs/kit';
import { redirect } from '@sveltejs/kit';
import { tokenConfig } from '$lib/server/token';

const PROTECTED_PATHS = new Set(['/u/dashboard']);

function isProtectedPath(path: string): boolean {
	return (
		PROTECTED_PATHS.has(path) ||
		Array.from(PROTECTED_PATHS).some((prefix) => path.startsWith(prefix))
	);
}

export const handle: Handle = async ({ event, resolve }) => {
	console.log('handle');
	const path = event.url.pathname;

	// if (path === '/') {
	// 	const user = await getUserFromToken(event, tokenConfig);
	// 	if (user) {
	// 		throw redirect(303, '/u/dashboard');
	// 	}
	// }

	if (!isProtectedPath(path)) {
		console.log('started');
		return resolve(event);
	}

	const user = await getUserFromToken(event, tokenConfig);
	if (!user) {
		throw redirect(302, '/auth/login');
	}

	const response = await resolve(event);
	return response;
};

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
	const url = new URL(request.url);

	if (url.pathname.startsWith('/axum/')) {
		const cleanPath = url.pathname.replace('/axum/', '/');
		// Create new URL with the base and path
		const newUrl = new URL(cleanPath, env.BACKEND_URL);

		// IMPORTANT: Copy all search parameters from original URL
		url.searchParams.forEach((value, key) => {
			newUrl.searchParams.set(key, value);
		});

		// Create new request with the full URL including query parameters
		request = new Request(newUrl, request);
	}

	request.headers.set('X-API-KEY', env.API_KEY_AXUM);
	request.headers.set('Content-Type', 'application/json');
	const accessToken = event.cookies.get('accessToken');
	if (accessToken) {
		request.headers.set('Authorization', `Bearer ${accessToken}`);
	}
	return fetch(request);
};
