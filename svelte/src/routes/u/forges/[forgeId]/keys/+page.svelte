<script lang="ts">
	import { ApiKeyCard } from '$lib/components';
	import AddCard from '$lib/components/cards/AddCard.svelte';

	const { data } = $props();
	const forge = $derived(data.forge);
</script>

<div class="grid grid-cols-3 gap-3">
	<AddCard></AddCard>
	{#each data.apiKeys as apiKey}
		<ApiKeyCard {apiKey}></ApiKeyCard>
	{/each}
</div>
