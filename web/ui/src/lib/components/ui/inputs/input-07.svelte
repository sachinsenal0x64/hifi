<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with gray background</Label>
	<Input
		id={uid}
		class="bg-muted border-transparent shadow-none"
		placeholder="Email"
		type="email"
	/>
</div>
