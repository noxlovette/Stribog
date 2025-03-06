<script lang="ts">
    import { H1, SparkCard, Button } from "$lib/components";
    import {CollaboratorCard, ApiKeyCard} from "$lib/components";
    import { formatDate } from '@noxlovette/svarog';
	import { Key, Newspaper, UserPlus } from "lucide-svelte";
  
    const { data } = $props();

  </script>
  
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
    <!-- Sparks Section -->
    <div class="col-span-1 lg:col-span-2 flex flex-col space-y-5 p-4">
      <div class="flex justify-between items-center">
        <H1>Искры</H1>
        <Button Icon={Newspaper} variant="primary">
          <svg xmlns="http://www.w3.org/2000/svg" class="mr-2 h-4 w-4" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
          Create Spark
        </Button>
      </div>
      
      {#if data.sparks && data.sparks.length > 0}
        <div class="grid grid-cols-1 gap-5 min-w-max">
          {#each data.sparks as spark}
            <SparkCard {spark} />
          {/each}
        </div>
      {:else}
        <div class="flex flex-col items-center justify-center p-8 bg-neutral-50 dark:bg-neutral-800/50 rounded-lg border border-dashed border-neutral-200 dark:border-neutral-700">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-neutral-400" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"/></svg>
          <p class="mt-4 text-lg font-medium text-neutral-600 dark:text-neutral-300">No sparks yet</p>
          <p class="mt-1 text-sm text-neutral-500 dark:text-neutral-400">Create your first spark to get started</p>
          <Button Icon={Newspaper} variant="primary">Create Spark</Button>
        </div>
      {/if}
    </div>
  
    <!-- Sidebar with API Keys and Collaborators -->
    <div class="col-span-1 flex flex-col space-y-6">
      <!-- API Keys Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-lg">
        <div class="flex justify-between items-center">
          <H1>API Keys</H1>
          <Button variant="outline" size="sm" Icon={Key}>
            <svg xmlns="http://www.w3.org/2000/svg" class="mr-1 h-4 w-4" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
            New Key
          </Button>
        </div>
        
        {#if data.apiKeys && data.apiKeys.length > 0}
          <div class="grid grid-cols-1 gap-4">
            {#each data.apiKeys as apiKey}
              <ApiKeyCard 
                {apiKey} 
              />
            {/each}
          </div>
        {:else}
          <div class="flex flex-col items-center justify-center p-4 bg-white dark:bg-neutral-800 rounded-lg border border-dashed border-neutral-200 dark:border-neutral-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-neutral-400" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"/></svg>
            <p class="mt-2 text-sm text-neutral-500 dark:text-neutral-400">No API keys yet</p>
            <Button variant="outline" size="sm" Icon={Key}>Create Key</Button>
          </div>
        {/if}
      </div>
  
      <!-- Collaborators Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-lg">
        <div class="flex justify-between items-center">
          <H1>Collaborators</H1>
          <Button variant="outline" size="sm" Icon={UserPlus}>
            <svg xmlns="http://www.w3.org/2000/svg" class="mr-1 h-4 w-4" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
            Add User
          </Button>
        </div>
        
        {#if data.collaborators && data.collaborators.length > 0}
          <div class="grid grid-cols-1 gap-4">
            {#each data.collaborators as collaborator}
              <CollaboratorCard 
                {collaborator} 
              />
            {/each}
          </div>
        {:else}
          <div class="flex flex-col items-center justify-center p-4 bg-white dark:bg-neutral-800 rounded-lg border border-dashed border-neutral-200 dark:border-neutral-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-neutral-400" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="7" r="4"/><path d="M5 22v-4a8 8 0 0 1 16 0v4"/></svg>
            <p class="mt-2 text-sm text-neutral-500 dark:text-neutral-400">No collaborators yet</p>
            <Button variant="outline" size="sm" Icon={UserPlus}>Add Collaborator</Button>
          </div>
        {/if}
      </div>
    </div>
  </div>