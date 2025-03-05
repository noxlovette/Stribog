<script lang="ts">
	let { children } = $props();
</script>

<h2 class="text-md snug font-medium text-balance text-slate-500 sm:text-lg lg:text-xl">
	{@render children()}
</h2>
