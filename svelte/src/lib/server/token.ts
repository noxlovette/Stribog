import { env } from '$env/dynamic/private';
import { importSPKI, jwtVerify } from 'jose';
import type { TokenValidationConfig } from '@noxlovette/svarog';

export const tokenConfig: TokenValidationConfig = {
	spki: env.SPKI || '',
	alg: env.ALG || ''
};

export async function ValidateAccess(jwt: string) {
	const spki = env.spki || '';
	const alg = env.alg || 'RS256';
	const publicKey = await importSPKI(spki, alg);

	const { payload } = await jwtVerify(jwt, publicKey, {
		issuer: 'auth:auth'
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
