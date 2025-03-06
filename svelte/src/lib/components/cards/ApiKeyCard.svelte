<script lang="ts">
    import { Button } from '$lib/components';
    import { formatDate } from '@noxlovette/svarog';
	import { Activity, Copy, Delete } from 'lucide-svelte';

    const {apiKey} = $props();
  
    function handleCopy() {
      navigator.clipboard.writeText(apiKey.id);

    }
  
  </script>
  
  <div class="flex flex-col p-4 bg-white dark:bg-neutral-800 rounded-lg shadow-sm border border-neutral-200 dark:border-neutral-700">
    <div class="flex justify-between items-start">
      <div>
        <h3 class="text-lg font-medium text-neutral-900 dark:text-white">{apiKey.title}</h3>
        <p class="text-sm text-neutral-500 dark:text-neutral-400 font-mono truncate">{apiKey.id}</p>
      </div>
      <div class="flex space-x-2">
        <Button Icon={Copy} variant="ghost" size="sm" onclick={handleCopy}>
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg>
          Copy
        </Button>
      </div>
    </div>
    <div class="mt-2 flex items-center">
      <span class={`inline-block w-2 h-2 rounded-full ${apiKey.is_active ? 'bg-green-500' : 'bg-red-500'} mr-2`}></span>
      <span class="text-sm text-neutral-600 dark:text-neutral-300">{apiKey.is_active ? 'Active' : 'Inactive'}</span>
    </div>
    <div class="mt-2 text-xs text-neutral-500 dark:text-neutral-400">
      <div>Created: {formatDate(apiKey.created_at)}</div>
      {#if apiKey.last_used_at}
        <div>Last used: {formatDate(apiKey.last_used_at)}</div>
      {:else}
        <div>Never used</div>
      {/if}
    </div>
    <div class="mt-4 flex justify-between">
      <Button Icon={Activity} variant="outline" size="sm" >
        {apiKey.is_active ? 'Deactivate' : 'Activate'}
      </Button>
      <Button Icon={Delete} variant="danger" size="sm" >Delete</Button>
    </div>
  </div>