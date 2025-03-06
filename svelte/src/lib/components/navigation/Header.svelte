<script>
	import { page } from '$app/state';
	import { notification, user } from '$lib/stores';
</script>

<header class="bg-gray-800 text-white shadow-lg">
	<div class="container mx-auto px-4 py-3">
		<div class="flex items-center justify-between">
			<!-- Logo -->
			<div class="flex items-center space-x-2">
				<a
					href="/"
					class="font-ponomar text-2xl font-bold tracking-tight transition hover:text-blue-400"
					>Стрибог</a
				>
			</div>

			<!-- Desktop Navigation -->
			<nav class="hidden items-center space-x-6 md:flex">
				<a
					href="/"
					class="transition hover:text-blue-400 {page.url.pathname === '/' ? 'text-blue-400' : ''}"
					>В Избу</a
				>
				<a
					href="/forges"
					class="transition hover:text-blue-400 {page.url.pathname.startsWith('/forges')
						? 'text-blue-400'
						: ''}">Кузни</a
				>

				{#if $user && $user.username}
					<!-- User is logged in -->
					<div class="group relative">
						<button class="flex items-center space-x-1 transition hover:text-blue-400">
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
							class="invisible absolute right-0 z-50 mt-2 w-48 rounded bg-gray-700 py-1 opacity-0 shadow-lg transition-all duration-200 group-hover:visible group-hover:opacity-100"
						>
							<a href="/account" class="block px-4 py-2 text-sm hover:bg-gray-600">Установки</a>
							<a href="/forges/new" class="block px-4 py-2 text-sm hover:bg-gray-600"
								>Открыть новую Кузню</a
							>
							<button
								onclick={() => notification.set({ message: 'not implemented', type: 'info' })}
								class="block w-full px-4 py-2 text-left text-sm text-red-400 hover:bg-gray-600"
							>
								Покинуть Деревню
							</button>
						</div>
					</div>
				{:else}
					<!-- User is not logged in -->
					<a
						href="/login"
						class="transition hover:text-blue-400 {page.url.pathname === '/login'
							? 'text-blue-400'
							: ''}">Login</a
					>
					<a
						href="/signup"
						class="rounded bg-blue-600 px-4 py-2 transition hover:bg-blue-700 {page.url.pathname ===
						'/signup'
							? 'bg-blue-700'
							: ''}"
					>
						Sign Up
					</a>
				{/if}
			</nav>
		</div>
	</div>
</header>
