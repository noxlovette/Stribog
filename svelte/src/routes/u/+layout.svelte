<script lang="ts">
	import { setUser } from '$lib/stores/user.js';
	import { onMount } from 'svelte';

	const { children, data } = $props();

	onMount(() => {
		setUser(data.user);
	});
</script>

{@render children?.()}
