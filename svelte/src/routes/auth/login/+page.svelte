<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	import { Input, Turnstile, Button, H1 } from '$lib/components';
	import H2 from '$lib/components/typography/H2.svelte';
	import H3 from '$lib/components/typography/H3.svelte';
	import { setUser, initialUser, notification } from '$lib/stores';
	import { enhanceForm } from '$lib/utils';
</script>

<H3>Welcome Back</H3>
<form
	method="POST"
	class="max-Yw-md flex min-w-72 flex-col items-center justify-center space-y-3"
	use:enhance={enhanceForm({
		handlers: {
			success: async (result) => {
				if (result.data) {
					const { user = initialUser } = result.data;
					setUser(user);
					localStorage.setItem('user', JSON.stringify(user));
					notification.set({ message: 'Samovar on the way...', type: 'success' });
					await goto('/u/dashboard');
				}
			}
		}
	})}
>
	<Input name="username" placeholder="Username" value="" />
	<Input name="password" placeholder="Password" value="" type="password" />
	<Button size="lg" type="submit" variant="primary">Login</Button>
</form>

<svelte:head>
	<title>Login</title>
</svelte:head>
