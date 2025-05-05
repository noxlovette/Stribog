import { env } from '$env/dynamic/private';
import redis from '$lib/redisClient';
import { ValidateAccess, setTokenCookie } from '$lib/server';
import type { RefreshResponse } from '$lib/types';
import type { Handle, HandleFetch, RequestEvent, ServerInit } from '@sveltejs/kit';
import { redirect } from '@sveltejs/kit';

const PROTECTED_PATHS = new Set(['/u/']);

function isProtectedPath(path: string): boolean {
	return (
		PROTECTED_PATHS.has(path) ||
		Array.from(PROTECTED_PATHS).some((prefix) => path.startsWith(prefix))
	);
}
export const init: ServerInit = async () => {};

export const handle: Handle = async ({ event, resolve }) => {
	const path = event.url.pathname;

	if (!isProtectedPath(path)) {
		return resolve(event);
	}

	const user = await getUserFromToken(event);
	if (!user) {
		throw redirect(302, '/auth/login');
	}

	const response = await resolve(event);
	return response;
};

async function handleTokenRefresh(event: RequestEvent) {
	const refreshToken = event.cookies.get('refresh_token');
	if (!refreshToken) {
		console.debug('no refresh token, redirecting to login');
		throw redirect(302, '/auth/login');
	}

	const refreshCacheKey = `refresh:${refreshToken}`;
	console.debug('set redis cache key');
	const inProgress = await redis.get(refreshCacheKey);

	if (inProgress === 'true') {
		for (let i = 0; i < 10; i++) {
			await new Promise((resolve) => setTimeout(resolve, 200));

			const cachedUser = await redis.get(`${refreshCacheKey}:result`);
			if (cachedUser) {
				return JSON.parse(cachedUser);
			}
		}
		await redis.del(refreshCacheKey);
	}

	await redis.set(refreshCacheKey, 'true', 'EX', 5);

	try {
		console.debug('sending request to refresh');
		const refreshRes = await event.fetch('/auth/refresh', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			}
		});

		if (!refreshRes.ok) {
			console.debug('refresh failed after request');
			throw new Error('Refresh failed');
		}

		const { accessToken } = (await refreshRes.json()) as RefreshResponse;

		if (!accessToken) {
			throw new Error('No access token in refresh response');
		}
		setTokenCookie(event.cookies, 'access_token', accessToken);

		const user = await ValidateAccess(accessToken.token);

		await redis.set(`${refreshCacheKey}:result`, JSON.stringify(user), 'EX', 3);

		return user;
	} catch (error) {
		console.error('Token refresh failed:', error);
		throw redirect(302, '/auth/login');
	} finally {
		await redis.del(refreshCacheKey);
	}
}

async function getUserFromToken(event: RequestEvent) {
	const accessToken = event.cookies.get('access_token');
	let user = null;

	if (accessToken) {
		try {
			user = await ValidateAccess(accessToken);
		} catch {
			user = await handleTokenRefresh(event);
		}
	} else if (event.cookies.get('refresh_token')) {
		user = await handleTokenRefresh(event);
	}

	return user;
}

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
	const url = new URL(request.url);

	if (url.pathname.startsWith('/backend/')) {
		const cleanPath = url.pathname.replace('/backend/', '/');
		const newUrl = new URL(cleanPath, env.BACKEND_URL);

		url.searchParams.forEach((value, key) => {
			newUrl.searchParams.set(key, value);
		});

		request = new Request(newUrl.toString(), request);
	}

	const accessToken = event.cookies.get('access_token');
	if (accessToken) {
		request.headers.set('Authorization', `Bearer ${accessToken}`);
	}

	return fetch(request);
};
