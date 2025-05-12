<script lang="ts">
	import type { Forge } from '$lib/types';
	import { page } from '$app/state';

	const { forge }: { forge: Forge } = $props();
	const chosen = $derived(page.params.forgeID === forge.id);
</script>

<a
	href="/u/forges/{forge.id}"
	class="flex size-full flex-col items-start justify-center rounded px-4 py-2 ring-2 transition-all
            duration-300
				hover:bg-orange-600/20 hover:ring-orange-800
                md:px-8
                md:py-4
                {chosen ? 'bg-orange-600/10 ring-orange-500' : 'ring-stone-600'}
				"
>
	<h2 class=" font-monospace text-xl font-semibold md:text-3xl">{forge.title}</h2>
	<h3 class="capitalize">{forge.description}</h3>
</a>
