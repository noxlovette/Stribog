import { writable } from 'svelte/store';

function createLoadingStore() {
	const { subscribe, set } = writable(false);

	return {
		subscribe,
		true: () => set(true),
		false: () => set(false),
		toggle: () => {
			let currentValue;
			subscribe((value) => {
				currentValue = value;
			})();
			set(!currentValue);
		}
	};
}

export const isLoading = createLoadingStore();
