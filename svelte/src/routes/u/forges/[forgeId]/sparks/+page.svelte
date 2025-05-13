<script lang="ts">
	import { H2, SparkCard, Button } from '$lib/components';

	const { data } = $props();
</script>

<div class="col-span-1 flex flex-col gap-3">
	<div class="flex items-center justify-between">
		<H2>Sparks</H2>
		<Button variant="primary">Create Spark</Button>
	</div>

	<div class="grid min-w-max gap-5 md:grid-cols-2 xl:grid-cols-3">
		{#each data.sparks as spark}
			<SparkCard {spark} />
		{/each}
	</div>
</div>
