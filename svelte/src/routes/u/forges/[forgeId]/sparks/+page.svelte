<script lang="ts">
	import { H2, SparkCard, Button } from '$lib/components';
	import AddCard from '$lib/components/cards/AddCard.svelte';

	const { data } = $props();
</script>

<div class="grid min-w-max gap-3 md:grid-cols-2 xl:grid-cols-3">
	<AddCard></AddCard>
	{#each data.sparks as spark}
		<SparkCard {spark} />
	{/each}
</div>
