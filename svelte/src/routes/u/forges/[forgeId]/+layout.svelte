<script lang="ts">
	const { data, children } = $props();
	import { H1, H2, H3, H4, TabElement, Tabs } from '$lib/components';
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

<div class="flex flex-col items-center gap-3 py-3 text-center">
	<a href="/u/forges/{page.params.forgeID}" class="group max-w-1/2">
		<H1 styling="transition-all group-hover:text-orange-300 {chosen ? 'text-orange-200' : ''}"
			>{data.forge.title}</H1
		>
	</a>
	<Tabs>
		{#each elements as element}
			<TabElement {element}></TabElement>
		{/each}
	</Tabs>
</div>
{@render children?.()}
