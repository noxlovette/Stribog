<script lang="ts">
	let { children, styling = '' } = $props();
</script>

<h1
	class="{styling} text-frosted my-0.5 text-center text-xl font-light tracking-tight text-balance md:text-left md:text-2xl"
>
	{@render children?.()}
</h1>
