<script lang="ts">
	import { Button } from '$lib/components';
	import { formatDate } from '@noxlovette/svarog';
	import { page } from '$app/state';

	const { data } = $props();

	const spark = $derived(data.spark);
</script>

<div class="flex flex-col rounded-sm border border-stone-600/40 bg-stone-900 p-5 shadow-sm">
	<div class="flex items-start justify-between">
		<h3 class="text-xl font-medium">{spark.title}</h3>
	</div>

	<div class="text-sm text-stone-500 dark:text-stone-400">
		Last updated: {formatDate(spark.updatedAt)}
	</div>

	<div class="prose dark:prose-invert prose-sm mt-4 max-w-none text-stone-600 dark:text-stone-300">
		<p>{spark.markdown}</p>
	</div>
</div>
