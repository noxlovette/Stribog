<script lang="ts">
    const {children} = $props()
</script>

<div class="items-center justify-center flex size-full">
{@render children?.()}
</div>