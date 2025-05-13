<script lang="ts">
	import { H2, H3, Button, SparkCard, ApiKeys, Collaborators, Tabs } from '$lib/components';
	import TabElement from '$lib/components/UI/TabElement.svelte';
	const { data } = $props();
</script>

<grid class="grid grid-cols-3 gap-3">
	<div class="flex flex-col gap-3">
		<H2>Sparks</H2>
		<div class="grid min-w-max grid-cols-1 gap-5">
			{#each data.sparks.slice(0, 3) as spark}
				<SparkCard {spark} />
			{/each}
		</div>
	</div>
	<ApiKeys apiKeys={data.apiKeys} />
	<Collaborators collaborators={data.collaborators} />
</grid>
