<script lang="ts">
	import type { Forge } from '$lib/types';
	import { page } from '$app/state';

	const { forge }: { forge: Forge } = $props();
	const chosen = $derived(page.params.forgeID === forge.id);
</script>

<a
	href="/u/forges/{forge.id}"
	class="group inline-flex min-h-36 flex-col rounded-sm px-3 py-2 ring transition-colors hover:bg-orange-600/20 hover:ring-orange-800
        {chosen ? 'bg-orange-600/10 ring-orange-500' : 'ring-stone-600'}"
>
	<div class="origin-left transition-all duration-300">
		<h2 class="font-semibold md:text-3xl">{forge.title}</h2>
		<h3>{forge.description}</h3>
	</div>
</a>
