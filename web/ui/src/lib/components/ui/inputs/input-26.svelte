<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import ArrowRight from '@lucide/svelte/icons/arrow-right';
	import Search from '@lucide/svelte/icons/search';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Search input with icon and button</Label>
	<div class="relative">
		<Input id={uid} class="peer ps-9 pe-9" placeholder="Search..." type="search" />
		<div
			class="text-muted-foreground/80 pointer-events-none absolute inset-y-0 start-0 flex items-center justify-center ps-3 peer-disabled:opacity-50"
		>
			<Search size={16} />
		</div>
		<button
			class="text-muted-foreground/80 ring-offset-background hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 absolute inset-y-px end-px flex h-full w-9 items-center justify-center rounded-e-lg transition-shadow focus-visible:border focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
			aria-label="Submit search"
			type="submit"
		>
			<ArrowRight size={16} aria-hidden="true" />
		</button>
	</div>
</div>
