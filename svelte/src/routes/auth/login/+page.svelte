<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	import { Input, Button } from '$lib/components';
	import H3 from '$lib/components/typography/H3.svelte';
	import { notification } from '$lib/stores';
	import { enhanceForm } from '$lib/utils';
</script>

<H3>Welcome Back</H3>
<form
	method="POST"
	class="flex max-w-md min-w-72 flex-col items-center justify-center space-y-3"
	use:enhance={enhanceForm({
		handlers: {
			success: async (result) => {
				if (result.data) {
					notification.set({ message: 'Samovar on the way...', type: 'success' });
					await goto('/u/forges');
				}
			}
		}
	})}
>
	<div class="my-8 grid grid-cols-1 gap-3 md:grid-cols-2">
		<Input name="email" placeholder="Email" value="" />
		<Input name="password" placeholder="Password" value="" type="password" />
	</div>
	<Button size="lg" type="submit" variant="primary">Login</Button>
</form>

<svelte:head>
	<title>Login</title>
</svelte:head>
