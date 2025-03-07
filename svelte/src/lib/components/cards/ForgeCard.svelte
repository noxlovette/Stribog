<script lang="ts">
	import type { Forge } from "$lib/types";
    import {page} from "$app/state"

const {forge}: {forge: Forge} = $props();
const chosen = $derived(page.params.forgeId === forge.id);
</script>

<a
	href="/u/forges/{forge.id}"
	class="flex size-full justify-center flex-col items-start py-2 px-4 md:py-4 md:px-8 rounded
            ring-2 
				transition-all duration-300
                hover:ring-teal-800
                hover:bg-teal-600/20
                {chosen ? 'ring-teal-500 bg-teal-600/10' : 'ring-zinc-600'}
				"
>
	<h2 class=" text-xl md:text-3xl font-monospace font-semibold">{forge.title}</h2>
	<h3 class="capitalize">{forge.description}</h3>
</a>