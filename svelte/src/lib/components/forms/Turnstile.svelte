<!-- Turnstile.svelte -->
<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { page } from '$app/state';

	let turnstileElement: HTMLDivElement;
	let widgetId: string | null = null;

	function renderWidget() {
		if (typeof window === 'undefined' || !window.turnstile || !turnstileElement) return;

		// Clean up previous widget if it exists
		if (widgetId) {
			try {
				window.turnstile.remove(widgetId);
			} catch (e) {
				console.error('Error removing Turnstile widget:', e);
			}
		}

		// Render new widget
		widgetId = window.turnstile.render(turnstileElement, {
			sitekey: '0x4AAAAAAA_p3isIo2-ed2pU',
			theme: 'auto'
		});
	}

	// Re-render on route changes
	$effect(() => {
		if (page.url.pathname) {
			setTimeout(() => renderWidget(), 100);
		}
	});

	onMount(() => {
		renderWidget();
	});

	onDestroy(() => {
		if (widgetId && window.turnstile) {
			window.turnstile.remove(widgetId);
		}
	});
</script>

<div bind:this={turnstileElement} class="cf-turnstile my-4"></div>
