<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with end button</Label>
	<div class="flex rounded-lg shadow-xs shadow-black/[.04]">
		<Input
			id={uid}
			class="-me-px flex-1 rounded-e-none shadow-none focus-visible:z-10"
			placeholder="Email"
			type="email"
		/>
		<button
			class="border-input bg-background text-foreground ring-offset-background hover:bg-accent hover:text-foreground focus-visible:border-ring focus-visible:ring-ring/30 inline-flex items-center rounded-e-lg border px-3 text-sm font-medium transition-shadow focus:z-10 focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
		>
			Send
		</button>
	</div>
</div>
