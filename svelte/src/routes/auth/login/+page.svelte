<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	import { Input, Turnstile, Button } from '$lib/components';
	import { setProfile, setUser, initialUser, initialProfile, notification } from '$lib/stores';
	import { enhanceForm } from '@noxlovette/svarog';
	import { DoorOpen } from 'lucide-svelte';
</script>

<div
	class="flex w-11/12 max-w-md flex-col items-center justify-center space-y-6 rounded-xl bg-white p-9 shadow-md dark:bg-slate-900"
>
	<div class="text-center">
		<h2 class="text-3xl font-bold text-sky-600 dark:text-slate-200">Welcome back</h2>
		<p class="mt-2 text-sm text-slate-600">
			Don't have an account?
			<a
				href="/auth/signup"
				class="font-medium text-sky-500 hover:text-sky-400 dark:text-slate-50 dark:hover:text-slate-200"
				>Sign up</a
			>
		</p>
	</div>
	<form
		method="POST"
		class="w flex flex-col items-center justify-center space-y-4"
		use:enhance={enhanceForm({
			messages: {
				failure: "Something's off"
			},
			handlers: {
				success: async (result) => {
					if (result.data) {
						const { user = initialUser, profile = initialProfile } = result.data;
						setUser(user);
						setProfile(profile);
						localStorage.setItem('user', JSON.stringify(user));
						localStorage.setItem('profile', JSON.stringify(profile));
						notification.set({ message: 'Welcome home', type: 'success' });
						await goto(user.role === 'teacher' ? '/t/dashboard' : '/s/dashboard');
					}
				}
			}
		})}
	>
		<div class="">
			<Input name="username" placeholder="Username" value="" />
		</div>
		<div class="">
			<Input name="password" placeholder="Password" value="" type="password" />
		</div>
		<Turnstile />
		<Button Icon={DoorOpen} type="submit" variant="primary" fullWidth={true}>Login</Button>
	</form>
</div>

<svelte:head>
	<title>Login</title>
</svelte:head>
