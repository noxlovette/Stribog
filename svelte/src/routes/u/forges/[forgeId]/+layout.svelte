<script lang="ts">
	const { data, children } = $props();
	import { H1, H2, H3, TabElement, Tabs } from '$lib/components';
	import { page } from '$app/state';
	import type { TabInsert } from '$lib/types';

	const elements: TabInsert[] = [
		{ title: 'Sparks', href: `/u/forges/${page.params.forgeID}/sparks` },
		{ title: 'Keys', href: `/u/forges/${page.params.forgeID}/keys` },
		{ title: 'Settings', href: `/u/forges/${page.params.forgeID}/settings` },
		{ title: 'Collaborators', href: `/u/forges/${page.params.forgeID}/collaborators` }
	];

	const chosen = $derived(page.url.pathname === `/u/forges/${page.params.forgeID}`);
</script>

<div class="mb-6 flex justify-between gap-3">
	<a
		href="/u/forges/{page.params.forgeID}"
		class="group w-1/2 max-w-1/2 rounded-b-md bg-stone-900/60 px-3 pt-1 pb-2"
	>
		<H1 styling="transition-all group-hover:text-orange-300 {chosen ? 'text-orange-200' : ''}"
			>{data.forge.title}</H1
		>
		<H3
			styling="{chosen
				? 'max-h-40 opacity-100'
				: 'max-h-0 overflow-hidden opacity-0'} transition-all duration-300 ease-in-out"
		>
			{data.forge.description}</H3
		>
	</a>
	<Tabs>
		{#each elements as element}
			<TabElement {element}></TabElement>
		{/each}
	</Tabs>
</div>
{@render children?.()}
