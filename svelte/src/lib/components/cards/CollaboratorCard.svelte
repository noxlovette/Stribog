<script lang="ts">
    import { Button } from '$lib/components';
	import { RollerCoaster, Trash } from 'lucide-svelte';

    const {collaborator}  = $props()
 
  </script>
  
  <div class="flex flex-col p-4 bg-white dark:bg-neutral-800 rounded-lg shadow-sm border border-neutral-200 dark:border-neutral-700">
    <div class="flex items-start justify-between">
      <div class="flex items-center">
        <div class="w-10 h-10 bg-neutral-200 dark:bg-neutral-700 rounded-full flex items-center justify-center text-neutral-500 dark:text-neutral-400">
          {collaborator.name.substring(0, 2).toUpperCase()}
        </div>
        <div class="ml-3">
          <h3 class="text-lg font-medium text-neutral-900 dark:text-white">{collaborator.name}</h3>
          <p class="text-sm text-neutral-500 dark:text-neutral-400">@{collaborator.username}</p>
        </div>
      </div>
      <div>
        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-800 dark:text-blue-100">
          {collaborator.role}
        </span>
      </div>
    </div>
    <div class="mt-2 text-sm text-neutral-600 dark:text-neutral-300">
      <p>{collaborator.email}</p>
    </div>
    <div class="mt-4 flex justify-between">
      <Button variant="outline" size="sm" Icon={RollerCoaster}>
        Change Role
      </Button>
      <Button variant="danger" size="sm" Icon={Trash} >Remove</Button>
    </div>
  </div>