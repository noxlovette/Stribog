<script lang="ts">
	import { ApiKeyCard, AddCard, CollaboratorCard } from '$lib/components';

	const { data } = $props();
</script>

<div class="grid grid-cols-3 gap-3">
	<AddCard></AddCard>
	{#each data.collaborators as collaborator}
		<CollaboratorCard {collaborator}></CollaboratorCard>
	{/each}
</div>
