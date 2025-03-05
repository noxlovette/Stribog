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
	class="h-screen w-screen overflow-auto bg-zinc-50 text-sky-800 antialiased dark:bg-zinc-950 dark:text-zinc-100"
	use:keyEscape={goBack}
>
	<Header />
	<div
		class="mx-auto flex h-full max-w-7xl flex-col items-center font-medium caret-sky-800 selection:bg-sky-500/20 dark:caret-zinc-200"
	>
		<div id="main" class="flex w-full flex-1 flex-row py-4">
			{@render children?.()}
		</div>
	</div>
	<Footer />
	<Notification />
</main>
