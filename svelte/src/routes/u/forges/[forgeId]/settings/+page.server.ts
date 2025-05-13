import { handleApiResponse, isSuccessResponse, type EmptyResponse } from '@noxlovette/svarog';
import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
	default: async ({ request, fetch, params }) => {
		const formData = await request.formData();
		const title = formData.get('title') as string;
		const description = formData.get('description') as string;

		const response = await fetch(`/backend/api/forge/${params.forgeID}`, {
			method: 'PATCH',
			body: JSON.stringify({
				title,
				description
			})
		});
		const result = await handleApiResponse<EmptyResponse>(response);
		if (!isSuccessResponse(result)) {
			return fail(result.status, { message: result.message });
		}

		if (response.ok) {
			return { success: true };
		}
	}
} satisfies Actions;
