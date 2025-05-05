import type { User } from '$lib/types';
import type { LayoutServerLoadEvent } from './$types';

export const load: LayoutServerLoadEvent = async ({ fetch }) => {
	const url = `/backend/api/me`;

	const user = await fetch(url).then((res) => res.json() as User);

	return {
		user
	};
};
