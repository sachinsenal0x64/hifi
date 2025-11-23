<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import CircleX from '@lucide/svelte/icons/circle-x';

	let inputValue = $state('Click to clear');
	let inputElement = $state<HTMLInputElement | null>(null)!;

	function handleClearInput() {
		inputValue = '';
	}

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with clear button</Label>
	<div class="relative">
		<Input
			id={uid}
			bind:ref={inputElement}
			bind:value={inputValue}
			class="pe-9"
			placeholder="Type something..."
			type="text"
		/>
		{#if inputValue}
			<button
				class="text-muted-foreground/80 ring-offset-background animate-in fade-in zoom-in-75 hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 absolute inset-y-px end-px flex h-full w-9 items-center justify-center rounded-e-lg border border-transparent transition-shadow focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
				aria-label="Clear input"
				onclick={handleClearInput}
			>
				<CircleX size={16} aria-hidden="true" />
			</button>
		{/if}
	</div>
</div>
