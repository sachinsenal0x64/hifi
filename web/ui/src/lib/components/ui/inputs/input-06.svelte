<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with error</Label>
	<Input
		id={uid}
		class="border-destructive/80 text-destructive focus-visible:border-destructive/80 focus-visible:ring-destructive/30"
		placeholder="Email"
		type="email"
		value="invalid@email.com"
	/>
	<p class="text-destructive text-xs" role="alert" aria-live="polite">Email is invalid</p>
</div>
