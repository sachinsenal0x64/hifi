<script lang="ts">
	import Input from '../ui/input.svelte';
	import Label from '../ui/label.svelte';

	const uid = $props.id();
</script>

<div class="[&>*:not(:first-child)]:mt-2">
	<Label class="flex-1">Range</Label>
	<div class="flex">
		<Input
			id="{uid}-1"
			class="flex-1 rounded-e-none [-moz-appearance:textfield] focus:z-10 [&::-webkit-inner-spin-button]:m-0 [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:m-0 [&::-webkit-outer-spin-button]:appearance-none"
			placeholder="From"
			type="number"
			aria-label="Min Value"
		/>
		<Input
			id="{uid}-2"
			class="-ms-px flex-1 rounded-s-none [-moz-appearance:textfield] focus:z-10 [&::-webkit-inner-spin-button]:m-0 [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:m-0 [&::-webkit-outer-spin-button]:appearance-none"
			placeholder="To"
			type="number"
			aria-label="Max Value"
		/>
	</div>
</div>
