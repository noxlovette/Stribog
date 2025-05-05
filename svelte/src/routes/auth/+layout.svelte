<script lang="ts">
	import { H2 } from '$lib/components';
	const { children } = $props();
</script>

<div class="flex flex-col items-center justify-center">
	<div class="size-72 overflow-hidden">
		<img
			src="/pictures/auth-logo-grind.png"
			alt="Forge and Sparkles"
			class="h-full w-full object-cover"
		/>
	</div>
	<H2>Account</H2>

	{@render children?.()}
</div>
