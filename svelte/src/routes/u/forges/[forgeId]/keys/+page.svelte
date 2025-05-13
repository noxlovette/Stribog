<script lang="ts">
	import { Button, Input } from '$lib/components';
	import { enhance } from '$app/forms';
	import { enhanceForm } from '$lib/utils';
	import ApiKeyCard from '$lib/components/cards/ApiKeyCard.svelte';

	const { data } = $props();
	const forge = $derived(data.forge);
</script>

<div class="grid grid-cols-3 gap-3">
	{#each data.apiKeys as apiKey}
		<ApiKeyCard {apiKey}></ApiKeyCard>
	{:else}
		No API Keys
	{/each}
</div>
