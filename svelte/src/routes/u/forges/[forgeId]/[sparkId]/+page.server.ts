import type { Spark } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load = (async ({ params, fetch }) => {
	const spark = await fetch(`/axum/spark/s/${params.sparkId}`).then(
		(res) => res.json() as Promise<Spark>
	);
	return { spark };
}) satisfies PageServerLoad;
