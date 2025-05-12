import type { LayoutServerLoad } from './$types';
import type { Forge, Spark, Collaborator, ApiKey } from '$lib/types';

export const load: LayoutServerLoad = async ({ fetch, params, depends }) => {
	depends('forge:general');
	const forgeID = params.forgeID;
	const endpoints = {
		forge: `/backend/api/forge/${forgeID}`,
		sparks: `/backend/api/forge/${forgeID}/sparks`,
		collaborators: `/backend/api/forge/${forgeID}/access`,
		apiKeys: `/backend/api/forge/${forgeID}/api-keys`
	};

	const [forge, sparks, collaborators, apiKeys] = await Promise.all([
		fetch(endpoints.forge).then((res) => res.json() as Promise<Forge>),
		fetch(endpoints.sparks).then((res) => res.json() as Promise<Spark[]>),
		fetch(endpoints.collaborators).then((res) => res.json() as Promise<Collaborator[]>),
		fetch(endpoints.apiKeys).then((res) => res.json() as Promise<ApiKey[]>)
	]);

	return { forge, sparks, collaborators, apiKeys };
};
