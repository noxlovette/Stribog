<script lang="ts">
	const { data, children } = $props();
	import { H2, SparkCard, Button, Collaborators, ApiKeys } from '$lib/components';
</script>

<div class="grid grid-cols-4">
	<div class="col-span-1 flex flex-col space-y-5 p-3">
		<div class="flex items-center justify-between">
			<H2>Sparks</H2>
			<Button variant="primary">Create Spark</Button>
		</div>

		<div class="grid min-w-max grid-cols-1 gap-5">
			{#each data.sparks as spark}
				<SparkCard {spark} />
			{/each}
		</div>
	</div>
	<div class="col-span-2 p-3">
		{@render children?.()}
	</div>
	<div class="col-span-1 flex flex-col space-y-6 p-3">
		<ApiKeys apiKeys={data.apiKeys} />
		<Collaborators collaborators={data.collaborators} />
	</div>
</div>
