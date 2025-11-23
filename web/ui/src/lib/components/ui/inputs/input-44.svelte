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
			'border-input bg-background text-foreground ring-offset-background relative flex size-9 items-center justify-center border-y border-e font-medium shadow-xs shadow-black/[.04] transition-all first:rounded-s-lg first:border-s last:rounded-e-lg',
			{ 'border-ring ring-ring/30 z-10 border ring-2 ring-offset-2': cell.isActive }
		)}
	>
		{#if cell.char !== null}
			<div>{cell.char}</div>
		{/if}
	</PinInput.Cell>
{/snippet}

<div class="*:not-first:mt-2">
	<Label for={uid}>OTP input single</Label>
	<PinInput.Root
		bind:value
		id={uid}
		class="flex items-center gap-3 has-disabled:opacity-50"
		maxlength={4}
	>
		{#snippet children({ cells })}
			<div class="flex">
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
