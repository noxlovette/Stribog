import type { Spark } from '$lib/types';
import { parseMarkdown } from '@noxlovette/svarog';
import type { LayoutServerLoad } from './$types';

export const load = (async ({ params, fetch }) => {
	const spark = await fetch(`/backend/api/sparks/${params.sparkID}`).then(
		(res) => res.json() as Promise<Spark>
	);

	const rendered = await parseMarkdown(spark.markdown);
	return { spark, rendered };
}) satisfies LayoutServerLoad;
