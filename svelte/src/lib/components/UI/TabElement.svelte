<script lang="ts">
	import type { TabInsert } from '$lib/types';
	import type { Snippet } from 'svelte';
	import { page } from '$app/state';

	const { children, element }: { children?: Snippet; element: TabInsert } = $props();

	const chosen = $derived(page.url.href.includes(element.title.toLocaleLowerCase()));
</script>

<a
	href={element?.href}
	class="flex items-center justify-center rounded-b-md
	bg-stone-900/60 px-5 py-3 text-center transition-all hover:text-orange-300 {chosen
		? ' text-orange-200'
		: ''}"
>
	{element?.title}
</a>
