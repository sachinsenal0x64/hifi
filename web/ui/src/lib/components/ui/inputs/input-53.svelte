<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';
</script>

<div class="*:not-first:mt-2">
	<Label for="input-52">Read-only input</Label>
	<Input
		id="input-52"
		class="read-only:bg-muted"
		value="This is a read-only input"
		readonly
		placeholder="Email"
		type="email"
	/>
</div>
