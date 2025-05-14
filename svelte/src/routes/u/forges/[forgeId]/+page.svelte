<script lang="ts">
	import SparkLogo from '$lib/components/images/SparkLogo.svelte';
	import { H3 } from '$lib/components';
	const { data } = $props();
</script>

<div class="flex w-full flex-col items-center justify-center">
	<SparkLogo></SparkLogo>
	<H3>
		{data.forge.description}
	</H3>
</div>
