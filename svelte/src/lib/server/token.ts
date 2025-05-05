import { env } from '$env/dynamic/private';
import { jwtVerify } from 'jose';
import type { TokenValidationConfig } from '@noxlovette/svarog';

export const tokenConfig: TokenValidationConfig = {
	spki: '',
	alg: 'HS256'
};

export async function ValidateAccess(jwt: string) {
	const secret = new TextEncoder().encode(env.JWT_KEY || '');
	const alg = 'HS256';

	const { payload } = await jwtVerify(jwt, secret, {
		issuer: 'auth:auth',
		algorithms: [alg]
	});

	const EXPIRY_BUFFER = 30;
	if (payload.exp && typeof payload.exp === 'number') {
		const now = Math.floor(Date.now() / 1000);
		if (payload.exp - now < EXPIRY_BUFFER) {
			throw new Error('Token about to expire');
		}
	}

	return payload;
}
