<script lang="ts">
	import { Button } from '$lib/components';
	import Card from './Card.svelte';
	import type { Collaborator } from '$lib/types';
	const { collaborator }: { collaborator: Collaborator } = $props();
	$inspect(collaborator);
</script>

<Card>
	<div class="flex items-start justify-between">
		<div class="flex items-center">
			<div
				class="flex h-10 w-10 items-center justify-center rounded-full bg-stone-200 text-stone-500 dark:bg-stone-700 dark:text-stone-400"
			>
				{collaborator?.userName?.substring(0, 2).toUpperCase()}
			</div>
			<div class="ml-3">
				<h3 class="text-lg font-medium text-stone-900 dark:text-white">
					{collaborator.userName}
				</h3>
			</div>
		</div>
		<div>
			<span
				class="inline-flex items-center rounded-full bg-blue-100 px-2.5 py-0.5 text-xs font-medium text-blue-800 dark:bg-blue-800 dark:text-blue-100"
			>
				{collaborator?.accessRole}
			</span>
		</div>
	</div>
	<div class="mt-2 text-sm text-stone-300">
		<p>{collaborator?.userEmail}</p>
	</div>
</Card>
