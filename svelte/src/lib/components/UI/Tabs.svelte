<script lang="ts">
	const { children } = $props();
</script>

<nav class="grid grid-cols-4 items-start gap-3">
	{@render children?.()}
</nav>
