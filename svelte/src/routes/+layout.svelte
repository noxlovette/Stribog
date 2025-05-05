<script lang="ts">
	import '../app.css';
	import { SEO } from '@noxlovette/svarog';
	import { Notification, Header, Footer } from '$lib/components';

	interface Props {
		children?: import('svelte').Snippet;
	}
	let { children }: Props = $props();
</script>

<SEO title="Stribog" description="CDN" keywords="CDN, stribog" />

<main class="flex flex-col overflow-auto bg-stone-950 text-stone-50 antialiased">
	<Header />
	<div
		class="flex min-h-screen flex-col font-medium caret-orange-800 selection:bg-orange-500/20 dark:caret-stone-200"
	>
		{@render children?.()}
	</div>
	<Footer />
	<Notification />
</main>
