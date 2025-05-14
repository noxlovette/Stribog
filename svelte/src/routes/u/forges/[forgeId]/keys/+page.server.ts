import { handleApiResponse, isSuccessResponse } from '@noxlovette/svarog';
import type { NewResponse } from '$lib/types';
import { fail, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ fetch, params, request }) => {
		// const formData = await request.formData();
		// const title = formData.get('title');
		const title = 'test';
		console.log('keys');

		const response = await fetch(`/backend/api/forge/${params.forgeID}/api-keys/`, {
			method: 'POST',
			body: JSON.stringify({ title })
		});

		console.log(response);
		const newResult = await handleApiResponse<NewResponse>(response);

		if (!isSuccessResponse(newResult)) {
			return fail(newResult.status, { message: newResult.message });
		}

		const { id } = newResult.data;
		console.log(id);

		if (response.ok) {
			return id;
		}
	}
};
