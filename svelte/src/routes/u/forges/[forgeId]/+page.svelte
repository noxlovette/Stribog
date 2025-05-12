<script lang="ts">
	import { H2, H3 } from '$lib/components';
	const { data } = $props();
</script>

<div class="flex flex-col items-center justify-center">
	<H2>{data.forge.title}</H2>
	<H3>{data.forge.description}</H3>
	<div class="size-72 overflow-hidden">
		<img
			src="/pictures/spark-logo.png"
			alt="Forge and Sparkles"
			class="h-full w-full object-cover"
		/>
	</div>
</div>
