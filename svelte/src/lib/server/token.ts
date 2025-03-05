import { env } from '$env/dynamic/private';

import type { TokenValidationConfig } from '@noxlovette/svarog';

export const tokenConfig: TokenValidationConfig = {
	spki: env.SPKI || '',
	alg: env.ALG || '',
	audience: 'auth:auth'
};
