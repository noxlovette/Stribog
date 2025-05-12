<script lang="ts">
	import { Button } from '$lib/components';
	import type { ApiKey } from '$lib/types';
	import { formatDate } from '@noxlovette/svarog';
	import { Activity, Copy, Delete } from 'lucide-svelte';

	const { apiKey }: { apiKey: ApiKey } = $props();

	function handleCopy() {
		navigator.clipboard.writeText(apiKey.id);
	}
</script>

<div
	class="flex flex-col rounded-sm border border-stone-200 bg-white p-4 dark:border-stone-700 dark:bg-stone-900"
>
	<div class="flex items-start justify-between">
		<div>
			<h3 class="text-lg font-medium text-stone-900 dark:text-white">{apiKey.title}</h3>
			<p class="truncate font-mono text-sm text-stone-500 dark:text-stone-400">{apiKey.id}</p>
		</div>
		<div class="flex space-x-2">
			<Button Icon={Copy} variant="ghost" size="sm" onclick={handleCopy}>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="18"
					height="18"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					><rect width="14" height="14" x="8" y="8" rx="2" ry="2" /><path
						d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"
					/></svg
				>
				Copy
			</Button>
		</div>
	</div>
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
		<Button Icon={Activity} variant="outline" size="sm">
			{apiKey.isActive ? 'Deactivate' : 'Activate'}
		</Button>
		<Button Icon={Delete} variant="danger" size="sm">Delete</Button>
	</div>
</div>
