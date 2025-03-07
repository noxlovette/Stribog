<script lang="ts">
	import { enhance } from "$app/forms";
    import { H1, SparkCard, Button, H2, H3, CollaboratorCard, ApiKeyCard } from "$lib/components";
    import { enhanceForm } from '@noxlovette/svarog';
	import { Ban, Key, Newspaper, Save, Sparkle, StopCircle, UserPlus } from "lucide-svelte";
  import {invalidate} from "$app/navigation"
	import { fade } from "svelte/transition";
	import Input from "$lib/components/forms/Input.svelte";
	import { isLoading, notification } from "$lib/stores";

  
    const { data } = $props();

    let apiKeyCreationDialogue = $state(false);

  </script>
  
    <!-- Sparks Section -->
    <div class="col-span-1 lg:col-span-2 flex flex-col space-y-5 p-4">
      <div class="flex justify-between items-center">
        <H1>Sparks</H1>
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
        <div class="flex flex-col items-center justify-center p-8 bg-neutral-50 dark:bg-neutral-800/50 rounded-sm border border-dashed border-neutral-200 dark:border-neutral-700">
          <p class="text-lg font-medium text-neutral-600 dark:text-neutral-300">No sparks yet</p>
          <p class="text-sm text-neutral-500 dark:text-neutral-400">Create your first spark to get started</p>
        </div>
      {/if}
    </div>
  
    <!-- Sidebar with API Keys and Collaborators -->
    <div class="col-span-1 flex flex-col space-y-6">
      <!-- API Keys Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-sm">
        <div class="flex flex-col">
          <div class="flex justify-between">
          <H2>API Keys</H2>
          <Button variant="outline" type="button" onclick={()=> apiKeyCreationDialogue = true} size="sm" Icon={Key}>
            New Key
          </Button>
        </div>
          {#if apiKeyCreationDialogue}
          <form method="POST" class="" action="?/newKey" use:enhance={enhanceForm({
            messages: {
              success: "Created New Key",
              failure: "Failed to create",
              defaultError: "Failed to create key",
            },
            handlers: {
              success: async () => {
                invalidate("forge:general");
              },
            },
            notificationStore: notification,
            isLoadingStore: isLoading
          })}>
          <div in:fade class="space-y-3">
            <H3>
              Name the Key
            </H3>
            <Input name="title" placeholder="Name the Key" value="" />
            <div class="flex space-x-2">
            <Button variant = "primary" Icon={Save} type="submit">Create</Button>
            <Button variant="ghost" Icon={Ban} onclick={()=> apiKeyCreationDialogue = false}> Cancel </Button>
          </div>
          </div>
        </form>
        {/if}
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
          <div class="flex flex-col items-center justify-center p-4 bg-white dark:bg-neutral-800 rounded-sm border border-dashed border-neutral-200 dark:border-neutral-700">
            <p class="text-sm text-neutral-500 dark:text-neutral-400">No API keys yet</p>
          </div>
        {/if}
      </div>
  
      <!-- Collaborators Section -->
      <div class="flex flex-col space-y-4 p-4 bg-neutral-50 dark:bg-neutral-800/30 rounded-sm">
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
          <div class="flex flex-col items-center justify-center p-4 bg-white dark:bg-neutral-800 rounded-sm border border-dashed border-neutral-200 dark:border-neutral-700">
            <p class="text-sm text-neutral-500 dark:text-neutral-400">No collaborators yet</p>
          </div>
        {/if}
      </div>
    </div>
