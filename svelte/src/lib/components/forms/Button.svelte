<script lang="ts">
	import { Component, Loader2 } from 'lucide-svelte';
	import { isLoading } from '$lib/stores';
	import type { ComponentType, Snippet } from 'svelte';
	import type { MouseEventHandler } from 'svelte/elements';
	import ModalBackground from '../UI/ModalBackground.svelte';

	type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'ghost' | 'link' | 'outline';

	type ButtonSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl';
	type ButtonType = 'button' | 'submit' | 'reset' | undefined;

	interface Props {
		variant: ButtonVariant;
		size?: ButtonSize;
		type?: ButtonType;
		href?: string | undefined;
		formaction?: string | undefined;
		styling?: string;
		disable?: boolean;
		Icon?: Component;
		iconPosition?: 'left' | 'right';
		fullWidth?: boolean;
		rounded?: boolean;
		confirmText?: string | undefined;
		confirmTitle?: string | undefined;
		onclick?: MouseEventHandler<HTMLButtonElement> | undefined;
		children?: Snippet;
	}

	let {
		variant = 'primary',
		size = 'md',
		type = 'button',
		href = undefined,
		formaction = undefined,
		styling = '',
		disable = false,
		Icon = undefined,
		iconPosition = 'left',
		fullWidth = false,
		rounded = false,
		confirmText = undefined,
		confirmTitle = undefined,
		onclick = undefined,
		children
	}: Props = $props();

	const isLink = $derived(!!href);
	let disabled = $derived($isLoading || disable);
	let showConfirmDialog = $state(false);

	function handleClick(event: any) {
		if (variant === 'danger' && (confirmText || confirmTitle)) {
			event.preventDefault();
			showConfirmDialog = true;
		}
	}

	const sizeClasses = {
		xs: 'px-2 py-1 text-tiny',
		sm: 'px-3 py-1.5 text-sm',
		md: 'px-3.5 py-1.5 text-[15px]',
		lg: 'px-5 py-2.5 text-base md:px-6',
		xl: 'px-6 py-3 text-lg md:px-8'
	};

	const baseClasses =
		'flex items-center justify-center rounded-full transition-all duration-150 ease-in-out font-medium select-none focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-orange-400 disabled:cursor-not-allowed disabled:opacity-50 backdrop-blur-md';

	const variantClasses = {
		primary:
			'bg-gradient-to-br ring-1 ring-orange-700 transition-colors shadow-sm ring-1 from-orange-900/70 to-orange-800/70 text-orange-50 hover:from-orange-800 hover:to-orange-700',
		secondary: 'ring-1 shadow-sm bg-stone-900/80 text-stone-200 hover:bg-stone-800 ring-stone-700',
		danger:
			'bg-gradient-to-br transition-colors ring-1 ring-red-600/50 shadow-sm from-red-700/70 to-red-800/70 text-white hover:from-orange-600 hover:to-orange-700',
		ghost: 'text-stone-400 hover:bg-stone-800/60',
		link: 'text-orange-600 underline hover:text-orange-800 p-0 ring-0 dark:text-orange-300 dark:hover:text-orange-100',
		outline: 'bg-transparent  ring-1 text-stone-200 ring-stone-600/50 hover:bg-stone-800'
	};

	// const shapeClasses = $derived(rounded ? 'rounded-full' : 'rounded-lg');
	const widthClasses = $derived(fullWidth ? 'w-full' : '');

	const allClasses = $derived(
		[
			baseClasses,
			variantClasses[variant],
			sizeClasses[size],
			// shapeClasses,
			widthClasses,
			styling
		].join(' ')
	);
</script>

{#if isLink}
	<a {href} class={allClasses} aria-disabled={disabled}>
		{#if $isLoading}
			<Loader2 class="mr-2 size-4 animate-spin" />
		{:else if Icon && iconPosition === 'left'}
			<Icon class="mr-2 size-4" />
		{/if}

		{#if $isLoading}Loading...
		{:else}
			{@render children?.()}
		{/if}

		{#if Icon && iconPosition === 'right'}
			<Icon class="ml-2 size-4" />
		{/if}
	</a>
{:else}
	<button
		{type}
		{formaction}
		{disabled}
		class={allClasses}
		onclick={variant === 'danger' ? handleClick : onclick}
	>
		{#if $isLoading}
			<Loader2 class="mr-2 size-4 animate-spin" />
		{:else if Icon && iconPosition === 'left'}
			<Icon class="mr-2 size-4" />
		{/if}

		{#if $isLoading}Loading...
		{:else}
			{@render children?.()}
		{/if}

		{#if Icon && iconPosition === 'right'}
			<Icon class="ml-2 size-4" />
		{/if}
	</button>
{/if}

{#if showConfirmDialog}
	<ModalBackground>
		<h3 class="text-lg font-semibold text-stone-800 dark:text-stone-200">
			{confirmTitle || 'Are you sure?'}
		</h3>
		<p class="mt-2 text-sm text-stone-600 dark:text-stone-400">
			{confirmText || 'This action cannot be undone.'}
		</p>
		<div class="mt-5 flex justify-end gap-2">
			<button
				type="button"
				class="rounded-lg bg-white px-4 py-2 text-stone-700/30 ring-1 ring-stone-300 transition-all hover:bg-stone-50 dark:bg-stone-800/30 dark:text-stone-200 dark:hover:bg-stone-700"
				onclick={() => (showConfirmDialog = false)}
			>
				Cancel
			</button>
			<button
				type="submit"
				class="rounded-lg bg-gradient-to-br from-orange-500 to-orange-600 px-4 py-2 text-white ring-1 ring-orange-300 transition-all hover:from-orange-500 hover:to-orange-700 focus:ring focus:ring-orange-400 focus:ring-offset-2"
				{formaction}
			>
				Confirm
			</button>
		</div>
	</ModalBackground>
{/if}
