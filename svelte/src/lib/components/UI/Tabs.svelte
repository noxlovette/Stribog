<script lang="ts">
	const { children } = $props();
</script>

<nav class="col-span-2 grid grid-cols-4 items-start gap-3">
	{@render children?.()}
</nav>
