import type { RefreshResponse } from '$lib/types';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ cookies, fetch }) => {
	const refreshToken = cookies.get('refresh_token');
	console.debug(refreshToken);
	console.debug('request has reached the server.ts');
	const response = await fetch('/backend/auth/refresh', {
		method: 'POST',
		body: JSON.stringify({ refreshToken })
	});
	console.debug('response on the server received');
	console.debug(response);
	const { accessToken } = (await response.json()) as RefreshResponse;

	return json({ success: true, accessToken });
};
