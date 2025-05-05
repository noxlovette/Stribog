<script>
	import { Button, H1, H3 } from '$lib/components';
</script>

<div class="flex flex-col items-center space-y-8 px-28 py-12">
	<div class="flex flex-col items-center space-y-2">
		<H1>Stribog</H1>
		<H3>Simple CDN</H3>
	</div>
	<div class="grid grid-cols-2 gap-3">
		<Button variant="secondary">Read the Docs</Button>
		<Button variant="primary">Get Started</Button>
	</div>
	<div class="h-96 w-96 overflow-hidden">
		<img
			src="/pictures/stribog-main-transparent.png"
			alt="Stribog at Forge"
			class="h-full w-full object-cover"
		/>
	</div>
</div>
