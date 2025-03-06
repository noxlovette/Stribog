<script lang="ts">
    import {ForgeCard, H1} from "$lib/components"
    const {data, children} = $props();
    
    </script>
    <div class="w-full flex flex-row">
        <div class="p-5 space-y-5">
        <H1>
            Кузни
        </H1>
        {#if data.forges}
        <div
            class="grid grid-cols-1 gap-5 min-w-max bg-stripes"
        >
            {#each data.forges as forge}
                <ForgeCard {forge} />
            {/each}
        </div>
        {/if}
    </div>
    {@render children()}
    </div>