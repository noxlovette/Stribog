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
			class={`inline-block h-2 w-2 rounded-full ${apiKey.is_active ? 'bg-green-500' : 'bg-red-500'} mr-2`}
		></span>
		<span class="text-sm text-stone-600 dark:text-stone-300"
			>{apiKey.is_active ? 'Active' : 'Inactive'}</span
		>
	</div>
	<div class="mt-2 text-xs text-stone-500 dark:text-stone-400">
		<div>Created: {formatDate(apiKey.createdAt)}</div>
		{#if apiKey.lastUsedAt}
			<div>Last used: {formatDate(apiKey.lastUsedAt)}</div>
		{:else}
			<div>Never used</div>
		{/if}
	</div>
	<div class="mt-4 flex justify-between">
		<Button variant="outline" size="sm">
			{apiKey.is_active ? 'Deactivate' : 'Activate'}
		</Button>
		<Button variant="danger" size="sm">Delete</Button>
	</div>
</Card>
