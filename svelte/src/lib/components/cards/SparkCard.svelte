<script lang="ts">

    import { Button } from '$lib/components';
    import { formatDate } from '@noxlovette/svarog';
	import type { Spark } from '$lib/types';
	import { Eye, Pencil, Trash, View } from 'lucide-svelte';
	import { page } from '$app/state';
  
    const {spark}:{spark: Spark} = $props()
    
    const preview = $derived(spark.markdown.length > 150 
      ? spark.markdown.substring(0, 150) + '...' 
      : spark.markdown)
  </script>
  
  <div class="flex flex-col p-5 bg-white dark:bg-neutral-800 rounded-sm shadow-sm border border-neutral-200 dark:border-neutral-700">
    <div class="flex justify-between items-start">
      <h3 class="text-xl font-medium text-neutral-900 dark:text-white">{spark.title}</h3>
      <div class="flex space-x-2">
        <Button Icon={Pencil} variant="ghost" size="sm" href="/u/forges/{page.params.forgeId}/{spark.id}/edit" >
          Edit
        </Button>
        <Button Icon={Eye} href="/u/forges/{page.params.forgeId}/{spark.id}" variant="outline" size="sm">
          View
        </Button>
      </div>
    </div>
  
    <div class="text-sm text-neutral-500 dark:text-neutral-400">
      Last updated: {formatDate(spark.updatedAt)}
    </div>
  
    <div class="mt-4 prose dark:prose-invert prose-sm max-w-none text-neutral-600 dark:text-neutral-300">
      <p>{preview}</p>
    </div>
  
    
  </div>