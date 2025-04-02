<script>
	import { page } from '$app/state';
	import { initialUser, notification, user } from '$lib/stores';
	import { Anvil, Home } from 'lucide-svelte';
	import { Button } from '../forms';
</script>

<header class="my-1 w-11/12 items-baseline md:w-full py-2 px-4 border-b border-zinc-200">

		<div class="flex items-center justify-between">
			<!-- Logo -->
			<div class="flex items-center space-x-2 text-teal-600">
				<a
					href="/"
					class="text-2xl font-teko font-bold tracking-tight transition hover:text-teal-400"
					>Stribog</a
				>

				
			</div>

			<!-- Desktop Navigation -->
			<nav class="hidden items-center space-x-3 md:flex">

				<Button Icon={Home} variant="ghost" href="/u/dashboard"
				styling={page.url.pathname === '/u/dashboard' ? 'bg-teal-50' : ''}
				>Izba</Button>
				<Button Icon={Anvil} variant="ghost" href="/u/forges"
				styling={page.url.pathname === '/u/forges' ? 'bg-teal-50' : ''}
				>Forges</Button>
				
				{#if $user && $user.username}
					<!-- User is logged in -->
					<div class="group relative">
						<button class="flex items-center space-x-1 transition hover:text-teal-400">
							<span>{$user.username}</span>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="h-4 w-4"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M19 9l-7 7-7-7"
								/>
							</svg>
						</button>
						<div
							class="invisible absolute right-0 z-50 mt-2 w-48 rounded bg-neutral-50 py-1 opacity-0 shadow-lg transition-all duration-200 group-hover:visible group-hover:opacity-100"
						>
							<a href="/account" class="block px-4 py-2 text-sm hover:bg-neutral-100">Settings</a>
							<a href="/forges/new" class="block px-4 py-2 text-sm hover:bg-neutral-100"
								>Open New Forge</a
							>
							<button
								onclick={() => {
									fetch("/auth/logout", {method: "POST"});
									user.set(initialUser);
									localStorage.clear()
									}}
								class="block w-full px-4 py-2 text-left text-sm text-red-400 hover:bg-neutral-100"
							>
								Leave the Village
							</button>
						</div>
					</div>
				{:else}
					<!-- User is not logged in -->
					<a
						href="/auth/login"
						class="transition hover:text-teal-400 {page.url.pathname === '/auth/login'
							? 'text-teal-400'
							: ''}">Login</a
					>
					<a
						href="/auth/signup"
						class="rounded bg-teal-600 px-4 py-2 transition hover:bg-teal-700 {page.url.pathname ===
						'/auth/signup'
							? 'bg-teal-700'
							: ''}"
					>
						Sign Up
					</a>
				{/if}
			</nav>
		</div>

</header>
