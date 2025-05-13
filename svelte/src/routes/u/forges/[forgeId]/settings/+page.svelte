<script lang="ts">
	import { Button, Input } from '$lib/components';
	import {enhance} from "$app/forms"
	import { enhanceForm } from '$lib/utils';

	const { data } = $props();
	const forge = $derived(data.forge);
</script>

<form method="POST" use:enhance={enhanceForm({})}
	>
	<Input name="title" placeholder="Title" value={forge.title}></Input>
	<Input name="description" placeholder="Description" value={forge.description}></Input>
	
	<Button type="submit" variant="primary">Submit</Button>
</form>
