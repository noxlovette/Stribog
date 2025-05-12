<script lang="ts">
	import { Button, H2, CollaboratorCard } from '$lib/components';
	import type { Collaborator } from '$lib/types';

	const { collaborators = [] }: { collaborators: Collaborator[] } = $props();
</script>

<div class="flex items-center justify-between">
	<H2>Collaborators</H2>
	<Button variant="outline" size="sm">Add User</Button>
</div>

{#if collaborators && collaborators.length > 0}
	<div class="grid grid-cols-1 gap-4">
		{#each collaborators as collaborator}
			<CollaboratorCard {collaborator} />
		{/each}
	</div>
{:else}
	<div
		class="flex flex-col items-center justify-center rounded-sm border border-dashed border-stone-200 bg-white p-4 dark:border-stone-700 dark:bg-stone-800"
	>
		<p class="text-sm text-stone-500 dark:text-stone-400">No collaborators yet</p>
	</div>
{/if}
