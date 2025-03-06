<script>
	import { page } from '$app/state';
	import { notification, user } from '$lib/stores';
</script>

<header class="my-2 w-11/12 items-baseline rounded-lg shadow-md ring ring-zinc-200 md:w-full dark:ring-zinc-900 ">
	<div class="container mx-auto px-4 py-3 max-w-7xl">
		<div class="flex items-center justify-between">
			<!-- Logo -->
			<div class="flex items-center space-x-2">
				<a
					href="/"
					class="font-ponomar text-2xl font-bold tracking-tight transition hover:text-indigo-400"
					>Стрибог</a
				>
			</div>

			<!-- Desktop Navigation -->
			<nav class="hidden items-center space-x-6 md:flex">
				<a
					href="/u/dashboard"
					class="transition hover:text-indigo-400 {page.url.pathname === '/u/dashboard' ? 'text-indigo-400' : ''}"
					>В Избу</a
				>
				<a
					href="/u/forges"
					class="transition hover:text-indigo-400 {page.url.pathname.startsWith('/u/forges')
						? 'text-indigo-400'
						: ''}">Кузни</a
				>

				{#if $user && $user.username}
					<!-- User is logged in -->
					<div class="group relative">
						<button class="flex items-center space-x-1 transition hover:text-indigo-400">
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
							class="invisible absolute right-0 z-50 mt-2 w-48 rounded bg-neutral-700 py-1 opacity-0 shadow-lg transition-all duration-200 group-hover:visible group-hover:opacity-100"
						>
							<a href="/account" class="block px-4 py-2 text-sm hover:bg-neutral-600">Установки</a>
							<a href="/forges/new" class="block px-4 py-2 text-sm hover:bg-neutral-600"
								>Открыть новую Кузню</a
							>
							<button
								onclick={() => notification.set({ message: 'not implemented', type: 'info' })}
								class="block w-full px-4 py-2 text-left text-sm text-red-400 hover:bg-neutral-600"
							>
								Покинуть Деревню
							</button>
						</div>
					</div>
				{:else}
					<!-- User is not logged in -->
					<a
						href="/login"
						class="transition hover:text-indigo-400 {page.url.pathname === '/login'
							? 'text-indigo-400'
							: ''}">Login</a
					>
					<a
						href="/signup"
						class="rounded bg-indigo-600 px-4 py-2 transition hover:bg-indigo-700 {page.url.pathname ===
						'/signup'
							? 'bg-indigo-700'
							: ''}"
					>
						Sign Up
					</a>
				{/if}
			</nav>
		</div>
	</div>
</header>
