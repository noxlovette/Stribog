<script lang="ts">
	import { enhance } from '$app/forms';
	import { Turnstile, Button } from '$lib/components';
	import { DoorOpen } from 'lucide-svelte';
	import { enhanceForm } from '@noxlovette/svarog';

	let password = $state('');
	let confirmPassword = $state('');
	let passwordMatch = $state(true);
</script>

<div
	class="flex w-11/12 max-w-md flex-col items-center justify-center space-y-6 rounded-xl bg-white p-9 shadow-md dark:bg-zinc-900"
>
	<div class="text-center">
		<h2 class="text-3xl font-bold text-indigo-600 dark:text-zinc-100">Create Account</h2>
		<p class="mt-2 text-sm text-zinc-600">
			Already have an account?
			<a href="/auth/login" class="font-medium text-indigo-500 hover:text-indigo-400 dark:text-zinc-100"
				>Sign in</a
			>
		</p>
	</div>

	<form
		method="post"
		class="w flex flex-col items-center justify-center space-y-4"
		use:enhance={enhanceForm({
			messages: {
				redirect: 'Welcome on board',
				defaultError: 'Signup Failed'
			},
			navigate: true
		})}
	>
		<div class="space-y-4">
			<div>
				<label for="name" class="block text-sm font-medium text-zinc-700">Full Name</label>
				<input
					type="text"
					name="name"
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				/>
			</div>

			<div>
				<label for="username" class="block text-sm font-medium text-zinc-700">Username</label>
				<input
					type="text"
					name="username"
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				/>
			</div>

			<div>
				<label for="role" class="block text-sm font-medium text-zinc-700">Role</label>
				<select
					name="role"
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				>
					<option value="">Select a role</option>
					<option value="teacher">Teacher</option>
					<option value="student">Student</option>
				</select>
			</div>

			<div>
				<label for="email" class="block text-sm font-medium text-zinc-700">Email</label>
				<input
					type="email"
					name="email"
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				/>
			</div>

			<div>
				<label for="password" class="block text-sm font-medium text-zinc-700">Password</label>
				<input
					type="password"
					name="password"
					bind:value={password}
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				/>
			</div>

			<div>
				<label for="confirmPassword" class="block text-sm font-medium text-zinc-700"
					>Confirm Password</label
				>
				<input
					type="password"
					name="confirmPassword"
					bind:value={confirmPassword}
					required
					class="w-full rounded-lg border border-zinc-200 px-4 py-2 transition duration-200 focus:ring focus:ring-indigo-500 focus:outline-none disabled:text-zinc-500
            dark:border-zinc-800 dark:bg-zinc-950 dark:focus:border-zinc-800 dark:focus:ring
                   dark:focus:ring-zinc-700 dark:focus:outline-none"
				/>
				{#if !passwordMatch}
					<p class="mt-1 text-sm text-red-600">Passwords don't match</p>
				{/if}
			</div>
		</div>
		<Turnstile />
		<Button Icon={DoorOpen} type="submit" variant="primary">Create Account</Button>
	</form>
</div>

<svelte:head>
	<title>Signup</title>
</svelte:head>
