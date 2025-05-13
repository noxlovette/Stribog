<script lang="ts">
	import type { Forge } from '$lib/types';
	import { page } from '$app/state';
	import { H3, Button } from '$lib/components';
	import Card from './Card.svelte';

	const { forge }: { forge: Forge } = $props();
	const chosen = $derived(page.params.forgeID === forge.id);
</script>

<div class="flex flex-col rounded-md bg-stone-900/60 p-5">
	<div class="flex items-start justify-between">
		<div class="items-start justify-between">
			<H3>
				{forge.title}
			</H3>
			<p>
				{forge.description}
			</p>
		</div>
		<div class="flex space-x-2">
			<Button variant="ghost" size="sm" href="/u/forges/{forge.id}/settings">Edit</Button>
			<Button href="/u/forges/{forge.id}" variant="outline" size="sm">View</Button>
		</div>
	</div>
</div>
