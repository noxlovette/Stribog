import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ fetch }) => {
	const forges = await fetch('/axum/forge').then((res) => res.json() as Promise<Forge[]>);
};
