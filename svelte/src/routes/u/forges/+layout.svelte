<script lang="ts">
    import {ForgeCard} from "$lib/components"
    const {data, children} = $props();
    
    </script>
    <div class="w-full flex justify-between">
        {#if data.forges}
        <div
            class="grid grid-cols-1 gap-5 p-5 border-x-2 border-dashed min-w-max md:w-1/5 bg-stripes"
        >
            {#each data.forges as forge}
                <ForgeCard {forge} />
            {/each}
        </div>
        {/if}
    {@render children()}
    </div>