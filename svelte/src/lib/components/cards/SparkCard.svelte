<script lang="ts">
	import { Button } from '$lib/components';
	import { formatDate } from '@noxlovette/svarog';
	import type { Spark } from '$lib/types';
	import { page } from '$app/state';

	const { spark }: { spark: Spark } = $props();

	const preview = $derived(
		spark.markdown.length > 150 ? spark.markdown.substring(0, 150) + '...' : spark.markdown
	);
</script>

<div class="flex flex-col rounded-sm border border-stone-600/40 bg-stone-900 p-5 shadow-sm">
	<div class="flex items-start justify-between">
		<h3 class="text-xl font-medium">{spark.title}</h3>
		<div class="flex space-x-2">
			<Button variant="ghost" size="sm" href="/u/forges/{page.params.forgeID}/{spark.id}/edit">
				Edit
			</Button>
			<Button href="/u/forges/{page.params.forgeID}/{spark.id}" variant="outline" size="sm">
				View
			</Button>
		</div>
	</div>

	<div class="text-sm text-stone-500 dark:text-stone-400">
		Last updated: {formatDate(spark.updatedAt)}
	</div>

	<div class="prose dark:prose-invert prose-sm mt-4 max-w-none text-stone-600 dark:text-stone-300">
		<p>{preview}</p>
	</div>
</div>
