<script lang="ts">
	import '../app.css';
	import { SEO } from '@noxlovette/svarog';
	import { Notification, Header, Footer } from '$lib/components';
	import { keyEscape } from '$lib/actions';

	function goBack() {
		window.history.back();
	}

	interface Props {
		children?: import('svelte').Snippet;
	}
	let { children }: Props = $props();
</script>

<SEO title="Stribog" description="CDN" keywords="CDN, stribog" />

<main
	class="flex flex-col h-screen w-screen items-center justify-center overflow-auto bg-zinc-50 text-zinc-800 antialiased dark:bg-zinc-950 dark:text-zinc-100"
	use:keyEscape={goBack}
>
	<Header />
	<div
		class="flex size-full max-w-7xl flex-col font-medium caret-teal-800 selection:bg-teal-500/20 dark:caret-zinc-200"
	>
			{@render children?.()}
	</div>
	<Footer />
	<Notification />
</main>
