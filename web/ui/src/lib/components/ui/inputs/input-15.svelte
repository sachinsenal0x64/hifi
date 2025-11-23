<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with end add-on</Label>
	<div class="flex rounded-lg shadow-xs shadow-black/[.04]">
		<Input
			id={uid}
			class="z-10 -me-px rounded-e-none shadow-none"
			placeholder="google"
			type="text"
		/>
		<span
			class="border-input bg-background text-muted-foreground inline-flex items-center rounded-e-lg border px-3 text-sm"
		>
			.com
		</span>
	</div>
</div>
