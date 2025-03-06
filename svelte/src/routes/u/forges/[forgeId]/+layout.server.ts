import type { LayoutServerLoad } from './$types';
import type { Forge, Spark } from '$lib/types';

export const load: LayoutServerLoad = async ({ fetch, params }) => {
	const forge = await fetch(`/axum/forge/${params.forgeId}`).then(
		(res) => res.json() as Promise<Forge>
	);

	const sparks = await fetch(`/axum/spark/${params.forgeId}`).then(
		(res) => res.json() as Promise<Spark[]>
	);

	return { forge, sparks };
};
