import { handleApiResponse, isSuccessResponse } from '@noxlovette/svarog';
import type { NewResponse } from '$lib/types';
import { fail, redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ fetch, params }) => {
		const response = await fetch(`/backend/api/forge/${params.forgeID}/sparks`, {
			method: 'POST'
		});

		const newResult = await handleApiResponse<NewResponse>(response);

		if (!isSuccessResponse(newResult)) {
			return fail(newResult.status, { message: newResult.message });
		}

		const { id } = newResult.data;

		if (response.ok) {
			return redirect(301, `/u/forges/${params.forgeID}/sparks/${id}/edit`);
		}
	}
};
