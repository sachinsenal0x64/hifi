<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';
	import { cn } from '$lib/utils.js';

	import { PinInput, type PinInputCell } from 'bits-ui';

	let value = $state('');
	const uid = $props.id();
</script>

{#snippet Cell(cell: PinInputCell)}
	<PinInput.Cell
		{cell}
		class={cn(
			'border-input bg-background text-foreground flex size-9 items-center justify-center rounded-md border font-medium shadow-xs transition-[color,box-shadow]',
			{ 'border-ring ring-ring/50 z-10 ring-[3px]': cell.isActive }
		)}
	>
		{#if cell.char !== null}
			<div>{cell.char}</div>
		{/if}
	</PinInput.Cell>
{/snippet}

<div class="*:not-first:mt-2">
	<Label for={uid}>OTP input (spaced)</Label>
	<PinInput.Root
		bind:value
		id={uid}
		class="flex items-center gap-3 has-disabled:opacity-50"
		maxlength={4}
	>
		{#snippet children({ cells })}
			<div class="flex gap-2">
				<!-- eslint-disable-next-line svelte/require-each-key -->
				{#each cells as cell}
					{@render Cell(cell)}
				{/each}
			</div>
		{/snippet}
	</PinInput.Root>
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with <a
			class="hover:text-foreground underline"
			href="https://next.bits-ui.com/docs/components/pin-input"
			target="_blank"
			rel="noopener nofollow"
		>
			Bits UI PIN Input
		</a>
	</p>
</div>
