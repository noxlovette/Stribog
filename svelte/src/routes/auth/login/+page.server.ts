import {
	handleApiResponse,
	isSuccessResponse,
	parseCookieOptions,
	turnstileVerify,
	ValidateAccess,
	validateRequired
} from '@noxlovette/svarog';
import type { AuthResponse, User } from '$lib/types';
import { fail, type Actions } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';
import { tokenConfig } from '$lib/server/token';

export const actions: Actions = {
	default: async ({ request, fetch, cookies }) => {
		try {
			const data = await request.formData();
			const username = data.get('username') as string;
			const pass = data.get('password') as string;
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
			const passError = validatePass(pass);
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

			// Login API call with typed response handling
			const response = await fetch('/axum/auth/signin', {
				method: 'POST',
				body: JSON.stringify({ username, pass })
			});

			const authResult = await handleApiResponse<AuthResponse>(response);

			console.log(authResult);
			if (!isSuccessResponse(authResult)) {
				return fail(authResult.status, { message: authResult.message });
			}

			// Handle cookies from response
			response.headers.getSetCookie().forEach((cookie) => {
				const [fullCookie, ...opts] = cookie.split(';');
				const [name, value] = fullCookie.split('=');
				const cookieOptions = parseCookieOptions(opts);
				cookies.set(name, value, cookieOptions);
			});

			// Validate the access token
			const { accessToken } = authResult.data;
			console.log(tokenConfig);
			const user = (await ValidateAccess(
				accessToken,
				tokenConfig.spki,
				tokenConfig.alg,
				tokenConfig.issuer
			)) as User;

			if (!user) {
				return fail(401, {
					message: 'Invalid access token'
				});
			}

			return {
				success: true,
				message: 'Login successful',
				user
			};
		} catch (error) {
			console.error('Signin error:', error);
			return fail(500, {
				message: error instanceof Error ? error.message : 'Internal server error'
			});
		}
	}
};
