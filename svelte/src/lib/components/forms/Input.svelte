<script lang="ts">
	import Label from '$lib/components/typography/Label.svelte';

	let {
		placeholder = 'Edit here',
		name = 'name',
		labelName = name,
		value = $bindable(),
		type = 'text'
	}: {
		placeholder: string;
		name: string;
		value: string | number | boolean | undefined;
		labelName?: string;
		type?: 'text' | 'number' | 'textarea' | 'password' | 'email' | 'checkbox';
	} = $props();

	const styling =
		'mt-2 w-full rounded-md p-3 text-lg shadow-sm dark:border-slate-700 dark:bg-slate-950';
</script>

<Label>{labelName}</Label>
{#if type === 'text'}
	<input
		{name}
		type="text"
		bind:value
		class="{styling} focus:border-sky-500 focus:ring-sky-500"
		{placeholder}
	/>
{:else if type === 'textarea'}
	<textarea
		{name}
		rows="5"
		bind:value
		class="{styling} focus:border-sky-500 focus:ring-sky-500"
		{placeholder}
	></textarea>
{:else if type === 'number'}
	<input
		type="number"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-sky-500 focus:ring-sky-500"
	/>
{:else if type === 'password'}
	<input
		type="password"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-sky-500 focus:ring-sky-500"
	/>
{:else if type === 'email'}
	<input
		type="email"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-sky-500 focus:ring-sky-500"
	/>
{/if}
