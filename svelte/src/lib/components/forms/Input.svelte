<script lang="ts">
	let {
		placeholder = '',
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
		'mt-2 w-full rounded-full p-3 font-medium border-stone-700 placeholder:text-stone-700 bg-stone-950';
</script>

{#if type === 'text'}
	<input
		{name}
		type="text"
		bind:value
		class="{styling} focus:border-orange-500 focus:ring-orange-500"
		{placeholder}
	/>
{:else if type === 'textarea'}
	<textarea
		{name}
		rows="5"
		bind:value
		class="{styling} focus:border-orange-500 focus:ring-orange-500"
		{placeholder}
	></textarea>
{:else if type === 'number'}
	<input
		type="number"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-orange-500 focus:ring-orange-500"
	/>
{:else if type === 'password'}
	<input
		type="password"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-orange-500 focus:ring-orange-500"
	/>
{:else if type === 'email'}
	<input
		type="email"
		{placeholder}
		{name}
		bind:value
		class="{styling} focus:border-orange-500 focus:ring-orange-500"
	/>
{/if}
