<script lang="ts">
	import { formatDate } from '@noxlovette/svarog';

	const { data } = $props();

	const spark = $derived(data.spark);
</script>

<div class="flex items-start justify-between">
	<h3 class="text-xl font-medium">{spark.title}</h3>
</div>

<div class="text-sm text-stone-500 dark:text-stone-400">
	Last updated: {formatDate(spark.updatedAt)}
</div>

<div class="prose dark:prose-invert prose-sm mt-4 max-w-none text-stone-600 dark:text-stone-300">
	<p>{@html data.rendered}</p>
</div>
