import type { SignupResponse } from '$lib/types';
import {
	validateEmail,
	validatePassword,
	validatePasswordMatch,
	validateUsername,
	isSuccessResponse,
	handleApiResponse,
	turnstileVerify
} from '@noxlovette/svarog';
import { env } from '$env/dynamic/private';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';
export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const data = await request.formData();
		const username = data.get('username') as string;
		const pass = data.get('password') as string;
		const confirmPassword = data.get('confirmPassword') as string;
		const email = data.get('email') as string;
		const role = data.get('role') as string;
		const name = data.get('name') as string;

		if (validateEmail(email)) {
			return fail(400, { message: 'Invalid Email' });
		}

		const usernameValidation = validateUsername(username);
		const passValidation = validatePassword(pass);
		const passMatchValidation = validatePasswordMatch(pass, confirmPassword);

		if (usernameValidation) {
			return fail(400, { message: usernameValidation });
		}

		if (pass !== confirmPassword) {
			return fail(400, { message: 'Passwords do not match' });
		}

		if (passValidation) {
			return fail(400, { message: passValidation });
		}

		if (passMatchValidation) {
			return fail(400, { message: passMatchValidation });
		}

		const turnstileToken = data.get('cf-turnstile-response') as string;
		if (!turnstileToken) {
			return fail(400, {
				message: 'Please complete the CAPTCHA verification'
			});
		}
		const turnstileResponse = await turnstileVerify(turnstileToken, env.TURNSTILE_SECRET);
		if (!turnstileResponse.ok) {
			return fail(400, {
				message: 'Turnstile verification failed'
			});
		}

		// Signup API call
		const response = await fetch('/backend/auth/signup', {
			method: 'POST',
			body: JSON.stringify({ username, pass, email, role, name })
		});

		const result = await handleApiResponse<SignupResponse>(response);

		if (!isSuccessResponse(result)) {
			return fail(result.status, { message: result.message });
		}

		return redirect(302, '/auth/login');
	}
} satisfies Actions;
