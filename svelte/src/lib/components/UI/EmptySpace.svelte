<script lang="ts">
	let { children } = $props();
</script>

<div
	class="flex min-h-40 w-full flex-col items-center justify-center rounded-lg border-2 border-dashed border-stone-300/30 p-4 text-center dark:border-stone-700"
>
	{@render children?.()}
</div>
