<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import Send from '@lucide/svelte/icons/send';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with end inline button</Label>
	<div class="relative">
		<Input id={uid} class="pe-9" placeholder="Email" type="email" />
		<button
			class="text-muted-foreground/80 ring-offset-background hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 absolute inset-y-px end-px flex h-full w-9 items-center justify-center rounded-e-lg border border-transparent transition-shadow focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
			aria-label="Subscribe"
		>
			<Send size={16} aria-hidden="true" />
		</button>
	</div>
</div>
