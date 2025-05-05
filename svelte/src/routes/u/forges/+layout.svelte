<script lang="ts">
    import {ForgeCard, H1} from "$lib/components"
	import { Anvil } from "lucide-svelte";
    const {data, children} = $props();
    
    </script>
    <div class="grid grid-cols-4 ">
        {#if data.forges && data.forges.length > 0}
          <div class="p-5 space-y-5">
            <h1>Forges</h1>
            <div class="grid grid-cols-1 gap-5 min-w-max">
              {#each data.forges as forge}
                <ForgeCard {forge} />
              {/each}
            </div>
          </div>
          {@render children()}
        {:else}
          <div class="col-span-4 flex flex-col justify-center items-center my-4">
            <div class="text-center mb-10 items-center justify-center flex flex-col">
                <Anvil class="size-32 my-2 text-neutral-500" />
              <h2 class="text-2xl font-bold text-gray-900 mb-2">No Forges Yet</h2>
              <p class="text-gray-600 mb-8 max-w-md mx-auto">Create your first forge to start building and managing your content distribution network.</p>
              
              <a href="/u/forges/new" class="inline-flex items-center px-8 py-4 bg-indigo-600 border border-transparent rounded-lg text-lg font-medium text-white shadow-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transform transition-all duration-200 hover:scale-105">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                Create Your First Forge
              </a>
            </div>
          </div>
        {/if}
      </div>