<script lang="ts">
	let { children, href = '/' } = $props();
</script>

<a {href} class="p-1 font-sans text-xs text-indigo-400 lg:text-sm">
	{@render children()}
</a>
