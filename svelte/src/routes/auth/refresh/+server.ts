import type { RefreshResponse } from '$lib/types';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ cookies, fetch }) => {
	const RefreshToken = cookies.get('refresh_tokeh');

	const response = await fetch('/backend/auth/refresh', {
		method: 'POST',
		body: JSON.stringify({ RefreshToken })
	});

	const { accessToken } = (await response.json()) as RefreshResponse;

	return json({ success: true, accessToken });
};
