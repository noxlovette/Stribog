import { handleApiResponse, isSuccessResponse, validateRequired } from '@noxlovette/svarog';
import type { AuthResponse } from '$lib/types';
import { fail, type Actions } from '@sveltejs/kit';
import { setTokenCookie } from '$lib/server';

export const actions: Actions = {
	default: async ({ request, fetch, cookies }) => {
		try {
			const data = await request.formData();
			const username = data.get('username') as string;
			const password = data.get('password') as string;
			const turnstileToken = data.get('cf-turnstile-response') as string;

			// const turnstileTokenError = validateRequired(turnstileToken);
			// if (turnstileTokenError) {
			// 	return fail(400, {
			// 		message: turnstileTokenError
			// 	});
			// }

			const validateUsername = validateRequired('username');
			const validatePass = validateRequired('password');

			const usernameError = validateUsername(username);
			const passError = validatePass(password);
			if (usernameError) {
				return fail(400, {
					message: usernameError
				});
			}

			if (passError) {
				return fail(400, {
					message: passError
				});
			}

			// const turnstileResponse = await turnstileVerify(turnstileToken, env.CLOUDFLARE_SECRET);

			// if (!turnstileResponse.ok) {
			// 	return fail(400, {
			// 		message: 'Turnstile verification failed'
			// 	});
			// }

			const response = await fetch('/backend/auth/login', {
				method: 'POST',
				body: JSON.stringify({ username, password })
			});

			const authResult = await handleApiResponse<AuthResponse>(response);

			if (!isSuccessResponse(authResult)) {
				return fail(authResult.status, { message: authResult.message });
			}
			const { accessToken, refreshToken } = authResult.data;
			setTokenCookie(cookies, 'access_token', accessToken);
			setTokenCookie(cookies, 'refresh_token', refreshToken);

			return {
				success: true,
				message: 'Login successful'
			};
		} catch (error) {
			console.error('Signin error:', error);
			return fail(500, {
				message: error instanceof Error ? error.message : 'Internal server error'
			});
		}
	}
};
