import type { LayoutServerLoad } from './$types';
import type { Forge } from '$lib/types';

export const load: LayoutServerLoad = async ({ fetch }) => {
	const forges = await fetch('/backend/api/forge').then((res) => res.json() as Promise<Forge[]>);

	return { forges };
};
