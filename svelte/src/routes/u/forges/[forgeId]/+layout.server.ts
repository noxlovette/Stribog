import type { LayoutServerLoad } from './$types';
import type { Forge, Spark, Collaborator, ApiKey } from '$lib/types';

export const load: LayoutServerLoad = async ({ fetch, params }) => {
	const forgeId = params.forgeId;
	const endpoints = {
		forge: `/axum/forge/${forgeId}`,
		sparks: `/axum/spark/${forgeId}`,
		collaborators: `/axum/forge/${forgeId}/access`,
		apiKeys: `/axum/key/${forgeId}`
	};

	const [forge, sparks, collaborators, apiKeys] = await Promise.all([
		fetch(endpoints.forge).then((res) => res.json() as Promise<Forge>),
		fetch(endpoints.sparks).then((res) => res.json() as Promise<Spark[]>),
		fetch(endpoints.collaborators).then((res) => res.json() as Promise<Collaborator[]>),
		fetch(endpoints.apiKeys).then((res) => res.json() as Promise<ApiKey[]>)
	]);

	return { forge, sparks, collaborators, apiKeys };
};
