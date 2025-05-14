<script lang="ts">
	import { enhance } from '$app/forms';
	import { Plus } from 'lucide-svelte';
	import { enhanceForm } from '$lib/utils';
	import { notification } from '$lib/stores';
</script>

<form
	action="?/create"
	method="POST"
	class="group flex min-h-36 flex-col items-center justify-center rounded-sm border border-stone-800/60 px-3 py-2 transition-colors hover:border-orange-800"
	use:enhance={enhanceForm({
		messages: {
			redirect: 'New Entity Created'
		},
		navigate: true
	})}
>
	<button type="submit" class="flex size-full items-center justify-center">
		<Plus
			class="size-24 text-stone-600 transition-colors group-hover:text-inherit"
			strokeWidth={1}
		/>
	</button>
</form>
