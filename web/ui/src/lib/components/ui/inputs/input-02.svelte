<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>
		Required input <span class="text-destructive">*</span>
	</Label>
	<Input id={uid} placeholder="Email" type="email" required />
</div>
