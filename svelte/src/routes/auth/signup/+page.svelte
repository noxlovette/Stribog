<script lang="ts">
	import { enhance } from '$app/forms';
	import { Button, Input, H3 } from '$lib/components';

	import { enhanceForm } from '$lib/utils';

	let password = $state('');
	let confirmPassword = $state('');
	let passwordMatch = $state(true);
</script>

<H3>Join Us</H3>
<form
	method="post"
	class="flex flex-col items-center justify-center space-y-4"
	use:enhance={enhanceForm({
		messages: {
			redirect: 'Welcome on board',
			defaultError: 'Signup Failed'
		}
	})}
>
	<div class="my-8 grid grid-cols-1 gap-3 md:grid-cols-2">
		<Input name="name" placeholder="Name" value=""></Input>
		<Input name="email" placeholder="Email" type="email" value=""></Input>
		<Input name="password" placeholder="Password" type="password" bind:value={password}></Input>
		<Input name="confirmPassword" placeholder="Again" type="password" bind:value={confirmPassword}
		></Input>

		{#if !passwordMatch}
			<p class="mt-1 text-sm text-red-600">Passwords don't match</p>
		{/if}
	</div>
	<Button size="lg" type="submit" variant="primary">Create Account</Button>
</form>

<svelte:head>
	<title>Signup</title>
</svelte:head>
