<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with start add-on</Label>
	<div class="flex rounded-lg shadow-xs shadow-black/5">
		<span
			class="border-input bg-background text-muted-foreground inline-flex items-center rounded-s-lg border px-3 text-sm"
		>
			https://
		</span>
		<Input
			id={uid}
			class="-ms-px rounded-s-none shadow-none"
			placeholder="google.com"
			type="text"
		/>
	</div>
</div>
