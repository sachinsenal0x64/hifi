<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with inline add-ons</Label>
	<div class="relative">
		<Input id={uid} class="peer ps-6 pe-12" placeholder="0.00" type="text" />
		<span
			class="text-muted-foreground pointer-events-none absolute inset-y-0 start-0 flex items-center justify-center ps-3 text-sm peer-disabled:opacity-50"
		>
			€
		</span>
		<span
			class="text-muted-foreground pointer-events-none absolute inset-y-0 end-0 flex items-center justify-center pe-3 text-sm peer-disabled:opacity-50"
		>
			EUR
		</span>
	</div>
</div>
