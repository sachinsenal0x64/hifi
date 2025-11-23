<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with inline start and end add-on</Label>
	<div class="relative flex rounded-lg shadow-xs shadow-black/[.04]">
		<span
			class="text-muted-foreground pointer-events-none absolute inset-y-0 start-0 flex items-center justify-center ps-3 text-sm"
		>
			€
		</span>
		<Input id={uid} class="-me-px rounded-e-none ps-6 shadow-none" placeholder="0.00" type="text" />
		<span
			class="border-input bg-background text-muted-foreground -z-10 inline-flex items-center rounded-e-lg border px-3 text-sm"
		>
			EUR
		</span>
	</div>
</div>
