<script lang="ts">
	import { ForgeCard, H1, H3 } from '$lib/components';
	import AddCard from '$lib/components/cards/AddCard.svelte';

	const { data } = $props();
</script>

<div class="flex flex-col items-center justify-center py-3">
	<H1>Forges</H1>
	<H3>Your Projects</H3>
	<div class="my-4 grid grid-cols-2 gap-4">
		{#each data.forges as forge}
			<ForgeCard {forge} />
		{/each}
		<AddCard href="/u/forges/create" />
	</div>
</div>
