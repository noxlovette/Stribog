<script lang="ts">
	import { Button } from '$lib/components';
	import type { ApiKey } from '$lib/types';
	import { formatDate } from '@noxlovette/svarog';

	import Card from './Card.svelte';
	import H3 from '../typography/H3.svelte';

	const { apiKey }: { apiKey: ApiKey } = $props();
</script>

<Card>
	<H3>{apiKey.title}</H3>

	<div class="mt-2 flex items-center">
		<span
			class={`inline-block h-2 w-2 rounded-full ${apiKey.isActive ? 'bg-green-500' : 'bg-red-500'} mr-2`}
		></span>
		<span class="text-sm text-stone-600 dark:text-stone-300"
			>{apiKey.isActive ? 'Active' : 'Inactive'}</span
		>
	</div>
	<div class="mt-2 text-xs text-stone-500 dark:text-stone-400">
		<div>Created: {formatDate(apiKey.createdAt)}</div>
		{#if apiKey.lastUsed}
			<div>Last used: {formatDate(apiKey.lastUsed)}</div>
		{:else}
			<div>Never used</div>
		{/if}
	</div>
</Card>
