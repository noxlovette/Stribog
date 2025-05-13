<script lang="ts">
	import { Button, Input } from '$lib/components';
	import {enhance} from "$app/forms"
	import { enhanceForm } from '$lib/utils';

	const { data } = $props();
	const spark = $derived(data.spark);
</script>

<form method="POST" use:enhance={enhanceForm({})}
	>
	<Input name="title" placeholder="Title" value={spark.title}></Input>
	<Input name="markdown" placeholder="Markdown" value={spark.markdown}></Input>
	{#each spark?.tags ?? [] as tag}
		<Input name="tag" placeholder="Tag" value={tag}></Input>
	{/each}
	<Button type="submit" variant="primary">Submit</Button>
</form>
