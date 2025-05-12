import type { Spark } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load = (async ({ params, fetch }) => {
	const spark = await fetch(`/backend/api/sparks/${params.sparkID}`).then(
		(res) => res.json() as Promise<Spark>
	);
	return { spark };
}) satisfies PageServerLoad;
