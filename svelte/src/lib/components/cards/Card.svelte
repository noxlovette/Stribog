<script lang="ts">
	const { children } = $props();
</script>

<div class="flex flex-col rounded-md bg-stone-900/60 p-5">
	{@render children()}
</div>
