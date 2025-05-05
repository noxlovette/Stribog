<script lang="ts">
	import { user } from '$lib/stores';
</script>

{#snippet link(href: string, name: string)}
	<a
		{href}
		class="p-3 text-center text-sm transition-colors hover:bg-stone-950 hover:ring hover:ring-stone-800/70"
		>{name}</a
	>
{/snippet}
<header class="flex items-center justify-center bg-stone-900">
	<div class="flex w-full max-w-md items-center justify-center">
		<nav class="relative grid w-full grid-cols-3">
			{#if $user && $user.username}
				{@render link('/u/dashboard', 'Izba')}
				{@render link('/u/forges', 'Forges')}
				{@render link('/u/account', 'Account')}
			{:else}
				{@render link('/', 'Home')}
				{@render link('/auth/login', 'Login')}
				{@render link('/auth/signup', 'Sign Up')}
			{/if}
		</nav>
	</div>
</header>
