<script lang="ts">
	import { ForgeCard, ForgeColumn, H1 } from '$lib/components';
	import { Anvil } from 'lucide-svelte';
	const { data, children } = $props();
</script>

<div class="flex flex-col">
	{@render children()}
</div>
