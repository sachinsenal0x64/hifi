<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2" style:--ring="234 89% 74%">
	<Label for={uid}>Input with colored border and ring</Label>
	<Input id={uid} placeholder="Email" type="email" />
</div>
