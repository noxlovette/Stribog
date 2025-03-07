<script lang="ts">
    import { H1, SparkCard, Button, H2 } from "$lib/components";
    import {CollaboratorCard, ApiKeyCard} from "$lib/components";
    import { formatDate } from '@noxlovette/svarog';
	import { Key, Newspaper, Sparkle, UserPlus } from "lucide-svelte";
  
    const { data } = $props();

  </script>
  
    <!-- Sparks Section -->
    <div class="col-span-1 lg:col-span-2 flex flex-col space-y-5 p-4">
      <div class="flex justify-between items-center">
        <H1>Искры</H1>
        <Button Icon={Sparkle} variant="primary">
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
        <div class="flex flex-col space-y-2 items-center justify-center p-8 bg-neutral-50 dark:bg-neutral-800/50 rounded-lg border border-dashed border-neutral-200 dark:border-neutral-700">
          <p class="text-lg font-medium text-neutral-600 dark:text-neutral-300">No sparks yet</p>
          <p class="text-sm text-neutral-500 dark:text-neutral-400">Create your first spark to get started</p>
          <Button Icon={Sparkle} variant="primary">Create Spark</Button>
        </div>
      {/if}
    </div>
  
    <!-- Sidebar with API Keys and Collaborators -->
    <div class="col-span-1 flex flex-col space-y-6">
      <!-- API Keys Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-lg">
        <div class="flex justify-between items-center">
          <H2>API Keys</H2>
          <Button variant="outline" size="sm" Icon={Key}>
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
            <p class="mt-2 text-sm text-neutral-500 dark:text-neutral-400">No API keys yet</p>
          </div>
        {/if}
      </div>
  
      <!-- Collaborators Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-lg">
        <div class="flex justify-between items-center">
          <H2>Collaborators</H2>
          <Button variant="outline" size="sm" Icon={UserPlus}>
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
            <p class="mt-2 text-sm text-neutral-500 dark:text-neutral-400">No collaborators yet</p>
            <Button variant="outline" size="sm" Icon={UserPlus}>Add Collaborator</Button>
          </div>
        {/if}
      </div>
    </div>
