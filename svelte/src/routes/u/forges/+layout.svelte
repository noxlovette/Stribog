<script lang="ts">
    import {ForgeCard, H1} from "$lib/components"
    const {data, children} = $props();
    
    </script>
    <div class="grid grid-cols-4">
        <div class="p-5 space-y-5">
        <H1>
            Forges
        </H1>
        {#if data.forges}
        <div
            class="grid grid-cols-1 gap-5 min-w-max"
        >
            {#each data.forges as forge}
                <ForgeCard {forge} />
            {/each}
        </div>
        {/if}
    </div>
    {@render children()}
    </div>