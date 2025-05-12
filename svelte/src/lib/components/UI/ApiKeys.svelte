<script lang="ts">
	import { enhance } from '$app/forms';
	import { enhanceForm } from '$lib/utils';
	import { invalidate } from '$app/navigation';
	import { H3, Button, Input, H2, ApiKeyCard } from '$lib/components';
	import type { ApiKey } from '$lib/types';

	const { apiKeys = [] }: { apiKeys: ApiKey[] } = $props();
	let apiKeyCreationDialogue = $state(false);
</script>

<div class="flex flex-col space-y-3">
	<div class="flex flex-col">
		<div class="flex items-center justify-between">
			<H2>API Keys</H2>
			<Button
				variant="outline"
				type="button"
				onclick={() => (apiKeyCreationDialogue = true)}
				size="sm"
			>
				New Key
			</Button>
		</div>
		{#if apiKeyCreationDialogue}
			<form
				method="POST"
				class=""
				action="?/newKey"
				use:enhance={enhanceForm({
					messages: {
						success: 'Created New Key',
						failure: 'Failed to create',
						defaultError: 'Failed to create key'
					},
					handlers: {
						success: async () => {
							invalidate('forge:general');
						}
					}
				})}
			>
				<div class="space-y-3">
					<H3>Name the Key</H3>
					<Input name="title" placeholder="Name the Key" value="" />
					<div class="flex space-x-2">
						<Button variant="primary" type="submit">Create</Button>
						<Button variant="ghost" onclick={() => (apiKeyCreationDialogue = false)}>Cancel</Button>
					</div>
				</div>
			</form>
		{/if}
	</div>

	{#if apiKeys && apiKeys.length > 0}
		<div class="grid grid-cols-1 gap-4">
			{#each apiKeys as apiKey}
				<ApiKeyCard {apiKey} />
			{/each}
		</div>
	{:else}
		<div
			class="flex flex-col items-center justify-center rounded-sm border border-dashed border-stone-600/50 bg-stone-950 p-4"
		>
			<p class="text-sm text-stone-400">No API keys yet</p>
		</div>
	{/if}
</div>
