<script lang="ts">
	import { Plus } from 'lucide-svelte';
	const { href = '/' } = $props();
</script>

<a
	{href}
	class="group flex min-h-36 flex-col items-center justify-center rounded-sm px-3 py-2 ring ring-stone-600 transition-colors hover:bg-orange-600/20 hover:ring-orange-800"
>
	<Plus class="size-24 text-stone-600 transition-colors group-hover:text-inherit" strokeWidth={1} />
</a>
