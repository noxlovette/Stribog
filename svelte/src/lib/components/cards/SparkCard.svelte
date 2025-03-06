<script lang="ts">

    import { Button } from '$lib/components';
    import { formatDate } from '@noxlovette/svarog';
	import type { Spark } from '$lib/types';
	import { Pencil, Trash, View } from 'lucide-svelte';
	import { page } from '$app/state';
  
    const {spark}:{spark: Spark} = $props()
    
    const preview = $derived(spark.markdown.length > 150 
      ? spark.markdown.substring(0, 150) + '...' 
      : spark.markdown)
  </script>
  
  <div class="flex flex-col p-5 bg-white dark:bg-neutral-800 rounded-lg shadow-sm border border-neutral-200 dark:border-neutral-700">
    <div class="flex justify-between items-start">
      <h3 class="text-xl font-medium text-neutral-900 dark:text-white">{spark.title}</h3>
      <div class="flex space-x-2">
        <Button Icon={Pencil} variant="ghost" size="sm" >
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/></svg>
          Edit
        </Button>
        <Button Icon={Trash} variant="danger" size="sm" >
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
          Delete
        </Button>
      </div>
    </div>
  
    <div class="mt-3 text-sm text-neutral-500 dark:text-neutral-400">
      Last updated: {formatDate(spark.updatedAt)}
    </div>
  
    <div class="mt-4 prose dark:prose-invert prose-sm max-w-none text-neutral-600 dark:text-neutral-300">
      <p>{preview}</p>
    </div>
  
    <Button Icon={View} href="/u/forges/{page.params.forgeId}/{spark.id}" variant="outline" size="sm">
      View Details
    </Button>
  </div>