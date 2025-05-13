<script lang="ts">
	const { data, children } = $props();
</script>

<div class="flex flex-col">
	{@render children()}
</div>
