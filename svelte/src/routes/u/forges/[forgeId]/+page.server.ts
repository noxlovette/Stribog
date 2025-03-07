import { handleApiResponse, isSuccessResponse } from '@noxlovette/svarog';
import type { NewResponse, Spark, ApiKey, Forge } from '$lib/types';
import { fail, redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	newKey: async ({ fetch, params, request }) => {
		const formData = await request.formData();

		const title = formData.get('title');

		const body = {
			title
		};

		const response = await fetch(`/axum/key/${params.forgeId}`, {
			method: 'POST',
			body: JSON.stringify(body)
		});

		const newResult = await handleApiResponse<NewResponse>(response);

		console.log(newResult);

		if (!isSuccessResponse(newResult)) {
			return fail(newResult.status, { message: newResult.message });
		}

		if (response.ok) {
			return { success: true };
		}
	}
};
