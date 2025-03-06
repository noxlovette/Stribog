<script lang="ts">
	import type { Forge } from "$lib/types";

const {forge}: {forge: Forge} = $props();
</script>

<div>
    {forge.title}
</div>